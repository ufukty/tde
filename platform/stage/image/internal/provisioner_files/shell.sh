#!/bin/bash

# IMPORTANT:
# This script should be called from the directory that this script located in

# ---------------------------------------------------------------------------- #
# Variables
# ---------------------------------------------------------------------------- #

PATH_OF_FILE="$(dirname "$(type -p $0)")"

# Name to append left sides of every line printed by script
ECHO_PREFIX="${ECHO_PREFIX:-"$(basename "$0")"}"

# A Linux and Postgres user will be created with this name, in addition to a Postgres Database
POSTGRES_USER="${POSTGRES_USER:?"Required. Example: picarus-sync-postgres"}"

POSTGRES_SERVER_PRIVATE_IP="${POSTGRES_SERVER_PRIVATE_IP:?"Required. Example: 10.137.0.3"}"

IPTABLES_PRIVATE_ETHERNET_INTERFACE="${IPTABLES_PRIVATE_ETHERNET_INTERFACE:?"Required. Example: eth1"}"

# ---------------------------------------------------------------------------- #
# Constants
# ---------------------------------------------------------------------------- #

PROVISIONER_FILES="/home/$SUDO_USER/provisioner_files"

# ---------------------------------------------------------------------------- #
# Include
# ---------------------------------------------------------------------------- #

# shell_commons.sh file will be located at same directory if the script
# is running by a provisioner (vagrant, packer, etc) in guest/remote server
[ -f shell_commons.sh ] && . shell_commons.sh || . $(realpath $PATH_OF_FILE/../../companions/shell_commons.sh)

# ---------------------------------------------------------------------------- #
# Function definitions
# ---------------------------------------------------------------------------- #

function wait_cloud_init() {
    cloud-init status --wait
}

function deploy_application_binary() {
    official_md5="04d93715d673213ade88f5822ec7cb74"

    mkdir -p /var/wwwapp
    tar -xzf "/home/$SUDO_USER/provisioner_files/linux-amd64-2021.11.21-16.51.06-UTC.tar.gz"
    mv linux-amd64 /var/wwwapp/picarus-sync
    chmod +x /var/wwwapp/picarus-sync
}

function create_and_enable_systemd_services() {
    info "deploy service definitions into /etc/systemd/system"
    mv \
        "$PROVISIONER_FILES/etc..systemd..system..picarus-sync-backend.service" \
        "/etc/systemd/system/picarus-sync-backend.service"
    mv \
        "$PROVISIONER_FILES/etc..systemd..system..picarus-sync-postgres-tunnel.service" \
        "/etc/systemd/system/picarus-sync-postgres-tunnel.service"

    info "template the service file"
    sed \
        --in-place \
        -e "s/<<SUDO_USER>>/$SUDO_USER/g" \
        -e "s/<<POSTGRES_USER>>/$POSTGRES_USER/g" \
        -e "s/<<POSTGRES_SERVER_PRIVATE_IP>>/$POSTGRES_SERVER_PRIVATE_IP/g" \
        "/etc/systemd/system/picarus-sync-postgres-tunnel.service"

    info "systemctl daemon-reload"
    systemctl daemon-reload
    
    info "enable and start 'picarus-sync-postgres-tunnel'"
    systemctl enable picarus-sync-postgres-tunnel
    systemctl start picarus-sync-postgres-tunnel
    
    info "enable and start 'picarus-sync-backend'"
    systemctl enable picarus-sync-backend 
    systemctl start picarus-sync-backend 
}

function iptables_configure() {
    info "Move the iptables-save template to its final location"
    mkdir -p "/etc/iptables"
    mv \
        "$PROVISIONER_FILES/etc..iptables..picarus-custom-firewall.v4" \
        "/etc/iptables/picarus-custom-firewall.v4"

    info "Modify the iptables file to reflect the correct network adapter"
    sed --in-place \
        -e "s/<<PRIVATE_ETHERNET_INTERFACE>>/$IPTABLES_PRIVATE_ETHERNET_INTERFACE/g" \
        "/etc/iptables/picarus-custom-firewall.v4"

    info "restart 'picarus-custom-firewall'"
    systemctl restart picarus-custom-firewall
}

function configure_ssh() {
    info "create .ssh directory"
    mkdir -p "/home/$SUDO_USER/.ssh"

    info "move private key of application server to .ssh directory"
    mv "$PROVISIONER_FILES/sync-application-server.private" "/home/$SUDO_USER/.ssh/sync-application-server.private"

    info "add public key of Postgres server to .ssh/known_hosts"
    ssh-keyscan $POSTGRES_SERVER_PRIVATE_IP >> "/home/$SUDO_USER/.ssh/known_hosts"
    ssh-keyscan $POSTGRES_SERVER_PRIVATE_IP >> "/root/.ssh/known_hosts"

    info "update .ssh directory with correct ownership and permissions"
    chmod -R 700 "/home/$SUDO_USER/.ssh"
    chown -R $SUDO_USER:$SUDO_USER "/home/$SUDO_USER/.ssh"
}

function deploy_tls_certificates() {
    info "moving files"
    mv "$PROVISIONER_FILES/sync.picarus.net.crt" "/etc/ssl/certs/sync.picarus.net.crt"
    mv "$PROVISIONER_FILES/sync.picarus.net.key" "/etc/ssl/private/sync.picarus.net.key"

    info "adjust permissions"
    chmod 755 "/etc/ssl/certs/sync.picarus.net.crt"
    chmod 755 "/etc/ssl/private/sync.picarus.net.key"

    info "adjust ownership"
    chown root:root "/etc/ssl/certs/sync.picarus.net.crt"
    chown root:root "/etc/ssl/private/sync.picarus.net.key"
}

function logging_enable() {
    info "creating directory and file for logs"
    mkdir "/var/log/picarus-sync-app"
    touch "/var/log/picarus-sync-app/backup.log"
    chown -R syslog:root "/var/log/picarus-sync-app"

    info "adjust permissions and ownership of /var/log/syslog"
    chown syslog:root "/var/log/syslog"
    chmod 640 "/var/log/syslog"

    info "moving rsyslog configuration"
    mv "$PROVISIONER_FILES/etc..rsyslog.d..picarus-sync-app.conf" "/etc/rsyslog.d/picarus-sync-app.conf"
    systemctl restart rsyslog

    info "managing logrotate"
    mv "$PROVISIONER_FILES/etc..logrotate.d..picarus-sync-app" "/etc/logrotate.d/picarus-sync-app"
    chown root:root "/etc/logrotate.d/picarus-sync-app"
}

# ---------------------------------------------------------------------------- #
# Main
# ---------------------------------------------------------------------------- #

execute_task assert_sudo
execute_task wait_cloud_init
execute_task remove_password_change_requirement

execute_task_in_golden_image logging_enable
execute_task_in_golden_image deploy_application_binary
execute_task_in_golden_image deploy_tls_certificates
execute_task_in_golden_image iptables_configure

execute_task_in_deployment configure_ssh
execute_task_in_deployment create_and_enable_systemd_services
