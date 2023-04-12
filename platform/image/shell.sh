#!/bin/bash

function retry() {
    count=0
    until $@; do
        count=$((count + 1))
        if [[ $count -le 300 ]]; then
            info "Attempted to run $1, but it's failed for $count times, now trying again..." && sleep 2
        else
            error "Seems like $1 is busy right now, please try again later."
        fi
    done
}

retry apt-get update
retry apt-get install -y python3-pip cpulimit

pip3 install numpy matplotlib networkx deap