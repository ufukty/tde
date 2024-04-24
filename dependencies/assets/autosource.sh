#!/usr/local/bin/bash

function _autosource() {
    local START_DIR="$(pwd -P)"
    local ANCHOR="$(pwd -P)"
    while true; do
        if test -f "$ANCHOR/autosource.sh"; then
            cd "$ANCHOR" && source "$ANCHOR/autosource.sh"
        fi

        if test "$ANCHOR" == "/"; then
            break
        else
            ANCHOR="$(dirname "$ANCHOR")"
        fi
    done
    cd "$START_DIR"
}

_autosource
_cd() {
    local START_DIR="$(pwd -P)"
    builtin cd "$@"
    _autosource "$@"
    OLDPWD="$START_DIR"
}
alias cd="_cd"
