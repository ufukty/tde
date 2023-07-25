#!/bin/bash

PS4="\033[36m$(grealpath --relative-to="$(dirname "$WORKSPACE")" "$(pwd -P)")\033[32m/\$(basename \"\${BASH_SOURCE}\"):\${LINENO}\033[0m\033[33m\${FUNCNAME[0]:+/\${FUNCNAME[0]}():}\033[0m "
set -o xtrace

SUDO_USER="${SUDO_USER:?"SUDO_USER is required"}"

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
    retry apt-get update
}

function restart_journald() {
    systemctl restart systemd-journald
}

function assert_sudo() {
    if [[ "$EUID" > 0 ]]; then error "You need to run this script as root user (or with sudo)"; fi
}

function remove_password_change_requirement() {
    # info "remove password change requirement to root"
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
