#!/bin/bash

set -e

# Base download URL (update version as needed)
VERSION="v1.0.0"
BASE_URL="https://github.com/Ljaewon-123/register-cli/releases/download/$VERSION"

# Installation directory
INSTALL_DIR="$HOME/.local/bin"
mkdir -p "$INSTALL_DIR"

# Detect OS
OS=$(uname -s)
FILE=""
BIN_NAME="register-cli"

if [[ "$OS" == "Linux" ]]; then
    FILE="register-cli-linux"
elif [[ "$OS" == "Darwin" ]]; then
    FILE="register-cli-macos"
elif [[ "$OS" == MINGW* || "$OS" == CYGWIN* || "$OS" == MSYS* ]]; then
    FILE="register-cli.exe"
    INSTALL_DIR="$HOME/AppData/Local/register-cli"
    mkdir -p "$INSTALL_DIR"
    BIN_NAME="register-cli.exe"
else
    echo "❌ Unsupported operating system: $OS"
    exit 1
fi

# Download and install
echo "⬇️  Downloading $FILE..."
curl -L "$BASE_URL/$FILE" -o "$INSTALL_DIR/$BIN_NAME"
chmod +x "$INSTALL_DIR/$BIN_NAME"

# Add to PATH
if [[ "$OS" == "Linux" || "$OS" == "Darwin" ]]; then
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
        SHELL_RC="$HOME/.bashrc"
        [[ "$SHELL" == *zsh ]] && SHELL_RC="$HOME/.zshrc"
        echo "export PATH=\"$INSTALL_DIR:\$PATH\"" >> "$SHELL_RC"
        echo "🔧 Added to PATH (will take effect in next shell session)"
    fi
    echo "✅ Installation complete! You can now run 'register-cli' from the terminal."
else
    # Windows: PATH needs to be set manually or via PowerShell
    echo "✅ Installation complete: $INSTALL_DIR\\register-cli.exe"
    echo "👉 On Windows, add the path to your system PATH manually or use a PowerShell script."
fi
