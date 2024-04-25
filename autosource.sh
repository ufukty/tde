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
