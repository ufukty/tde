all: ast-inspect client evolver-server runner-server customs

ast-inspect:
	go build -o build/ast-inspect ./cmd/ast-inspect

client:
	go build -o build/client ./cmd/client

evolver-server:
	go build -o build/evolver-server ./cmd/evolver-server

customs:
	go build -o build/customs ./cmd/customs

runner-server:
	go build -o build/runner-server ./cmd/runner-server

test-word-reverse:
	go run -tags="tde" tde/examples/word-reverse/word_reverse/tde

initial-environment-setup:
	go install golang.org/x/tools/cmd/stringer