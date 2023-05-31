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

# ---------------------------------------------------------------------------- #
# Tasks
# ---------------------------------------------------------------------------- #

function iptables_configure() {
    sed --in-place \
        -e "s/{{PRIVATE_ETHERNET_INTERFACE}}/$IPTABLES_PRIVATE_ETHERNET_INTERFACE/g" \
        "/etc/iptables/iptables-rules.v4"

    systemctl daemon-reload
    systemctl enable iptables-activation
    systemctl restart iptables-activation
}

function sshd_configure() {
    sed --in-place \
        -E 's;^AllowUsers (.*);AllowUsers \1 git;' \
        /etc/ssh/sshd_config
}

function fail2ban_configure() {
    systemctl restart fail2ban
}

# ---------------------------------------------------------------------------- #
# Main
# ---------------------------------------------------------------------------- #

with-echo assert_sudo
with-echo restart_journald
with-echo remove_password_change_requirement
with-echo wait_cloud_init

with-echo deploy_provisioner_files

with-echo iptables_configure
# with-echo sshd_configure
with-echo fail2ban_configure
