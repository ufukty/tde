#!/bin/bash

# ---------------------------------------------------------------------------- #
# Documentation
# ---------------------------------------------------------------------------- #

# Summary:           This script creates a OpenVPN server.

# Features:        - It Installs EasyRSA & OpenVPN, configures
#                    OpenVPN & iptables, creates .ovpn files for
#                    clients.
#                  - Clients can access to internet through
#                    server.
#                  - Clients can access to hosts available in
#                    the LAN of server
#                  - Clients are isolated from each other.
#                  - Same settings can be changed with optional
#                    variables including the encryption.

# Requirements:    - Fresh Ubuntu 20.04 server
#                  - Must run with sudo or by root user

# Example usage:     PUBLIC_IP="192.168.33.8" \
#                    VPC_ADDRESS="10.170.0.0" \
#                    PUBLIC_ETHERNET_INTERFACE="eth0" \
#                    PRIVATE_ETHERNET_INTERFACE="eth1" \
#                    SERVER_NAME="my_server" \
#                    EASYRSA_REQ_CN="my_server_common_name" \
#                    sudo --preserve-env bash sh.sh \
#                    my_client_1 my_client_2 my_client_n

# Fork:              This script is derived from another script
#                    https://github.com/angristan/openvpn-install

# ---------------------------------------------------------------------------- #
# Required Environment Variables
# ---------------------------------------------------------------------------- #

USER_ACCOUNT_NAME="${USER_ACCOUNT_NAME:?"USER_ACCOUNT_NAME is required."}"

PUBLIC_IP="${PUBLIC_IP:?"PUBLIC_IP is required."}"
PRIVATE_IP="${PRIVATE_IP:?"PRIVATE_IP is required."}"

# VPC_CIDR is the address of Virtual Private Cloud (LAN) + range mask
VPC_CIDR="${VPC_CIDR:?"VPC_CIDR is required."}"

# VPC_ADDRESS is the address of Virtual Private Cloud (LAN)
# It should be the address without the range suffix (like /24)
VPC_ADDRESS="${VPC_ADDRESS:?"VPC_ADDRESS is required."}"

# SUBNET_ADDRESS is the address of the subnet will be used by
# OpenVPN to handle underlying networking jobs. e.g. $SUBNET_ADDRESS
SUBNET_ADDRESS="${SUBNET_ADDRESS:?"SUBNET_ADDRESS is required. Example: 10.147.0.0"}"

# DNS_ADDRESS is the address od the internal DNS server (Unbound). It
# could be the address of first host in subnet.
DNS_ADDRESS="${DNS_ADDRESS:?"DNS_ADDRESS is required."}"

# PUBLIC_ETHERNET_INTERFACE
# Like eth0
PUBLIC_ETHERNET_INTERFACE="${PUBLIC_ETHERNET_INTERFACE:?"Required. Example: eth0 (for DigitalOcean)"}"

# PRIVATE_ETHERNET_INTERFACE is the name of the ethernet adapter
# used for connecting public internet. Like eth1
PRIVATE_ETHERNET_INTERFACE="${PRIVATE_ETHERNET_INTERFACE:?"Required. Example: eth1 (for DigitalOcean)"}"

# SERVER_NAME is used for EasyRSA, and it could be an arbitrary string
SERVER_NAME="${SERVER_NAME:?"SERVER_NAME is required."}"

# EASYRSA_REQ_CN is the common name of the server, and it could be an arbitrary string
EASYRSA_REQ_CN="${EASYRSA_REQ_CN:?"EASYRSA_REQ_CN is required."}"

# OTP_URI_ISSUER_NAME will be used to make OTP URI distinctive
OTP_URI_ISSUER_NAME="${OTP_URI_ISSUER_NAME:?"OTP_URI_ISSUER_NAME is required."}"

# OVPN_USERNAME and OVPN_HASH will be used as login credentials
OVPN_USERNAME="${OVPN_USERNAME:?"OVPN_USERNAME is required"}"
OVPN_HASH="${OVPN_HASH:?"OVPN_HASH is required"}"

# ---------------------------------------------------------------------------- #
# Optional Environment Variables
# ---------------------------------------------------------------------------- #

# PROTOCOL valid values:
# [ udp, tcp ]
PROTOCOL="${PROTOCOL:-"tcp"}"

PORT="${PORT:-"443"}"

# ENCRYPTION_TLS_SIG valid values:
# [ tls-crypt, tls-auth ]
ENCRYPTION_TLS_SIG=${ENCRYPTION_TLS_SIG:-"tls-crypt"}

# ENCRYPTION_CIPHER valid values:
# [ AES-128-GCM, AES-192-GCM, AES-256-GCM, AES-128-CBC, AES-192-CBC, AES-256-CBC ]
ENCRYPTION_CIPHER="${ENCRYPTION_CIPHER:-"AES-128-GCM"}"

# ENCRYPTION_CERT_TYPE valid values:
# [ ECDSA, RSA ]
ENCRYPTION_CERT_TYPE="${ENCRYPTION_CERT_TYPE:-"ECDSA"}"

# !!! Only usable when ENCRYPTION_CERT_TYPE="ECDSA"
# ENCRYPTION_CERT_CURVE valid values:
# [ prime256v1, secp384r1, secp521r1 ]
ENCRYPTION_ECDSA_CERT_CURVE="${ENCRYPTION_ECDSA_CERT_CURVE:-"prime256v1"}"

# !!! Only usable when ENCRYPTION_CERT_TYPE="ECDSA"
# ENCRYPTION_ECDSA_CC_CIPHER valid values:
# [ ECDHE-ECDSA-AES-128-GCM-SHA256, ECDHE-ECDSA-AES-256-GCM-SHA384 ]
ENCRYPTION_ECDSA_CC_CIPHER="${ENCRYPTION_ECDSA_CC_CIPHER:-"TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256"}"

# !!! Only usable when ENCRYPTION_CERT_TYPE="RSA"
# RSA_KEY_SIZE valid values:
# [ 2048, 3072, 4096 ]
ENCRYPTION_RSA_KEY_SIZE="${ENCRYPTION_RSA_KEY_SIZE:-"2048"}"

# !!! Only usable when ENCRYPTION_CERT_TYPE="RSA"
# ENCRYPTION_RSA_CC_CIPHER valid values:
# [ ECDHE-RSA-AES-128-GCM-SHA256, ECDHE-RSA-AES-256-GCM-SHA384 ]
ENCRYPTION_RSA_CC_CIPHER="${ENCRYPTION_RSA_CC_CIPHER:-"ECDHE-RSA-AES-128-GCM-SHA256"}"

# ENCRYPTION_DH_TYPE valid values:
# [ ECDH, DH ]
ENCRYPTION_DH_TYPE="${ENCRYPTION_DH_TYPE:-"ECDH"}"

# !!! Only usable when ENCRYPTION_DH_TYPE="ECDH"
# ENCRYPTION_ECDH_DH_CURVE valid values:
# [ prime256v1, secp384r1, secp521r1 ]
ENCRYPTION_ECDH_CURVE="${ENCRYPTION_ECDH_CURVE:-"prime256v1"}"

# !!! Only usable when ENCRYPTION_DH_TYPE="DH"
# ENCRYPTION_DH_KEY_SIZE valid values:
# [ 2048, 3072, 4096 ]
ENCRYPTION_DH_KEY_SIZE="${ENCRYPTION_DH_KEY_SIZE:-"2048"}"

# ENCRYPTION_HMAC_ALG valid values:
# [ SHA-256, SHA-384, SHA-512 ]
# - When GCM type ciphers are used, the algorithm is used only for
#   encryption of tls-auth packets from the control channel.
# - If, CBC type ciphers are used, the algorithm is used in addition
#   for authenticates data channel packets too.
ENCRYPTION_HMAC_ALG="${ENCRYPTION_HMAC_ALG:-"SHA256"}"

# TLS_SIG valid values:
# [ tls-crypt, tls-auth ]
# - Those will add additional layer of security to the control channel.
# - tls-auth authenticates the packets, while tls-crypt authenticate
#   and encrypt them.
TLS_SIG="${TLS_SIG:-"tls-crypt"}"

EASYRSA_REQ_COUNTRY="${EASYRSA_REQ_COUNTRY:-"US"}"
EASYRSA_REQ_PROVINCE="${EASYRSA_REQ_PROVINCE:-"NewYork"}"
EASYRSA_REQ_CITY="${EASYRSA_REQ_CITY:-"New York City"}"
EASYRSA_REQ_ORG="${EASYRSA_REQ_ORG:-"Lorem Ipsum"}"
EASYRSA_REQ_EMAIL="${EASYRSA_REQ_EMAIL:-"admin@loremi"}"
EASYRSA_REQ_OU="${EASYRSA_REQ_OU:-"Lorem Ipsum Pri"}"

# Iptables template will be filled with those variables' values
IPTABLES_PUBLIC_ETHERNET_INTERFACE="${IPTABLES_PUBLIC_ETHERNET_INTERFACE:-"$PUBLIC_ETHERNET_INTERFACE"}"
IPTABLES_PRIVATE_ETHERNET_INTERFACE="${IPTABLES_PRIVATE_ETHERNET_INTERFACE:-"$PRIVATE_ETHERNET_INTERFACE"}"
IPTABLES_OPENVPN_CUSTOM_PROTOCOL="${IPTABLES_OPENVPN_CUSTOM_PROTOCOL:-"$PROTOCOL"}"
IPTABLES_OPENVPN_CUSTOM_PORT="${IPTABLES_OPENVPN_CUSTOM_PORT:-"$PORT"}"
IPTABLES_OPENVPN_SUBNET_ADDRESS="${IPTABLES_OPENVPN_SUBNET_ADDRESS:-"$SUBNET_ADDRESS"}"

# ---------------------------------------------------------------------------- #
# Constants
# ---------------------------------------------------------------------------- #

PROVISIONER_FILES="$(pwd -P)"

# ---------------------------------------------------------------------------- #
# Imports
# ---------------------------------------------------------------------------- #

. provisioner-utilities.sh

# ---------------------------------------------------------------------------- #
# Definitions
# ---------------------------------------------------------------------------- #

function configure_easy_rsa() {
    (
        cd /etc/openvpn/easy-rsa

        echo "set_var EASYRSA_REQ_CN         \"$EASYRSA_REQ_CN\"" >vars
        echo "set_var EASYRSA_REQ_COUNTRY    \"$EASYRSA_REQ_COUNTRY\"" >>vars
        echo "set_var EASYRSA_REQ_PROVINCE   \"$EASYRSA_REQ_PROVINCE\"" >>vars
        echo "set_var EASYRSA_REQ_CITY       \"$EASYRSA_REQ_CITY\"" >>vars
        echo "set_var EASYRSA_REQ_ORG        \"$EASYRSA_REQ_ORG\"" >>vars
        echo "set_var EASYRSA_REQ_EMAIL      \"$EASYRSA_REQ_EMAIL\"" >>vars
        echo "set_var EASYRSA_REQ_OU         \"$EASYRSA_REQ_OU\"" >>vars

        if [[ $ENCRYPTION_CERT_TYPE == "ECDSA" ]]; then
            echo "set_var EASYRSA_ALGO           \"ec\"" >>vars
            echo "set_var EASYRSA_CURVE          \"$ENCRYPTION_ECDSA_CERT_CURVE\"" >>vars
        elif [[ $ENCRYPTION_CERT_TYPE == "RSA" ]]; then
            echo "set_var EASYRSA_KEY_SIZE       \"$ENCRYPTION_RSA_KEY_SIZE\"" >>vars
        fi

        # Create the PKI, set up the CA, the DH params and the server certificate
        ./easyrsa init-pki
        ./easyrsa --batch build-ca nopass

        # ECDH keys are generated on-the-fly so we don't need to generate them beforehand
        if [[ $ENCRYPTION_DH_TYPE == "DH" ]]; then
            openssl dhparam -out dh.pem $ENCRYPTION_DH_KEY_SIZE
        fi

        ./easyrsa build-server-full "$SERVER_NAME" nopass
        EASYRSA_CRL_DAYS=3650 ./easyrsa gen-crl

        # Generate tls-crypt or tls-auth key
        openvpn --genkey --secret /etc/openvpn/${TLS_SIG}.key
    )
}

function create_secrets_file_for_ovpn_auth() {
    (
        cd /etc/openvpn

        otp_secret="$(head -n 100 /dev/urandom | base32 | cut -b 1-64 | head -n 1)"
        authenticator_string="otpauth://totp/OpenVPN:${TOTP_USERNAME}@${OTP_URI_ISSUER_NAME}?secret=${otp_secret}&issuer=OpenVPN"

        (
            USERNAME=$(who am i | awk '{ print $1 }')
            sudo su "$USERNAME" -l
            touch "$HOME/otp-uri.txt"
            chmod 700 "$HOME/otp-uri.txt"
            chown $USER_ACCOUNT_NAME:$USER_ACCOUNT_NAME "$HOME/otp-uri.txt"
            echo -n "$authenticator_string" >$HOME/otp-uri.txt
        )

        sed --in-place \
            "<<OVPN_USERNAME>>" "$OVPN_USERNAME" \
            "<<OVPN_HASH>>" "$OVPN_HASH" \
            "<<OVPN_SECRET>>" "$otp_secret" \
            /etc/openvpn/secrets.yml

        chmod 744 secrets.yml
        chown root:root secrets.yml
    )
}

function configure_openvpn() {
    (
        cd /etc/openvpn/easy-rsa

        # Move all the generated files
        cp \
            "/etc/openvpn/easy-rsa/pki/ca.crt" \
            "/etc/openvpn/easy-rsa/pki/private/ca.key" \
            "/etc/openvpn/easy-rsa/pki/issued/$SERVER_NAME.crt" \
            "/etc/openvpn/easy-rsa/pki/private/$SERVER_NAME.key" \
            "/etc/openvpn/easy-rsa/pki/crl.pem" \
            "/etc/openvpn"

        if [[ $ENCRYPTION_DH_TYPE == "ECDH" ]]; then
            DH_CONF_STR="dh none"$'\n'"ecdh-curve $ENCRYPTION_ECDH_CURVE"
        elif [[ $ENCRYPTION_DH_TYPE == "DH" ]]; then
            cp dh.pem /etc/openvpn
            DH_CONF_STR="dh dh.pem"
        fi

        chmod 644 /etc/openvpn/crl.pem

        # Find out if the machine uses nogroup or nobody for the permissionless group
        if grep -qs "^nogroup:" /etc/group; then
            NOGROUP=nogroup
        else
            NOGROUP=nobody
        fi

        if [[ $ENCRYPTION_CERT_TYPE == "ECDSA" ]]; then
            ENCRYPTION_CC_CIPHER="$ENCRYPTION_ECDSA_CC_CIPHER"
        elif [[ $ENCRYPTION_CERT_TYPE == "RSA" ]]; then
            ENCRYPTION_CC_CIPHER="$ENCRYPTION_RSA_CC_CIPHER"
        fi

        # "Populating the configure file at: /etc/openvpn/server.conf"
        sed --in-place \
            -e "s;{{DH_CONF_STR}};$DH_CONF_STR;g" \
            -e "s;{{DNS_ADDRESS}};$DNS_ADDRESS;g" \
            -e "s;{{ENCRYPTION_CC_CIPHER}};$ENCRYPTION_CC_CIPHER;g" \
            -e "s;{{ENCRYPTION_CIPHER}};$ENCRYPTION_CIPHER;g" \
            -e "s;{{ENCRYPTION_HMAC_ALG}};$ENCRYPTION_HMAC_ALG;g" \
            -e "s;{{NOGROUP}};$NOGROUP;g" \
            -e "s;{{PORT}};$PORT;g" \
            -e "s;{{PROTOCOL}};$PROTOCOL;g" \
            -e "s;{{SERVER_NAME}};$SERVER_NAME;g" \
            -e "s;{{SUBNET_ADDRESS}};$SUBNET_ADDRESS;g" \
            -e "s;{{TLS_SIG}};$TLS_SIG;g" \
            -e "s;{{VPC_ADDRESS}};$VPC_ADDRESS;g" \
            /etc/openvpn/server.conf

        # Create client-config-dir dir
        mkdir -p /etc/openvpn/ccd

        # Create log dir
        mkdir -p /var/log/openvpn

        # Enable routing
        # echo 'net.ipv4.ip_forward=1' >/etc/sysctl.d/20-openvpn.conf

        # "Apply sysctl rules"
        sysctl --system

        # If SELinux is enabled and a custom port was selected, we need this
        if hash sestatus 2>/dev/null; then
            if sestatus | grep "Current mode" | grep -qs "enforcing"; then
                if [[ $PORT != '1194' ]]; then
                    semanage port -a -t openvpn_port_t -p "$PROTOCOL" "$PORT"
                fi
            fi
        fi

        # Don't modify package-provided service
        cp /lib/systemd/system/openvpn\@.service /etc/systemd/system/openvpn\@.service

        # Workaround to fix OpenVPN service on OpenVZ
        sed -i 's|LimitNPROC|#LimitNPROC|' /etc/systemd/system/openvpn\@.service
        # Another workaround to keep using /etc/openvpn/
        sed -i 's|/etc/openvpn/server|/etc/openvpn|' /etc/systemd/system/openvpn\@.service

        with-echo systemctl daemon-reload
        with-echo systemctl enable openvpn@server
        with-echo systemctl restart openvpn@server
    )
}

function configure_iptables() {
    sed --in-place \
        -e "s;{{PRIVATE_ETHERNET_INTERFACE}};$IPTABLES_PRIVATE_ETHERNET_INTERFACE;g" \
        -e "s;{{PUBLIC_ETHERNET_INTERFACE}};$IPTABLES_PUBLIC_ETHERNET_INTERFACE;g" \
        -e "s;{{OPENVPN_CUSTOM_PROTOCOL}};$IPTABLES_OPENVPN_CUSTOM_PROTOCOL;g" \
        -e "s;{{OPENVPN_CUSTOM_PORT}};$IPTABLES_OPENVPN_CUSTOM_PORT;g" \
        -e "s;{{OPENVPN_SUBNET_ADDRESS}};$IPTABLES_OPENVPN_SUBNET_ADDRESS;g" \
        /etc/iptables/iptables-rules.v4

    with-echo systemctl daemon-reload
    with-echo systemctl enable iptables-activation
    with-echo systemctl restart iptables-activation
}

function configure_unbound() {
    sed --in-place \
        -e "s;{{DNS_ADDRESS}};$DNS_ADDRESS;g" \
        -e "s;{{SUBNET_ADDRESS}};$SUBNET_ADDRESS;g" \
        -e "s;{{HOST_ADDRESS}};$PRIVATE_IP;g" \
        -e "s;{{VPC_CIDR}};$VPC_CIDR;g" \
        /etc/unbound/unbound.conf

    systemctl enable unbound
    systemctl restart unbound
}

with-echo assert_sudo
with-echo check_tun_availability
with-echo wait_cloud_init

with-echo configure_easy_rsa
with-echo configure_openvpn
with-echo create_secrets_file_for_ovpn_auth
with-echo configure_iptables
with-echo configure_unbound
