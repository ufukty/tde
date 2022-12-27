all: ast-inspect client server

ast-inspect:
	go build -o build/ast-inspect ./cmd/ast-inspect

client:
	go build -o build/agent ./cmd/client

server:
	go build -o build/server ./cmd/server

test-word-reverse:
	go run -tags="tde" tde/examples/word-reverse/word_reverse/tde