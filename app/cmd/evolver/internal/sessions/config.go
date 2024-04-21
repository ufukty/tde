package sessions

import (
	"go/ast"
	"time"
)

type (
	Runner struct {
		Address string `json:"address"` // ip:port
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

	Config struct {
		Timeout time.Duration `json:"timeout"` // in seconds
		Runner  Runner        `json:"runner"`  // ip address

		Probabilities   Probabilities `json:"probabilities"`
		Population      int           `json:"population"`
		SizeLimit       int           `json:"size_limit"`
		AllowedPackages []string      `json:"allowed_packages"` // packages allowed to import

		Iterate  int    `json:"iterate"`
		TestName string `json:"test_name"`
	}
)
