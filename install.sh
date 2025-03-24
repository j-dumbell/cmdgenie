#!/bin/sh
set -e

REPO="j-dumbell/cmdgenie"
BIN_NAME="cmdgenie"
INSTALL_DIR="/usr/local/bin"

OS=$(uname -s)
ARCH=$(uname -m)

case $OS in
  Darwin) PLATFORM="darwin" ;;
  Linux) PLATFORM="linux" ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

case $ARCH in
  x86_64) ARCH="amd64" ;;
  arm64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

LATEST_VERSION=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep '"tag_name":' | cut -d '"' -f 4)

URL="https://github.com/$REPO/releases/download/$LATEST_VERSION/${BIN_NAME}-${PLATFORM}-${ARCH}"
echo "Downloading $BIN_NAME from $URL..."
curl -L -o "$BIN_NAME" "$URL"

chmod +x "$BIN_NAME"
sudo mv "$BIN_NAME" "$INSTALL_DIR"

echo "$BIN_NAME installed successfully in $INSTALL_DIR"
