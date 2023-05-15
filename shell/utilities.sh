#!/usr/local/bin/bash -i

with-echo() {
    echo -e "\033[35m@echo\033[0m $@" && $@
    ret=$?
    if [ $ret -ne 0 ]; then
        echo -e "\033[35m@echo run has failed\033[0m" && exit $ret
    fi
}

shortlist() {
    typeset -f | awk '/ \(\) $/ { print $1 } ' | grep -v -e shortlist -e with-echo -e install_completion
}
