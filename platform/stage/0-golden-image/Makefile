SHELL = /usr/local/bin/bash
.PHONY: *

note-ssh-agent := echo -e "\033[30m\033[43m\033[1m ssh key isn't decrypted \033[0m"

check := \
	(ssh-add -l | grep mbp-ed > /dev/null) || ($(note-ssh-agent) && ssh-agent && ssh-add) &&\
	which ansible > /dev/null || . ~/venv/bin/activate &&\
	test -n "$DIGITALOCEAN_TOKEN" || . ../.env

base: 
	@$(check) && cd base && echo packer build base.pkr.hcl

internal: 
	@$(check) && cd internal && packer build internal.pkr.hcl

vpn: 
	@$(check) && cd vpn && packer build vpn.pkr.hcl