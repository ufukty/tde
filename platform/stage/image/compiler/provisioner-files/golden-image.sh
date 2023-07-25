#!/bin/bash

# ---------------------------------------------------------------------------- #
# Variables
# ---------------------------------------------------------------------------- #

SUDO_USER="${SUDO_USER:?"SUDO_USER is required."}"

# ---------------------------------------------------------------------------- #
# Include
# ---------------------------------------------------------------------------- #

PROVISIONER_FILES="/home/$SUDO_USER/provisioner-files"
cd "$PROVISIONER_FILES"
. utilities.sh

# ---------------------------------------------------------------------------- #
# Function definitions
# ---------------------------------------------------------------------------- #

function go-remove-previous-versions() {
    rm -rf /usr/local/go
}

function go-env-var() {
    touch /etc/profile
    echo "export PATH=\$PATH:/usr/local/go/bin" >>/etc/profile
    . /etc/profile
}

function go-install() {
    tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
}

function go-check() {
    go version
}

# ---------------------------------------------------------------------------- #
# Main
# ---------------------------------------------------------------------------- #

assert_sudo
restart_journald
remove_password_change_requirement
wait_cloud_init

go-remove-previous-versions
go-env-var
go-install
go-check
