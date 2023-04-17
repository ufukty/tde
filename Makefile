PROGRAMS := ast-inspect client evolver-server runner-server customs

all: $(PROGRAMS)

$(PROGRAMS):
	bash commands compile $@

dev-deploy: runner-server evolver-server
	cd platform && make dev-deploy

test-word-reverse:
	go run -tags="tde" tde/examples/word-reverse/word_reverse/tde

initial-environment-setup:
	go install golang.org/x/tools/cmd/stringer

# example: make run-client arg1 arg2
.PHONY: $(addprefix run-,$(PROGRAMS))
$(addprefix run-,$(PROGRAMS)): 
	build/$$(bash commands last-build $(subst run-,,$@))/$(subst run-,,$@)-darwin-amd64 $(filter-out $@,$(MAKECMDGOALS))