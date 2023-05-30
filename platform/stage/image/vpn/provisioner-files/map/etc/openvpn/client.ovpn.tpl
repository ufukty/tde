client
{{PROTOCOL_CONF_STR}}
remote {{PUBLIC_IP}} {{OPENVPN_PORT}}
dev tun

resolv-retry infinite
nobind
persist-key
persist-tun

remote-cert-tls server
verify-x509-name {{EASYRSA_SERVER_NAME}} name

auth {{ENCRYPTION_HMAC_ALG}}
auth-nocache

cipher {{ENCRYPTION_CIPHER}}
tls-client
tls-version-min 1.2
tls-cipher {{ENCRYPTION_CC_CIPHER}}

ignore-unknown-option block-outside-dns
setenv opt block-outside-dns # Prevent Windows 10 DNS leak

auth-user-pass

verb 4

<ca>
{{EASYRSA_CA_CERT_CONTENT}}
</ca>

<cert>
{{EASYRSA_CLIENT_CERT_CONTENT}}
</cert>

<key>
{{EASYRSA_CLIENT_KEY_CONTENT}}
</key>

{{TLS_AUTH_KEY_DIRECTION}}
<{{TLS_SIG}}>
{{TLS_SIG_KEY_CONTENT}}
</{{TLS_SIG}}>
