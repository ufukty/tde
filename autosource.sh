#!/usr/local/bin/bash

export WORKSPACE
WORKSPACE="$(pwd -P)"

with-echo() {
    echo -e "\033[35m@echo\033[0m $*" && "$@"
    ret=$?
    if [ $ret -ne 0 ]; then
        echo -e "\033[35m@echo run has failed\033[0m" && exit $ret
    fi
}
export -f with-echo

note() {
    echo -e "\033[30m\033[43m\033[1m $* \033[0m"
}
export -f note

error() {
    echo -e "\033[38m\033[41m\033[1m $* \033[0m"
}
export -f error

alias ssh="ssh -F $WORKSPACE/platform/stage/artifacts/ssh.conf"

_check_ssh() {
    SSH_KEY_NAME="mbp-ed"
    ssh-add -l | grep ${SSH_KEY_NAME} >/dev/null
}

_enable_ssh() {
    note "Calling ssh-agent" && ssh-agent && ssh-add
}

_check_ssh || _enable_ssh

_check_env_vars() {
    DIGITALOCEAN_TOKEN="${DIGITALOCEAN_TOKEN:?}"
    TF_VAR_DIGITALOCEAN_TOKEN="${TF_VAR_DIGITALOCEAN_TOKEN:?}"
    TF_VAR_OVPN_USER="${TF_VAR_OVPN_USER:?}"
    TF_VAR_OVPN_HASH="${TF_VAR_OVPN_HASH:?}"
}
_check_env_vars

. "$HOME/venv/bin/activate"

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
