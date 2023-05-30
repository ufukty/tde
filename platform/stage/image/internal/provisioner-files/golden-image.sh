#!/bin/bash

# ---------------------------------------------------------------------------- #
# Variables
# ---------------------------------------------------------------------------- #

SUDO_USER="${SUDO_USER:?"SUDO_USER is required."}"

# A Linux and Postgres user will be created with this name, in addition to a Postgres Database
POSTGRES_USER="${POSTGRES_USER:?"POSTGRES_USER is required"}"

POSTGRES_SERVER_PRIVATE_IP="${POSTGRES_SERVER_PRIVATE_IP:?"POSTGRES_SERVER_PRIVATE_IP is required"}"

IPTABLES_PRIVATE_ETHERNET_INTERFACE="${IPTABLES_PRIVATE_ETHERNET_INTERFACE:?"IPTABLES_PRIVATE_ETHERNET_INTERFACE is required"}"

# ---------------------------------------------------------------------------- #
# Runtime values
# ---------------------------------------------------------------------------- #

# ---------------------------------------------------------------------------- #
# Include
# ---------------------------------------------------------------------------- #

PROVISIONER_FILES="/home/$SUDO_USER/provisioner-files"
cd "$PROVISIONER_FILES"
. utilities.sh

# ---------------------------------------------------------------------------- #
# Function definitions
# ---------------------------------------------------------------------------- #

function app-db-tunnel() {
    sed --in-place \
        -e "s/{{SUDO_USER}}/$SUDO_USER/g" \
        -e "s/{{POSTGRES_USER}}/$POSTGRES_USER/g" \
        -e "s/{{POSTGRES_SERVER_PRIVATE_IP}}/$POSTGRES_SERVER_PRIVATE_IP/g" \
        "/etc/systemd/system/app-db-tunnel.service"

    systemctl daemon-reload
    systemctl enable app-db-tunnel
    systemctl start app-db-tunnel
}

function app-service() {
    systemctl enable picarus-sync-backend
    systemctl start picarus-sync-backend
}

function configure-iptables() {
    sed --in-place \
        -e "s/{{PRIVATE_ETHERNET_INTERFACE}}/$IPTABLES_PRIVATE_ETHERNET_INTERFACE/g" \
        "/etc/iptables/custom-rules.v4"
    systemctl restart custom-rules
}

function configure-ssh() {

    mv "$PROVISIONER_FILES/sync-application-server.private" "/home/$SUDO_USER/.ssh/sync-application-server.private"

    # "add public key of Postgres server to .ssh/known_hosts"
    ssh-keyscan $POSTGRES_SERVER_PRIVATE_IP >>"/home/$SUDO_USER/.ssh/known_hosts"
    ssh-keyscan $POSTGRES_SERVER_PRIVATE_IP >>"/root/.ssh/known_hosts"

    # "update .ssh directory with correct ownership and permissions"
    chmod -R 700 "/home/$SUDO_USER/.ssh"
    chown -R $SUDO_USER:$SUDO_USER "/home/$SUDO_USER/.ssh"
}

function deploy-tls-certificates() {
    # "moving files"
    mv "$PROVISIONER_FILES/sync.picarus.net.crt" "/etc/ssl/certs/sync.picarus.net.crt"
    mv "$PROVISIONER_FILES/sync.picarus.net.key" "/etc/ssl/private/sync.picarus.net.key"

    # "adjust permissions"
    chmod 755 "/etc/ssl/certs/sync.picarus.net.crt"
    chmod 755 "/etc/ssl/private/sync.picarus.net.key"

    # "adjust ownership"
    chown root:root "/etc/ssl/certs/sync.picarus.net.crt"
    chown root:root "/etc/ssl/private/sync.picarus.net.key"
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

with-echo assert_sudo
with-echo restart_journald
with-echo remove_password_change_requirement
with-echo check_tun_availability
with-echo wait_cloud_init


with-echo app-db-tunnel
with-echo app-service
with-echo configure-iptables
with-echo configure-ssh
with-echo deploy-tls-certificates
with-echo configure-logging

with-echo apt_update

with-echo deploy_provisioner_files
