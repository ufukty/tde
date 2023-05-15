#!/usr/local/bin/bash -i

with-echo() {
    echo -e "\033[35m@echo\033[0m $@" && $@
    ret=$?
    if [ $ret -ne 0 ]; then
        echo -e "\033[35m@echo run has failed\033[0m" && exit $ret
    fi
}

note() {
    echo -e "\033[30m\033[43m\033[1m ${@} \033[0m"
}
