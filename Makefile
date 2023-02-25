all: ast-inspect client evolver-server runner-server

ast-inspect:
	go build -o build/ast-inspect ./cmd/ast-inspect

client:
	go build -o build/agent ./cmd/client

evolver-server:
	go build -o build/evolver-server ./cmd/evolver-server

runner-server:
	go build -o build/runner-server ./cmd/runner-server

test-word-reverse:
	go run -tags="tde" tde/examples/word-reverse/word_reverse/tde

initial-environment-setup:
	go install golang.org/x/tools/cmd/stringer