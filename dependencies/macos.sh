#!/usr/local/bin/bash

set -e # exit on error

test -f "$HOME/venv/bin/activate" ||
    python -m venv "$HOME/venv"
source "$HOME/venv/bin/activate"

type _autosource ||
    (echo "copy assets/autosource.sh content into bash_profile" && exit 1)
type _commands_completion ||
    (echo "copy assets/commands.sh content into bash_profile" && exit 1)

test "$WORKSPACE" ||
    exit 1

which make ||
    xcode-select --install

which go ||
    (open "https://go.dev/dl" && exit 1)

which stringer ||
    go install "golang.org/x/tools/cmd/stringer"
which gonfique ||
    go install "github.com/ufukty/gonfique@v1.3.1"
which sqlc ||
    go install "github.com/sqlc-dev/sqlc/cmd/sqlc@latest"
which d2 ||
    go install "oss.terrastruct.com/d2@v0.6.3"

test -f "/usr/local/bin/bash" ||
    brew install "bash"
which psql ||
    (brew install "postgresql@15" && brew services start "postgresql@15")
(which terraform && which packer) ||
    brew tap "hashicorp/tap"
which terraform ||
    brew install "hashicorp/tap/terraform"
which packer ||
    brew install "hashicorp/tap/packer"
which openvpn ||
    brew install "openvpn"
which easyrsa ||
    (open "https://github.com/OpenVPN/easy-rsa" && exit 1)
which jq ||
    brew install jq

(which ansible && which qr) ||
    pip install -r "$WORKSPACE/dependencies/requirements.txt"

which argon2 ||
    (open "https://github.com/P-H-C/phc-winner-argon2/releases/tag/20190702" && exit 1)

which npm ||
    (open "https://nodejs.org/en/download" && exit 1)
which mmdc ||
    npm install -g "@mermaid-js/mermaid-cli"
