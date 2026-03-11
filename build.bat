@echo off
setlocal

echo [1/2] Building spty.exe ...
go build -o installer\spty.exe .
if %ERRORLEVEL% neq 0 (
    echo Build failed.
    exit /b 1
)

echo [2/2] Creating MSI installer ...
wix build installer\spty.wxs -o installer\spty.msi
if %ERRORLEVEL% neq 0 (
    echo MSI creation failed.
    exit /b 1
)

echo.
echo Done: installer\spty.msi
