#!/bin/bash

PROVISIONER_FILES="$(pwd -P)"
. utilities.sh

# ---------------------------------------------------------------------------- #
# Definitions
# ---------------------------------------------------------------------------- #

function install_utilities() {
    retry apt-get install -y ipcalc
}

function install_openvpn() {
    retry apt-get install -y ca-certificates gnupg openvpn iptables openssl wget ca-certificates curl
    test -d /etc/openvpn/easy-rsa && rm -rf /etc/openvpn/easy-rsa
    return "0"
}

function install_argon2() {
    (
        cd dependencies
        tar -xvf phc-winner-argon2-20190702.tar.gz
        retry apt-get install -y gcc build-essential
        (
            cd phc-winner-argon2-20190702
            make install
        )
    )
}

function install_easy_rsa() {
    mkdir -p /etc/openvpn/easy-rsa
    (
        cd dependencies
        tar xzf EasyRSA-3.1.3.tgz --strip-components=1 --directory /etc/openvpn/easy-rsa
    )
}

function install_ovpn_auth() {
    (
        cd dependencies
        tar -xzf ovpn-auth-210813-linux-amd64.tar.gz
        mv ovpn-auth /etc/openvpn/ovpn-auth
        chmod 755 /etc/openvpn/ovpn-auth
        chown root:root /etc/openvpn/ovpn-auth
    )
}

function install_unbound() {
    retry apt-get install -y unbound
}

# ---------------------------------------------------------------------------- #
# Main
# ---------------------------------------------------------------------------- #

assert_sudo
restart_journald
check_tun_availability
wait_cloud_init

apt_update

install_utilities
install_openvpn
install_argon2
install_easy_rsa
install_ovpn_auth
install_unbound

deploy_provisioner_files
