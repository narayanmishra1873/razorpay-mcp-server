@echo off
setlocal enabledelayedexpansion

:: Start script for Razorpay MCP Server (Windows)
:: Usage: start.bat [http|stdio] [additional_flags]

:: Default values
set TRANSPORT=%1
if "%TRANSPORT%"=="" set TRANSPORT=http

if "%RAZORPAY_API_KEY%"=="" (
    echo ❌ Error: RAZORPAY_API_KEY environment variable must be set
    echo Set them as environment variables:
    echo   set RAZORPAY_API_KEY=your_key
    echo   set RAZORPAY_API_SECRET=your_secret
    echo   start.bat
    echo.
    echo Or pass as flags:
    echo   start.bat http --key your_key --secret your_secret
    exit /b 1
)

if "%RAZORPAY_API_SECRET%"=="" (
    echo ❌ Error: RAZORPAY_API_SECRET environment variable must be set
    exit /b 1
)

:: Default values
if "%ADDRESS%"=="" set ADDRESS=:8080
if "%ENDPOINT_PATH%"=="" set ENDPOINT_PATH=/mcp
if "%READ_ONLY%"=="" set READ_ONLY=false

echo 🚀 Starting Razorpay MCP Server
echo Transport: %TRANSPORT%

:: Check if binary exists
if not exist "razorpay-mcp-server.exe" (
    if not exist "server.exe" (
        echo ❌ Binary not found. Building...
        go build -o razorpay-mcp-server.exe ./cmd/razorpay-mcp-server
        echo ✅ Build complete
    )
)

:: Use the right binary name
set BINARY=razorpay-mcp-server.exe
if exist "server.exe" set BINARY=server.exe

:: Build command arguments
set ARGS=

if "%TRANSPORT%"=="http" (
    echo Starting HTTP server on %ADDRESS%%ENDPOINT_PATH%
    set ARGS=!ARGS! --address %ADDRESS% --endpoint-path %ENDPOINT_PATH%
) else if "%TRANSPORT%"=="stdio" (
    echo Starting stdio server
    set ARGS=stdio !ARGS!
) else (
    echo ❌ Invalid transport: %TRANSPORT%. Use 'http' or 'stdio'
    exit /b 1
)

:: Add API keys
set ARGS=!ARGS! --key %RAZORPAY_API_KEY% --secret %RAZORPAY_API_SECRET%

:: Add optional parameters
if not "%TOOLSETS%"=="" set ARGS=!ARGS! --toolsets %TOOLSETS%
if "%READ_ONLY%"=="true" set ARGS=!ARGS! --read-only

:: Add any additional arguments (skip first argument which is transport)
shift
:loop
if "%1"=="" goto continue
set ARGS=!ARGS! %1
shift
goto loop
:continue

echo Command: %BINARY% %ARGS%
echo Press Ctrl+C to stop
echo.

:: Start the server
%BINARY% %ARGS%
