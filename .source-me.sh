#!/usr/bin/env bash -i

export WORKSPACE="$(pwd -P)"
. $WORKSPACE/shell/utilities.sh
test -f .source-me-untracked.sh && . .source-me-untracked.sh

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
    opts=$(grep '^Host' $WORKSPACE/platform/stage/ssh.conf 2>/dev/null | grep -v '[?*]' | cut -d ' ' -f 2-)

    COMPREPLY=($(compgen -W "$opts" -- ${cur}))
    return 0
}
complete -F _ssh_completion ssh

alias ssh="ssh -F $WORKSPACE/platform/stage/ssh.conf"
