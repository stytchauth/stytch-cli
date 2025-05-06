# stytch-cli
[Hackweek] Stytch CLI

To run commands:
- Set your workspace key id and secret in .env
- Run `go build -o ~/go/bin/stytch-cli` (ensure $HOME/go/bin is in your $PATH)
- Use `stytch-cli <command>`
- Available commands:
  - `version` - prints version
  - `create-b2b-project` - creates a new project in your workspace
