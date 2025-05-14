#!/bin/bash

set -e

# 다운로드 URL 베이스 (버전에 맞게 수정하세요)
VERSION="v1.0.0"
BASE_URL="https://github.com/Ljaewon-123/register-cli/releases/download/$VERSION"

# 설치 위치
INSTALL_DIR="$HOME/.local/bin"
mkdir -p "$INSTALL_DIR"

# OS 구분
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
    echo "❌ 지원되지 않는 운영체제: $OS"
    exit 1
fi

# 다운로드 및 설치
echo "⬇️  $FILE 다운로드 중..."
curl -L "$BASE_URL/$FILE" -o "$INSTALL_DIR/$BIN_NAME"
chmod +x "$INSTALL_DIR/$BIN_NAME"

# PATH에 등록
if [[ "$OS" == "Linux" || "$OS" == "Darwin" ]]; then
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
        SHELL_RC="$HOME/.bashrc"
        [[ "$SHELL" == *zsh ]] && SHELL_RC="$HOME/.zshrc"
        echo "export PATH=\"$INSTALL_DIR:\$PATH\"" >> "$SHELL_RC"
        echo "🔧 PATH에 등록 완료 (다음 셸에서 적용됨)"
    fi
    echo "✅ 설치 완료! 터미널에서 'register-cli' 실행 가능"
else
    # Windows는 PATH 수동 설정 또는 PowerShell 필요
    echo "✅ 설치 완료: $INSTALL_DIR\\register-cli.exe"
    echo "👉 Windows는 PATH에 수동 등록하거나 PowerShell 스크립트 사용 필요"
fi
