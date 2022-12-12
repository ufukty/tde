ast-inspect:
	go build -o build/ast-inspect GoGP/agent/cmd/ast-inspect

agent:
	go build -o build/agent GoGP/agent/cmd/server

