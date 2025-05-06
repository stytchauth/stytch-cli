# stytch-cli
[Hackweek] Stytch CLI

To run commands:
- Set your workspace key id and secret in .env
- Run `bash install.sh`
- Use `stytch-cli <command>`
- Available commands:
  - `version` - prints version
  - `authenticate` - authenticate via connected apps, and store your access token locally
  - `project <command>`
    - `project create -v <vertical> -n <name>`
