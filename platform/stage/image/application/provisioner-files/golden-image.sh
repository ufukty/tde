#!/bin/bash

# ---------------------------------------------------------------------------- #
# Variables
# ---------------------------------------------------------------------------- #

SUDO_USER="${SUDO_USER:?"SUDO_USER is required."}"
IPTABLES_PRIVATE_ETHERNET_INTERFACE="${IPTABLES_PRIVATE_ETHERNET_INTERFACE:?"IPTABLES_PRIVATE_ETHERNET_INTERFACE is required"}"

# ---------------------------------------------------------------------------- #
# Include
# ---------------------------------------------------------------------------- #

PROVISIONER_FILES="/home/$SUDO_USER/provisioner-files"
cd "$PROVISIONER_FILES"
. utilities.sh

# ---------------------------------------------------------------------------- #
# Function definitions
# ---------------------------------------------------------------------------- #

function app-service() {
    systemctl enable tde-app
    systemctl start tde-app
}

function deploy-tls-certificates() {
    # "moving files"
    mv "$PROVISIONER_FILES/app-db.crt" "/etc/ssl/certs/app-db.crt"
    mv "$PROVISIONER_FILES/app-db.key" "/etc/ssl/private/app-db.key"

    # "adjust permissions"
    chmod 755 "/etc/ssl/certs/app-db.crt"
    chmod 755 "/etc/ssl/private/app-db.key"

    # "adjust ownership"
    chown root:root "/etc/ssl/certs/app-db.crt"
    chown root:root "/etc/ssl/private/app-db.key"
}

function configure-logging() {
    # "creating directory and file for logs"
    mkdir "/var/log/picarus-sync-app"
    touch "/var/log/picarus-sync-app/backup.log"
    chown -R syslog:root "/var/log/picarus-sync-app"

    # "adjust permissions and ownership of /var/log/syslog"
    chown syslog:root "/var/log/syslog"
    chmod 640 "/var/log/syslog"

    # "moving rsyslog configuration"
    systemctl restart rsyslog

    # "managing logrotate"
}

# ---------------------------------------------------------------------------- #
# Main
# ---------------------------------------------------------------------------- #

assert_sudo
restart_journald
remove_password_change_requirement
check_tun_availability
wait_cloud_init
apt_update

deploy_provisioner_files

app-service
deploy-tls-certificates
configure-logging
