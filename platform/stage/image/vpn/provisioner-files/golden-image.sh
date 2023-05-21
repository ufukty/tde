#!/bin/bash

PROVISIONER_FILES="$(pwd -P)"

. provisioner-utilities.sh

# ---------------------------------------------------------------------------- #
# Definitions
# ---------------------------------------------------------------------------- #

function install_openvpn() {
    with-echo retry apt-get update
    with-echo retry apt-get install -y ca-certificates gnupg openvpn iptables openssl wget ca-certificates curl
    [[ -d /etc/openvpn/easy-rsa/ ]] && rm -rf /etc/openvpn/easy-rsa/
}

function install_argon2() {
    with-echo tar -xvf phc-winner-argon2-20190702.tar.gz
    with-echo retry apt-get install -y gcc build-essential
    (
        with-echo cd phc-winner-argon2-20190702
        with-echo make install
    )
    with-echo rm -rf $PROVISIONER_FILES/phc-winner-argon2-20190702 $PROVISIONER_FILES/phc-winner-argon2-20190702.tar.gz
}

function install_easy_rsa() {
    with-echo mkdir -p /etc/openvpn/easy-rsa
    (
        with-echo cd /etc/openvpn/easy-rsa
        with-echo tar xzf $PROVISIONER_FILES/EasyRSA-3.1.3.tgz --strip-components=1 --directory /etc/openvpn/easy-rsa
        with-echo rm -f $PROVISIONER_FILES/EasyRSA-3.1.3.tgz
    )
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

assert_sudo
check_tun_availability
wait_cloud_init

install_openvpn
install_argon2
install_easy_rsa
install_ovpn_auth
install_unbound
