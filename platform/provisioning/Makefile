dev-up:
	cd terraform &&	terraform apply

dev-down:
	cd terraform &&	terraform destroy
	
dev-conf:
	cd ansible && ansible-playbook playbook.yml

dev: dev-up dev-conf