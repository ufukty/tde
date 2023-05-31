#!/bin/bash

function with-echo() {
    echo -e "\033[35m@echo\033[0m $@" && $@
    ret=$?
    if [ $ret -ne 0 ]; then
        echo -e "\033[35m@echo run has failed\033[0m" && exit $ret
    fi
}

function retry() {
    count=0
    until $@; do
        count=$((count + 1))
        if [[ $count -le 300 ]]; then
            echo "Attempted to run $1, but it's failed for $count times, now trying again..." && sleep 2
        else
            echo "Seems like $1 is busy right now, please try again later."
        fi
    done
}

function apt_update() {
    with-echo retry apt-get update
}

function restart_journald() {
    systemctl restart systemd-journald
}

function assert_sudo() {
    if [[ "$EUID" > 0 ]]; then error "You need to run this script as root user (or with sudo)"; fi
}

function remove_password_change_requirement() {
    sed --in-place -E 's/root:(.*):0:0:(.*):/root:\1:18770:0:\2:/g' /etc/shadow
}

function wait_cloud_init() {
    cloud-init status --wait >/dev/null
}

function check_tun_availability() {
    if [ ! -e /dev/net/tun ]; then error "TUN is not available at /dev/net/tun"; fi
}

function deploy_provisioner_files() {
    chmod 700 -R "$PROVISIONER_FILES/map"
    chown root:root -R "$PROVISIONER_FILES/map"
    rsync --verbose --recursive "$PROVISIONER_FILES/map/" "/"
    rm -rfv "$PROVISIONER_FILES/map"
}

export DEBIAN_FRONTEND=noninteractive
