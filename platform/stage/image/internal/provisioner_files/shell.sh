#!/bin/bash

# IMPORTANT:
# This script should be called from the directory that this script located in

# ---------------------------------------------------------------------------- #
# Required variables
# ---------------------------------------------------------------------------- #

IPTABLES_PRIVATE_ETHERNET_INTERFACE="${IPTABLES_PRIVATE_ETHERNET_INTERFACE:?"Example: eth1 (for DigitalOcean)"}"

# ---------------------------------------------------------------------------- #
# Optional Variables
# ---------------------------------------------------------------------------- #

PATH_OF_FILE="$(dirname "$(type -p $0)")"

# Name to append left sides of every line printed by script
ECHO_PREFIX="${ECHO_PREFIX:-"$(basename "$0")"}"

# ---------------------------------------------------------------------------- #
# Constants
# ---------------------------------------------------------------------------- #

PROVISIONER_FILES="/home/$SUDO_USER/provisioner_files"

# ---------------------------------------------------------------------------- #
# Includes
# ---------------------------------------------------------------------------- #

# shell_commons.sh file will be located at same directory if the script
# is running by a provisioner (vagrant, packer, etc) in guest/remote server
[ -f shell_commons.sh ] && . shell_commons.sh || . $(realpath $PATH_OF_FILE/../../companions/shell_commons.sh)

# ---------------------------------------------------------------------------- #
# Tasks
# ---------------------------------------------------------------------------- #

function iptables_configure() {
    info "Move the iptables-save template to its final location"
    mkdir -p "/etc/iptables"
    mv "$PROVISIONER_FILES/etc..iptables..picarus-custom-firewall.v4" \
        "/etc/iptables/picarus-custom-firewall.v4"

    info "Modify the iptables file to reflect the correct network adapter"
    sed --in-place \
        -e "s/<<PRIVATE_ETHERNET_INTERFACE>>/$IPTABLES_PRIVATE_ETHERNET_INTERFACE/g" \
        "/etc/iptables/picarus-custom-firewall.v4"

    info "Move the systemd service file to its final location"
    mv "$PROVISIONER_FILES/etc..systemd..system..picarus-custom-firewall.service" \
        "/etc/systemd/system/picarus-custom-firewall.service"

    info "Enable the custom systemd service and apply rules"
    systemctl daemon-reload
    systemctl enable picarus-custom-firewall
    systemctl restart picarus-custom-firewall
}

function sshd_configure() {
    info "configuring sshd to allow logins to user git"
    sed -r 's;^AllowUsers (.*);AllowUsers \1 git;' --in-place /etc/ssh/sshd_config
}

function fail2ban_configure() {
    info "configuring fail2ban to allow private network hosts"
    cat >/etc/fail2ban/jail.local <<HERE
[DEFAULT]
ignoreip  = 127.0.0.1/8 ::1 10.0.0.0/8
HERE
    systemctl restart fail2ban
}

# ---------------------------------------------------------------------------- #
# Main
# ---------------------------------------------------------------------------- #

execute_task assert_sudo
execute_task remove_password_change_requirement
execute_task wait_cloud_init

execute_task_in_golden_image iptables_configure
# execute_task_in_golden_image sshd_configure
execute_task_in_golden_image fail2ban_configure