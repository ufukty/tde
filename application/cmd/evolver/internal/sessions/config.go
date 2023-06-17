package sessions

import "go/ast"

type Config struct {
	Timeout    int      // in seconds
	Runner     string   // ip address
	Continue   string   // session
	Model      string   //
	Ratios     string   //
	Population int      //
	Iterate    int      //
	Size       int      //
	Package    []string // packages allowed to import
	TestName   string
}

type (
	Runner struct {
		Address string `json:"address"`
		Port    string `json:"port"`
		Token   string `json:"token"`
	}

	Target struct {
		// TODO: Module          map[string]*ast.Package // import path -> package
		Package  *ast.Package  `json:"package"`
		File     *ast.File     `json:"file"`
		FuncDecl *ast.FuncDecl `json:"func_decl"`
	}

	Probabilities struct {
		CrossOver float64 `json:"cross_over"`
		Mutation  float64 `json:"mutation"`
	}

	Config2 struct {
		Timeout int    `json:"timeout"` // in seconds
		Runner  string `json:"runner"`  // ip address

		// Model         string        `json:"model"`
		Probabilities Probabilities `json:"probabilities"`

		Population      int      `json:"population"`
		SizeLimit       int      `json:"size_limit"`
		AllowedPackages []string `json:"allowed_packages"` // packages allowed to import

		Iterate  int    `json:"iterate"`
		TestName string `json:"test_name"`
	}
)
