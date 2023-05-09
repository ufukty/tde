dev: dev-up dev-deploy

dev-up:
	cd 1-provisioning && terraform apply

dev-down:
	cd 1-provisioning && terraform destroy
	
dev-deploy:
	@echo "Activate virtual environnment and ssh-agent before run ansible"
	cd 2-deployment && ansible-playbook --forks 20 playbook.yml

dev-redeploy:
    @echo "Activate virtual environnment and ssh-agent before run ansible"
	cd platform/2-deployment && with_echo ansible-playbook --forks=20 --tags=redeploy playbook.yml