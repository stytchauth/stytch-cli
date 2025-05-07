#!/usr/bin/env sh

set -euo pipefail

echo "Installing stytch-cli..."
go install github.com/stytchauth/stytch-cli@latest

# First check for $(go env GOBIN)
gobin=$(go env GOBIN)
echo "gobin: $gobin"

# If that isn't set, use $(go env GOPATH)/bin
if [ -z "$gobin" ]; then
	gobin=$(go env GOPATH)/bin
fi

# If *that* isn't set, use $HOME/go/bin
if [ -z "$gobin" ]; then
	gobin=$HOME/go/bin
fi

echo "Removing existing stytch-cli..."
echo `rm -f $gobin/stytch-cli`

echo "Building and installing stytch-cli..."
echo `go build -o $gobin/stytch-cli main.go`

# Get the user's default shell
shell=$(basename $SHELL)
echo "Adding shell completion for $shell..."

# Run the completion script
$gobin/stytch-cli completion $shell >$HOME/.stytch-cli-completion.$shell

# Check if the user's rc file sources the completion file
if ! grep -q ".stytch-cli-completion.$shell" $HOME/.${shell}rc; then
       echo "Adding source line to .${shell}rc"
       echo "\nsource $HOME/.stytch-cli-completion.$shell" >>$HOME/.${shell}rc
fi

echo
echo "Installation complete. Run 'source $HOME/.${shell}rc' or restart your terminal to enable shell completions."