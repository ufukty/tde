#!/bin/bash

# ---------------------------------------------------------------------------- #
# Required variables
# ---------------------------------------------------------------------------- #

IPTABLES_PRIVATE_ETHERNET_INTERFACE="${IPTABLES_PRIVATE_ETHERNET_INTERFACE:?"IPTABLES_PRIVATE_ETHERNET_INTERFACE is required."}"
IPTABLES_PUBLIC_ETHERNET_INTERFACE="${IPTABLES_PUBLIC_ETHERNET_INTERFACE:?"IPTABLES_PUBLIC_ETHERNET_INTERFACE is required."}"

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
        -e "s/{{PUBLIC_ETHERNET_INTERFACE}}/$IPTABLES_PUBLIC_ETHERNET_INTERFACE/g" \
        "/etc/iptables/iptables-rules.v4"

    systemctl daemon-reload
    systemctl enable iptables-activation
    systemctl restart iptables-activation
}

# ---------------------------------------------------------------------------- #
# Main
# ---------------------------------------------------------------------------- #

assert_sudo
restart_journald
remove_password_change_requirement
wait_cloud_init

deploy_provisioner_files

iptables_configure
