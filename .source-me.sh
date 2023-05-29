#!/usr/bin/env bash -i

export WORKSPACE="$(pwd -P)"

with-echo() {
    echo -e "\033[35m@echo\033[0m $@" && $@
    ret=$?
    if [ $ret -ne 0 ]; then
        echo -e "\033[35m@echo run has failed\033[0m" && exit $ret
    fi
}
export -f with-echo

note() {
    echo -e "\033[30m\033[43m\033[1m ${@} \033[0m"
}
export -f note

error() {
    echo -e "\033[38m\033[41m\033[1m ${@} \033[0m"
}
export -f error

alias ssh="ssh -F $WORKSPACE/platform/stage/artifacts/ssh.conf"

test -f .source-me-untracked.sh && . .source-me-untracked.sh

check-python-pkg() {
    CLI_NAME="$1" && shift
    PIP_COMMAND="$@"

    if ! which "$CLI_NAME" >/dev/null; then
        note "CLI command '$CLI_NAME' is not found."
        echo "Run: $PIP_COMMAND"
    fi
}

_commands_completion() {
    _shortlist="$(cat commands | grep "() {" | tr -cd '[A-Za-z\-\n]')"
    local cur
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    COMPREPLY=($(compgen -W "${_shortlist}" -- ${cur}))
    return 0
}

complete -o nospace -F _commands_completion commands

_ssh_completion() {
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD - 1]}"
    opts=$(grep '^Host' $WORKSPACE/platform/stage/artifacts/ssh.conf 2>/dev/null | grep -v '[?*]' | cut -d ' ' -f 2-)

    COMPREPLY=($(compgen -W "$opts" -- ${cur}))
    return 0
}
complete -F _ssh_completion ssh

check-python-pkg ansible "python3 -m pip install --user ansible"
check-python-pkg qr "pip install qrcode"

check-argon2() {
    which argon2 >/dev/null || error "argon2 not found"
}

check-argon2

check-env-vars() {
    test -z "$DIGITALOCEAN_TOKEN" && error "DIGITALOCEAN_TOKEN is empty"
    test -z "$TF_VAR_DIGITALOCEAN_TOKEN" && error "TF_VAR_DIGITALOCEAN_TOKEN is empty"
    test -z "$TF_VAR_OVPN_USER" && error "TF_VAR_OVPN_USER is empty"
    test -z "$TF_VAR_OVPN_HASH" && error "TF_VAR_OVPN_HASH is empty"
}

check-env-vars
