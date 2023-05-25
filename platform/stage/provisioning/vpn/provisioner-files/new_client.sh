#!/bin/bash

# ---------------------------------------------------------------------------- #
# Required Environment Variables
# ---------------------------------------------------------------------------- #

PUBLIC_IP="${PUBLIC_IP:-"PUBLIC_IP is required."}"
CLIENT_NAME="${CLIENT_NAME:-"CLIENT_NAME is required."}"

# ---------------------------------------------------------------------------- #
# Optional Environment Variables
# ---------------------------------------------------------------------------- #

# ENCRYPTION_CERT_TYPE valid values:
# [ ECDSA, RSA ]
ENCRYPTION_CERT_TYPE="${ENCRYPTION_CERT_TYPE:-"ECDSA"}"

# !!! Only usable when ENCRYPTION_CERT_TYPE="ECDSA"
# ENCRYPTION_ECDSA_CC_CIPHER valid values:
# [ ECDHE-ECDSA-AES-128-GCM-SHA256, ECDHE-ECDSA-AES-256-GCM-SHA384 ]
ENCRYPTION_ECDSA_CC_CIPHER="${ENCRYPTION_ECDSA_CC_CIPHER:-"TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256"}"

# !!! Only usable when ENCRYPTION_CERT_TYPE="RSA"
# ENCRYPTION_RSA_CC_CIPHER valid values:
# [ ECDHE-RSA-AES-128-GCM-SHA256, ECDHE-RSA-AES-256-GCM-SHA384 ]
ENCRYPTION_RSA_CC_CIPHER="${ENCRYPTION_RSA_CC_CIPHER:-"ECDHE-RSA-AES-128-GCM-SHA256"}"

# ENCRYPTION_CIPHER valid values:
# [ AES-128-GCM, AES-192-GCM, AES-256-GCM, AES-128-CBC, AES-192-CBC, AES-256-CBC ]
ENCRYPTION_CIPHER="${ENCRYPTION_CIPHER:-"AES-128-GCM"}"

# ENCRYPTION_HMAC_ALG valid values:
# [ SHA-256, SHA-384, SHA-512 ]
# - When GCM type ciphers are used, the algorithm is used only for
#   encryption of tls-auth packets from the control channel.
# - If, CBC type ciphers are used, the algorithm is used in addition
#   for authenticates data channel packets too.
ENCRYPTION_HMAC_ALG="${ENCRYPTION_HMAC_ALG:-"SHA256"}"

# OPENVPN_PROTOCOL valid values:
# [ udp, tcp ]
OPENVPN_PROTOCOL="${OPENVPN_PROTOCOL:-"tcp"}"
OPENVPN_PORT="${OPENVPN_PORT:-"443"}"

# TLS_SIG valid values:
# [ tls-crypt, tls-auth ]
# - Those will add additional layer of security to the control channel.
# - tls-auth authenticates the packets, while tls-crypt authenticate
#   and encrypt them.
TLS_SIG="${TLS_SIG:-"tls-crypt"}"

# ---------------------------------------------------------------------------- #
# Runtime Variables
# ---------------------------------------------------------------------------- #

EASYRSA_SERVER_NAME="$(cat /etc/openvpn/easy-rsa/generated/server_name)"

if [[ $OPENVPN_PROTOCOL =~ "udp" ]]; then
    PROTOCOL_CONF_STR="proto udp"$'\n'"explicit-exit-notify"
else
    PROTOCOL_CONF_STR="proto tcp-client"
fi

if test "$TLS_SIG" == "tls-crypt"; then
    TLS_AUTH_KEY_DIRECTION="key-direction 1"
else
    TLS_AUTH_KEY_DIRECTION="#"
fi

if [[ $ENCRYPTION_CERT_TYPE == "ECDSA" ]]; then
    ENCRYPTION_CC_CIPHER="$ENCRYPTION_ECDSA_CC_CIPHER"
elif [[ $ENCRYPTION_CERT_TYPE == "RSA" ]]; then
    ENCRYPTION_CC_CIPHER="$ENCRYPTION_RSA_CC_CIPHER"
fi

# ---------------------------------------------------------------------------- #
# Key generation
# ---------------------------------------------------------------------------- #

cd /etc/openvpn/easy-rsa/
./easyrsa build-client-full mbp nopass

# ---------------------------------------------------------------------------- #
# Templating
# ---------------------------------------------------------------------------- #

EASYRSA_CA_CERT_CONTENT="$(cat "/etc/openvpn/easy-rsa/pki/ca.crt")"
EASYRSA_CLIENT_KEY_CONTENT="$(cat "/etc/openvpn/easy-rsa/pki/private/$CLIENT_NAME.key")"
EASYRSA_CLIENT_CERT_CONTENT="$(awk '/BEGIN/,/END/' "/etc/openvpn/easy-rsa/pki/issued/$CLIENT_NAME.crt")"
TLS_SIG_KEY_CONTENT="$(cat "/etc/openvpn/$TLS_SIG.key")"

echo "templating the /etc/openvpn/client.ovpn.tpl for client '$CLIENT_NAME' with"
echo "EASYRSA_CA_CERT_CONTENT      = $EASYRSA_CA_CERT_CONTENT"
echo "EASYRSA_CLIENT_CERT_CONTENT  = $EASYRSA_CLIENT_CERT_CONTENT"
echo "EASYRSA_CLIENT_KEY_CONTENT   = $EASYRSA_CLIENT_KEY_CONTENT"
echo "EASYRSA_SERVER_NAME          = $EASYRSA_SERVER_NAME"
echo "ENCRYPTION_CC_CIPHER         = $ENCRYPTION_CC_CIPHER"
echo "ENCRYPTION_CIPHER            = $ENCRYPTION_CIPHER"
echo "ENCRYPTION_HMAC_ALG          = $ENCRYPTION_HMAC_ALG"
echo "OPENVPN_PORT                 = $OPENVPN_PORT"
echo "PROTOCOL_CONF_STR            = $PROTOCOL_CONF_STR"
echo "PUBLIC_IP                    = $PUBLIC_IP"
echo "TLS_AUTH_KEY_DIRECTION       = $TLS_AUTH_KEY_DIRECTION"
echo "TLS_SIG                      = $TLS_SIG"
echo "TLS_SIG_KEY_CONTENT          = $TLS_SIG_KEY_CONTENT"

cat /etc/openvpn/client.ovpn.tpl | sed \
    -e "s;{{EASYRSA_CA_CERT_CONTENT}};$EASYRSA_CA_CERT_CONTENT;" \
    -e "s;{{EASYRSA_CLIENT_CERT_CONTENT}};$EASYRSA_CLIENT_CERT_CONTENT;" \
    -e "s;{{EASYRSA_CLIENT_KEY_CONTENT}};$EASYRSA_CLIENT_KEY_CONTENT;" \
    -e "s;{{EASYRSA_SERVER_NAME}};$EASYRSA_SERVER_NAME;" \
    -e "s;{{ENCRYPTION_CC_CIPHER}};$ENCRYPTION_CC_CIPHER;" \
    -e "s;{{ENCRYPTION_CIPHER}};$ENCRYPTION_CIPHER;" \
    -e "s;{{ENCRYPTION_HMAC_ALG}};$ENCRYPTION_HMAC_ALG;" \
    -e "s;{{OPENVPN_PORT}};$OPENVPN_PORT;" \
    -e "s;{{PROTOCOL_CONF_STR}};$PROTOCOL_CONF_STR;" \
    -e "s;{{PUBLIC_IP}};$PUBLIC_IP;" \
    -e "s;{{TLS_AUTH_KEY_DIRECTION}};$TLS_AUTH_KEY_DIRECTION;" \
    -e "s;{{TLS_SIG}};$TLS_SIG;" \
    -e "s;{{TLS_SIG_KEY_CONTENT}};$TLS_SIG_KEY_CONTENT;" \
    >"/home/$USER_ACCOUNT_NAME/$CLIENT_NAME.ovpn"
