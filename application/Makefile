test-word-reverse:
	go run -tags="tde" tde/examples/word-reverse/word_reverse/tde

initial-environment-setup:
	brew update && brew install \
		clang-format
	
	go install golang.org/x/tools/cmd/stringer

	go mod tidy
