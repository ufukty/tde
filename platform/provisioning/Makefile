development:
	cd terraform &&	terraform apply 
	cd ansible && ansible-playbook playbook.yaml

development-with-test: development
	ssh "ufukty@$$(cat ansible/inventory.cfg | grep thesis-evolution-0 | cut -f 2 -d ' ' | cut -f 2 -d '=')" "cat runner_ip_addresses.txt"