test-word-reverse:
	go run -tags="tde" tde/examples/word-reverse/word_reverse/tde

initial-environment-setup:
	brew update && brew install \
		protobuf bufbuild/buf/buf clang-format bazel
	
	go install golang.org/x/tools/cmd/stringer

	go get -d github.com/envoyproxy/protoc-gen-validate; go install github.com/envoyproxy/protoc-gen-validate
	
	# Requirements of: https://github.com/grpc-ecosystem/grpc-gateway
	go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc

	go mod tidy

buf-generate-deps:
	cd buf && curl
	@echo "This might need to run twice"
	cd buf && bazel build //...

buf-generate:
	cd buf && buf generate
	go mod tidy

buf-update:
	cd buf && buf mod update