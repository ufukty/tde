note: 
	@echo "\033]34mDon't forget: ssh-agent; ssh-add; venv\033]0m"

base: note
	cd base && packer build base.pkr.hcl

internal: note
	cd internal && packer build internal.pkr.hcl

vpn: note
	cd vpn && packer build vpn.pkr.hcl