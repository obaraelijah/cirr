#!/bin/bash

set -e

INSTALL_DIR="$HOME/.cirr/bin"
CIRR_PATH="$INSTALL_DIR/cirr"

if [ ! -f "$CIRR_PATH" ]; then
    echo "cirr is not installed in $INSTALL_DIR."
    exit 0
fi

rm "$CIRR_PATH"

if [ -z "$(ls -A "$INSTALL_DIR")" ]; then
    rm -rf "$INSTALL_DIR"
    echo "Removed empty directory: $INSTALL_DIR"
fi

echo "cirr has been successfully uninstalled."

echo "Please remember to remove the following line from your shell configuration file (e.g., .bashrc, .zshrc):"
echo "export PATH=\"\$HOME/.cirr/bin:\$PATH\""
echo "After removing the line, restart your terminal or run 'source ~/.bashrc' (or the appropriate config file) to update your PATH."

echo "Uninstallation complete."