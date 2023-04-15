PROGRAMS := ast-inspect client evolver-server runner-server customs

all: $(PROGRAMS)

$(PROGRAMS):
	bash commands compile $@

deploy-dev:
	cd platform/provisioning/ansible && ansible-playbook playbook.yml

test-word-reverse:
	go run -tags="tde" tde/examples/word-reverse/word_reverse/tde

initial-environment-setup:
	go install golang.org/x/tools/cmd/stringer
