package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/T4ko0522/spotify-cli/internal/config"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

const redirectURI = "http://127.0.0.1:8888/callback"

var scopes = []string{
	"user-read-playback-state",
	"user-modify-playback-state",
	"user-read-currently-playing",
}

func newAuthenticator() *spotifyauth.Authenticator {
	return spotifyauth.New(
		spotifyauth.WithClientID(config.ClientID),
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(scopes...),
	)
}

func generateVerifier() (string, error) {
	buf := make([]byte, 64)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}

func generateChallenge(verifier string) string {
	h := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(h[:])
}

func generateState() (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(buf), nil
}

func Login() error {
	auth := newAuthenticator()

	verifier, err := generateVerifier()
	if err != nil {
		return fmt.Errorf("failed to generate PKCE verifier: %w", err)
	}
	challenge := generateChallenge(verifier)

	state, err := generateState()
	if err != nil {
		return fmt.Errorf("failed to generate state: %w", err)
	}

	codeCh := make(chan string, 1)
	errCh := make(chan error, 1)

	mux := http.NewServeMux()
	server := &http.Server{Addr: ":8888", Handler: mux}

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		if s := r.URL.Query().Get("state"); s != state {
			errCh <- fmt.Errorf("state mismatch")
			fmt.Fprint(w, "State mismatch. Please try again.")
			return
		}
		if e := r.URL.Query().Get("error"); e != "" {
			errCh <- fmt.Errorf("authorization denied: %s", e)
			fmt.Fprintf(w, "Authorization denied: %s", e)
			return
		}
		code := r.URL.Query().Get("code")
		if code == "" {
			errCh <- fmt.Errorf("no code in callback")
			fmt.Fprint(w, "No authorization code received.")
			return
		}
		fmt.Fprint(w, "Login successful! You can close this tab.")
		codeCh <- code
	})

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			errCh <- fmt.Errorf("callback server error: %w", err)
		}
	}()

	url := auth.AuthURL(state,
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("code_challenge", challenge),
	)

	fmt.Println("Opening browser for Spotify login...")
	if err := openBrowser(url); err != nil {
		fmt.Printf("Open this URL in your browser:\n%s\n", url)
	}

	var code string
	select {
	case code = <-codeCh:
	case err := <-errCh:
		_ = server.Shutdown(context.Background())
		return err
	}

	_ = server.Shutdown(context.Background())

	token, err := auth.Exchange(context.Background(), code,
		oauth2.SetAuthURLParam("code_verifier", verifier),
	)
	if err != nil {
		return fmt.Errorf("token exchange failed: %w", err)
	}

	if err := SaveToken(token); err != nil {
		return fmt.Errorf("failed to save token: %w", err)
	}

	fmt.Println("Logged in successfully.")
	return nil
}

var currentTokenSource oauth2.TokenSource

func openBrowser(url string) error {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("open", url).Start()
	case "linux":
		return exec.Command("xdg-open", url).Start()
	default: // windows
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	}
}

func oauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID: config.ClientID,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://accounts.spotify.com/authorize",
			TokenURL:  "https://accounts.spotify.com/api/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: redirectURI,
		Scopes:      scopes,
	}
}

// GetClient returns an HTTP client with automatic token refresh.
func GetClient(ctx context.Context) (*http.Client, error) {
	token, err := LoadToken()
	if err != nil {
		return nil, err
	}
	cfg := oauthConfig()
	currentTokenSource = cfg.TokenSource(ctx, token)
	return oauth2.NewClient(ctx, currentTokenSource), nil
}

// PersistToken saves the current (possibly refreshed) token to disk.
func PersistToken() error {
	if currentTokenSource == nil {
		return nil
	}
	token, err := currentTokenSource.Token()
	if err != nil {
		return nil
	}
	return SaveToken(token)
}
