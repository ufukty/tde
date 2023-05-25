#!/bin/bash

PROVISIONER_FILES="$(pwd -P)"

. utilities.sh

# ---------------------------------------------------------------------------- #
# Definitions
# ---------------------------------------------------------------------------- #

function install_utilities() {
    with-echo retry apt-get install -y ipcalc
}

function install_openvpn() {
    with-echo retry apt-get install -y ca-certificates gnupg openvpn iptables openssl wget ca-certificates curl
    test -d /etc/openvpn/easy-rsa && rm -rf /etc/openvpn/easy-rsa
    return "0"
}

function install_argon2() {
    with-echo tar -xvf phc-winner-argon2-20190702.tar.gz
    with-echo retry apt-get install -y gcc build-essential
    (
        with-echo cd phc-winner-argon2-20190702
        with-echo make install
    )
    with-echo rm -rf phc-winner-argon2-20190702 phc-winner-argon2-20190702.tar.gz
}

function install_easy_rsa() {
    mkdir -p /etc/openvpn/easy-rsa
    with-echo tar xzf EasyRSA-3.1.3.tgz --strip-components=1 --directory /etc/openvpn/easy-rsa
    with-echo rm -f EasyRSA-3.1.3.tgz
}

function install_ovpn_auth() {
    with-echo tar -xzf ovpn-auth-210813-linux-amd64.tar.gz
    with-echo mv ovpn-auth /etc/openvpn/ovpn-auth
    with-echo chmod 755 /etc/openvpn/ovpn-auth
    with-echo chown root:root /etc/openvpn/ovpn-auth
}

function install_unbound() {
    with-echo retry apt-get install -y unbound
}

# ---------------------------------------------------------------------------- #
# Main
# ---------------------------------------------------------------------------- #

with-echo assert_sudo
with-echo restart_journald
with-echo check_tun_availability
with-echo wait_cloud_init

with-echo apt_update

with-echo install_utilities
with-echo install_openvpn
with-echo install_argon2
with-echo install_easy_rsa
with-echo install_ovpn_auth
with-echo install_unbound

with-echo deploy_provisioner_files
