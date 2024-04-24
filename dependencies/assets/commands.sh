#!/usr/local/bin/bash

_commands_completion() {
    _shortlist="$(cat commands | grep -E "^[^ ]+\(\) [\{\(]" | tr -cd '[A-Za-z\-\n]')"
    local cur
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    COMPREPLY=($(compgen -W "${_shortlist}" -- ${cur}))
    return 0
}
complete -o nospace -F _commands_completion commands
