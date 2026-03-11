@echo off
setlocal

echo [1/2] Building spt.exe ...
go build -o installer\spt.exe .
if %ERRORLEVEL% neq 0 (
    echo Build failed.
    exit /b 1
)

echo [2/2] Creating MSI installer ...
wix build installer\spt.wxs -o installer\spt.msi
if %ERRORLEVEL% neq 0 (
    echo MSI creation failed.
    exit /b 1
)

echo.
echo Done: installer\spt.msi
