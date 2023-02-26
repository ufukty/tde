package main

import "net/http"

// TODO: accepts uploaded files
// TODO: , puts in sandboxed directory
// TODO: , checks against maliciousness
// TODO: , transforms into AST
// TODO: uploads AST to evolver server when requested

func main() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}
