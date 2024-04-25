#!/usr/local/bin/bash

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
