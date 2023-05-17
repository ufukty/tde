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

# ---------------------------------------------------------------------------- #
# Optional Environment Variables
# ---------------------------------------------------------------------------- #

# PROTOCOL valid values:
# [ udp, tcp ]
PROTOCOL="${PROTOCOL:-"tcp"}"

PORT="${PORT:-"443"}"

DNS_PRIMARY="${DNS:-"208.67.222.222"}"   # OpenDNS
DNS_SECONDARY="${DNS:-"208.67.220.220"}" # OpenDNS

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

EASY_RSA_VERSION="${EASY_RSA_VERSION:-"3.0.7"}"

EASYRSA_REQ_COUNTRY="${EASYRSA_REQ_COUNTRY:-"US"}"
EASYRSA_REQ_PROVINCE="${EASYRSA_REQ_PROVINCE:-"NewYork"}"
EASYRSA_REQ_CITY="${EASYRSA_REQ_CITY:-"New York City"}"
EASYRSA_REQ_ORG="${EASYRSA_REQ_ORG:-"Lorem Ipsum"}"
EASYRSA_REQ_EMAIL="${EASYRSA_REQ_EMAIL:-"admin@loremi"}"
EASYRSA_REQ_OU="${EASYRSA_REQ_OU:-"Lorem Ipsum Pri"}"

PATH_OF_FILE="$(dirname "$(type -p $0)")"

# Name to append left sides of every line printed by script
ECHO_PREFIX="${ECHO_PREFIX:-"$(basename "$0")"}"

# Iptables template will be filled with those variables' values
IPTABLES_PUBLIC_ETHERNET_INTERFACE="${IPTABLES_PUBLIC_ETHERNET_INTERFACE:-"$PUBLIC_ETHERNET_INTERFACE"}"
IPTABLES_PRIVATE_ETHERNET_INTERFACE="${IPTABLES_PRIVATE_ETHERNET_INTERFACE:-"$PRIVATE_ETHERNET_INTERFACE"}"
IPTABLES_OPENVPN_CUSTOM_PROTOCOL="${IPTABLES_OPENVPN_CUSTOM_PROTOCOL:-"$PROTOCOL"}"
IPTABLES_OPENVPN_CUSTOM_PORT="${IPTABLES_OPENVPN_CUSTOM_PORT:-"$PORT"}"
IPTABLES_OPENVPN_SUBNET_ADDRESS="${IPTABLES_OPENVPN_SUBNET_ADDRESS:-"$SUBNET_ADDRESS"}"

# ---------------------------------------------------------------------------- #
# Constants
# ---------------------------------------------------------------------------- #

PROVISIONER_FILES="/home/$SUDO_USER/provisioner_files"

# ---------------------------------------------------------------------------- #
# Imports
# ---------------------------------------------------------------------------- #

# provisioner_utilities.sh file will be located at same directory if the script
# is running by a provisioner (vagrant, packer, etc) in guest/remote server
[ -f provisioner_utilities.sh ] && . provisioner_utilities.sh || . $(realpath $PATH_OF_FILE/../../../../../shell/provisioner_utilities.sh)

# ---------------------------------------------------------------------------- #
# Definitions
# ---------------------------------------------------------------------------- #

function replace_in_template_file() {
    sed "s/$2/$3/g" --in-place $1
}

function install_unbound() {
    retry apt-get install -y unbound
}

function configure_unbound() {
    mv ~/provisioner_files/etc/unbound/unbound.conf /etc/unbound/unbound.conf

    touch /etc/unbound/custom.conf

    replace_in_template_file "/etc/unbound/unbound.conf" "<<DNS_ADDRESS>>" "$DNS_ADDRESS"
    replace_in_template_file "/etc/unbound/unbound.conf" "<<SUBNET_ADDRESS>>" "$SUBNET_ADDRESS"
    replace_in_template_file "/etc/unbound/unbound.conf" "<<HOST_ADDRESS>>" "$PRIVATE_IP"
    sed "s;<<VPC_CIDR>>;$VPC_CIDR;g" --in-place "/etc/unbound/unbound.conf"

    systemctl enable unbound
    systemctl restart unbound
}

function install_openvpn() {
    retry apt-get update
    retry apt-get install -y ca-certificates gnupg openvpn iptables openssl wget ca-certificates curl
    # An old version of easy-rsa was available by default in some openvpn packages
    [[ -d /etc/openvpn/easy-rsa/ ]] && rm -rf /etc/openvpn/easy-rsa/
}

function install_argon2() {
    cd ~

    info "Download & check argon2 from github"
    wget https://github.com/P-H-C/phc-winner-argon2/archive/refs/tags/20190702.tar.gz
    sha512sum 20190702.tar.gz | grep -q 0a4cb89e8e63399f7d || error "Bad sha512 for argon2 download"

    info "Extract files"
    tar -xvf 20190702.tar.gz

    info "Install gcc & build-essential"
    retry apt install -y gcc build-essential

    info "Compile and install libargon2"
    cd phc-winner-argon2-20190702
    sudo make install

    info "Clean source code"
    cd ~
    rm -rf ~/phc-winner-argon2-20190702 ~/20190702.tar.gz
}

function install_easy_rsa() {
    mkdir -p /etc/openvpn/easy-rsa
    cd /etc/openvpn/easy-rsa

    info "Install the latest version of easy-rsa from source"
    wget -O ~/easy-rsa.tgz \
        https://github.com/OpenVPN/easy-rsa/releases/download/v${EASY_RSA_VERSION}/EasyRSA-${EASY_RSA_VERSION}.tgz
    tar xzf ~/easy-rsa.tgz --strip-components=1 --directory /etc/openvpn/easy-rsa
    rm -f ~/easy-rsa.tgz
}

function configure_easy_rsa() {
    cd /etc/openvpn/easy-rsa

    info "Creating vars file for easy-rsa"

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
}

function install_ovpn_auth() {
    cd ~/provisioner_files

    info "checking sha512sum of archive"
    sha512sum ovpn-auth-210813-linux-amd64.tar.gz | grep -q af730e946e477aa7d1153f836a68b8706170 || error "Bad sha512 for ovpn-auth download"

    info "unarchive ovpn-auth-210813-linux-amd64.tar.gz"
    tar -xzf ovpn-auth-210813-linux-amd64.tar.gz

    info "moving ovpn-auth into /etc/openvpn"
    mv etc/openvpn/ovpn-auth /etc/openvpn/ovpn-auth

    info "change the file permission and ownership of ovpn-auth"
    chmod 755 /etc/openvpn/ovpn-auth
    chown root:root /etc/openvpn/ovpn-auth
}

function create_secrets_file_for_ovpn_auth() {
    cd /etc/openvpn/
    mv ~/provisioner_files/etc/openvpn/secrets.yml /etc/openvpn/secrets.yml

    otp_secret="$(head -n 100 /dev/urandom | base32 | cut -b 1-64 | head -n 1)"
    authenticator_string="otpauth://totp/OpenVPN:ufukty@${OTP_URI_ISSUER_NAME}?secret=${otp_secret}&issuer=OpenVPN"

    touch "$HOME/otp-uri.txt"
    chmod 700 "$HOME/otp-uri.txt"
    chown $USER_ACCOUNT_NAME:$USER_ACCOUNT_NAME "$HOME/otp-uri.txt"
    echo -n "$authenticator_string" >$HOME/otp-uri.txt

    replace_in_template_file "/etc/openvpn/secrets.yml" "<<OTP_SECRET>>" "$otp_secret"

    chmod 744 secrets.yml
    chown root:root secrets.yml
}

function configure_openvpn() {
    cd /etc/openvpn/easy-rsa

    # Move all the generated files
    cp "pki/ca.crt" \
        "pki/private/ca.key" \
        "pki/issued/$SERVER_NAME.crt" \
        "pki/private/$SERVER_NAME.key" \
        "/etc/openvpn/easy-rsa/pki/crl.pem" \
        "/etc/openvpn"

    if [[ $ENCRYPTION_DH_TYPE == "ECDH" ]]; then
        DH_CONF_STR="dh none"$'\n'"ecdh-curve $ENCRYPTION_ECDH_CURVE"
    elif [[ $ENCRYPTION_DH_TYPE == "DH" ]]; then
        cp dh.pem /etc/openvpn
        DH_CONF_STR="dh dh.pem"
    fi

    # Make cert revocation list readable for non-root
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

    info "Populating the configure file at: /etc/openvpn/server.conf"
    cat >/etc/openvpn/server.conf <<HERE
port $PORT
proto $PROTOCOL
dev tun
user nobody
group $NOGROUP

persist-key
persist-tun
keepalive 10 120

topology subnet

server $SUBNET_ADDRESS 255.255.255.0
ifconfig-pool-persist ipp.txt

push "dhcp-option DNS $DNS_ADDRESS"
push "route $VPC_ADDRESS 255.255.255.0"

$DH_CONF_STR

$TLS_SIG ${TLS_SIG}.key 0
crl-verify crl.pem
ca ca.crt

cert ${SERVER_NAME}.crt
key ${SERVER_NAME}.key

auth $ENCRYPTION_HMAC_ALG
cipher $ENCRYPTION_CIPHER
ncp-ciphers $ENCRYPTION_CIPHER

tls-server
tls-version-min 1.2
tls-cipher $ENCRYPTION_CC_CIPHER

client-config-dir /etc/openvpn/ccd

status /var/log/openvpn/status.log

script-security 2
auth-gen-token 86400
auth-user-pass-verify /etc/openvpn/ovpn-auth via-file

verb 4
HERE

    # Create client-config-dir dir
    mkdir -p /etc/openvpn/ccd

    # Create log dir
    mkdir -p /var/log/openvpn

    # Enable routing
    echo 'net.ipv4.ip_forward=1' >/etc/sysctl.d/20-openvpn.conf

    info "Apply sysctl rules"
    sysctl --system

    info "Checking if SELinux is enabled"
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

    systemctl daemon-reload
    systemctl enable openvpn@server
    systemctl restart openvpn@server
}

function configure_iptables() {
    info "Move the iptables-save template to its final location"
    mkdir -p "/etc/iptables"
    mv "$PROVISIONER_FILES/etc/iptables/iptables-rules.v4" \
        "/etc/iptables/iptables-rules.v4"

    info "Modify the iptables file to reflect the correct network adapter"
    sed --in-place \
        -e "s/<<PRIVATE_ETHERNET_INTERFACE>>/$IPTABLES_PRIVATE_ETHERNET_INTERFACE/g" \
        -e "s/<<PUBLIC_ETHERNET_INTERFACE>>/$IPTABLES_PUBLIC_ETHERNET_INTERFACE/g" \
        -e "s/<<OPENVPN_CUSTOM_PROTOCOL>>/$IPTABLES_OPENVPN_CUSTOM_PROTOCOL/g" \
        -e "s/<<OPENVPN_CUSTOM_PORT>>/$IPTABLES_OPENVPN_CUSTOM_PORT/g" \
        -e "s/<<OPENVPN_SUBNET_ADDRESS>>/$IPTABLES_OPENVPN_SUBNET_ADDRESS/g" \
        "/etc/iptables/iptables-rules.v4"

    info "Move the systemd service file to its final location"
    mv "$PROVISIONER_FILES/etc/systemd/system/iptables-activation.service" \
        "/etc/systemd/system/iptables-activation.service"

    info "Enable the custom systemd service and apply rules"
    systemctl daemon-reload
    systemctl enable iptables-activation
    systemctl restart iptables-activation
}

# Usage: $( create_client_configuration my_client_1 )
function create_client_configuration() {

    info "starting to create client configuration for client: $1"

    export CLIENT_NAME_ON_FILE="$1"

    #
    # Creating the client
    #

    info "checking if a config with that name already exists"

    if [[ $(tail -n +2 /etc/openvpn/easy-rsa/pki/index.txt | grep -c -E "/CN=$CLIENT_NAME_ON_FILE\$") == "1" ]]; then
        error "Client '$CLIENT_NAME_ON_FILE' already exists."
    fi

    info "calling easy-rsa for build-client-full"

    cd /etc/openvpn/easy-rsa/
    ./easyrsa build-client-full "$CLIENT_NAME_ON_FILE" nopass

    #
    # Creating the configuration file
    #

    if [[ $PROTOCOL_CHOICE =~ "udp" ]]; then
        PROTOCOL_CONF_STR="proto udp"$'\n'"explicit-exit-notify"
    else
        PROTOCOL_CONF_STR="proto tcp-client"
    fi

    if [[ -n "$SUDO_USER" ]]; then
        TARGET_DIR="/home/$SUDO_USER"
    else
        TARGET_DIR="/root"
    fi

    info "creating config file at: ${TARGET_DIR}/${CLIENT_NAME_ON_FILE}.ovpn"

    cat >"${TARGET_DIR}/${CLIENT_NAME_ON_FILE}.ovpn" <<EOF
client
$PROTOCOL_CONF_STR
remote $PUBLIC_IP $PORT
dev tun

resolv-retry infinite
nobind
persist-key
persist-tun

remote-cert-tls server
verify-x509-name $SERVER_NAME name

auth $ENCRYPTION_HMAC_ALG
auth-nocache

cipher $ENCRYPTION_CIPHER
tls-client
tls-version-min 1.2
tls-cipher $ENCRYPTION_CC_CIPHER

ignore-unknown-option block-outside-dns
setenv opt block-outside-dns # Prevent Windows 10 DNS leak

auth-user-pass

verb 4

<ca>
$(cat "/etc/openvpn/easy-rsa/pki/ca.crt")
</ca>

<cert>
$(awk '/BEGIN/,/END/' "/etc/openvpn/easy-rsa/pki/issued/${CLIENT_NAME}.crt")
</cert>

<key>
$(cat "/etc/openvpn/easy-rsa/pki/private/${CLIENT_NAME}.key")
</key>

$(if [[ $TLS_SIG == "tls-auth" ]]; then echo "key-direction 1"; fi)
<${ENCRYPTION_TLS_SIG}>
$(cat "/etc/openvpn/${ENCRYPTION_TLS_SIG}.key")
</${ENCRYPTION_TLS_SIG}>
EOF
}

function create_all_client_configurations() {
    for CLIENT_NAME in "$@"; do
        create_client_configuration "$CLIENT_NAME"
    done
}

function check_tun_availability() {
    if [ ! -e /dev/net/tun ]; then error "TUN is not available at /dev/net/tun"; fi
}

# ---------------------------------------------------------------------------- #
# Main
# ---------------------------------------------------------------------------- #

execute_task assert_sudo
execute_task check_tun_availability
execute_task wait_cloud_init

execute_task_in_golden_image install_openvpn
execute_task_in_golden_image install_argon2
execute_task_in_golden_image install_easy_rsa

execute_task_in_deployment configure_easy_rsa
execute_task_in_deployment configure_openvpn

execute_task_in_golden_image install_ovpn_auth

execute_task_in_deployment create_secrets_file_for_ovpn_auth
execute_task_in_deployment configure_iptables

execute_task_in_golden_image install_unbound

execute_task_in_deployment configure_unbound
execute_task_in_deployment create_all_client_configurations "$@"
