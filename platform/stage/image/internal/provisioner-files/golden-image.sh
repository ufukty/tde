#!/bin/bash

# ---------------------------------------------------------------------------- #
# Required variables
# ---------------------------------------------------------------------------- #

IPTABLES_PRIVATE_ETHERNET_INTERFACE="${IPTABLES_PRIVATE_ETHERNET_INTERFACE:?"IPTABLES_PRIVATE_ETHERNET_INTERFACE is required."}"

# ---------------------------------------------------------------------------- #
# Constants
# ---------------------------------------------------------------------------- #

PROVISIONER_FILES="/home/$SUDO_USER/provisioner-files"
cd "$PROVISIONER_FILES"
. utilities.sh
. secrets.sh

# ---------------------------------------------------------------------------- #
# Tasks
# ---------------------------------------------------------------------------- #

function create-app-user() {
    adduser --disabled-password --gecos "" "$APP_USER"
}

function iptables_configure() {
    sed --in-place \
        -e "s/{{PRIVATE_ETHERNET_INTERFACE}}/$IPTABLES_PRIVATE_ETHERNET_INTERFACE/g" \
        "/etc/iptables/iptables-rules.v4"

    systemctl daemon-reload
    systemctl enable iptables-activation
    systemctl restart iptables-activation
}

function sshd_configure() {
    sed --in-place -E 's;^AllowUsers (.*);AllowUsers \1 git;' /etc/ssh/sshd_config
}

function fail2ban_configure() {
    systemctl restart fail2ban
}

# ---------------------------------------------------------------------------- #
# Main
# ---------------------------------------------------------------------------- #

assert_sudo
restart_journald
remove_password_change_requirement
wait_cloud_init

deploy_provisioner_files

create-app-user
iptables_configure
# sshd_configure
fail2ban_configure
