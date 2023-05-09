stage: stage-up stage-deploy

stage-up:
	cd 1-provisioning && terraform apply

stage-down:
	cd 1-provisioning && terraform destroy
	
stage-deploy:
	@echo "Activate virtual environnment and ssh-agent before run ansible"
	cd 2-deployment && ansible-playbook --forks 20 playbook.yml

stage-redeploy:
	@echo "Activate virtual environnment and ssh-agent before run ansible"
	cd 2-deployment && with_echo ansible-playbook --forks=20 --tags=redeploy playbook.yml