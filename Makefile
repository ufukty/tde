PROGRAMS := ast-inspect client evolver-server runner-server customs

all: $(PROGRAMS)

$(PROGRAMS):
	bash commands compile $@

dev-deploy: runner-server evolver-server
	cd platform && make dev-deploy

test-word-reverse:
	go run -tags="tde" tde/examples/word-reverse/word_reverse/tde

# example: make run-client arg1 arg2
.PHONY: $(addprefix run-,$(PROGRAMS))
$(addprefix run-,$(PROGRAMS)): 
	build/$$(bash commands last-build $(subst run-,,$@))/$(subst run-,,$@)-darwin-amd64 $(filter-out $@,$(MAKECMDGOALS))

initial-environment-setup:
	go install golang.org/x/tools/cmd/stringer
	brew update && brew install protobuf bufbuild/buf/buf clang-format
	go get -d github.com/envoyproxy/protoc-gen-validate; go install github.com/envoyproxy/protoc-gen-validate
	
	# Requirements of: https://github.com/grpc-ecosystem/grpc-gateway
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc

	go mod tidy

build-api:
	cd api && protoc -I . -I vendor --go_out=build *.proto