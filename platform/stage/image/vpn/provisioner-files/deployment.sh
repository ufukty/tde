#!/bin/bash




assert_sudo
check_tun_availability
wait_cloud_init

configure_easy_rsa
configure_openvpn
create_secrets_file_for_ovpn_auth
configure_iptables
configure_unbound
# create_all_client_configurations "$@"
