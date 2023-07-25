#!/bin/bash

# ---------------------------------------------------------------------------- #
# Variables
# ---------------------------------------------------------------------------- #

# A Linux and Postgres user will be created with this name, in addition to a Postgres Database
POSTGRES_USER="${POSTGRES_USER:?"POSTGRES_USER is required"}"
POSTGRES_SERVER_PRIVATE_IP="${POSTGRES_SERVER_PRIVATE_IP:?"POSTGRES_SERVER_PRIVATE_IP is required"}"

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

function configure-iptables() {
    sed --in-place \
        -e "s/{{PRIVATE_ETHERNET_INTERFACE}}/$IPTABLES_PRIVATE_ETHERNET_INTERFACE/g" \
        "/etc/iptables/custom-rules.v4"

    systemctl restart custom-rules
}

function configure-ssh() {
    # "add public key of Postgres server to .ssh/known_hosts"
    ssh-keyscan $POSTGRES_SERVER_PRIVATE_IP >>"/home/$SUDO_USER/.ssh/known_hosts"
    ssh-keyscan $POSTGRES_SERVER_PRIVATE_IP >>"/root/.ssh/known_hosts"

    # "update .ssh directory with correct ownership and permissions"
    chmod -R 700 "/home/$SUDO_USER/.ssh"
    chown -R $SUDO_USER:$SUDO_USER "/home/$SUDO_USER/.ssh"
}

app-db-tunnel
configure-iptables
configure-ssh
