package list_test

import (
	"log"
	astw_util "tde/internal/astw/astwutl"
	"tde/internal/folders/discovery"
	"tde/internal/utilities"

	"errors"
	"fmt"
	"go/ast"
	"go/scanner"
	"path/filepath"
	"reflect"
)

type Command struct {
	Root    string `precedence:"0"`
	Verbose bool   `short:"v"`
}

func (c *Command) validateArgs() {
	if c.Root == "" {
		var err error
		c.Root, err = discovery.ModuleRoot()
		if err != nil {
			log.Fatalln("Could not detect module root:", err)
		}
	}
}

func extractTargetFunction(callExp *ast.CallExpr) (funcName string, receiverName string, err error) {
	switch fun := callExp.Fun.(type) {
	case *ast.Ident:
		funcName = fun.Name
	case *ast.SelectorExpr:
		if ident, ok := fun.X.(*ast.Ident); ok {
			receiverName = ident.Name
		}
		funcName = fun.Sel.Name
	default:
		return "", "", errors.New("unexpected type of call expression: " + reflect.TypeOf(fun).String())
	}
	return
}

func PrintModule(name, path string) {
	headerPrefix := "\033[1;33m"
	headerSuffix := "\033[1;0m"
	fmt.Printf(". %s%-19s%s %s\n",
		headerPrefix, name, headerSuffix,
		path,
	)
}

func PrintTDEFunction(path string, line int, function string, last bool) {
	headerPrefix := "\033[1;34m"
	headerSuffix := "\033[1;0m"
	// symbolPrefix := "\033[1;1m"
	// symbolSuffix := "\033[1;0m"
	treeInd := ""
	if last {
		treeInd += "└─"
	} else {
		treeInd += "├─"
	}
	fmt.Printf("%s %s%-18s%s %s:%d\n",
		treeInd,
		headerPrefix, function, headerSuffix,
		path, line,
	)
}

func printTestFunc(path string, line int, function string, lastInLevel bool, contFirstLevel bool) {
	headerPrefix := "\033[1;35m"
	headerSuffix := "\033[1;0m"
	treeInd := ""
	if contFirstLevel {
		treeInd += "│  "
	} else {
		treeInd += "   "
	}
	if lastInLevel {
		treeInd += "└─"
	} else {
		treeInd += "├─"
	}
	fmt.Printf("%s %s%-15s%s %s:%d\n",
		treeInd,
		headerPrefix, function, headerSuffix,
		path, line,
	)
}

func (c *Command) Run() {
	c.validateArgs()

	root, err := discovery.ModuleRoot()
	if err != nil {
		log.Fatalln("Failed to detect module root path.")
	}
	PrintModule(filepath.Base(root), filepath.Dir(root))

	tests, skipped, err := discovery.TestFunctionsInSubdirs(root)
	if err != nil {
		log.Fatalln("Failed to detect test functions:", err)
	}

	for i, test := range tests {
		rel, err := filepath.Rel(root, test.Path)
		if err != nil {
			rel = test.Path
		}

		PrintTDEFunction(rel, test.LineStart, test.Name, i == len(tests)-1)

		type Call struct {
			RelativePath  string
			LineNumber    int
			FunctionPrint string
		}
		calls := utilities.FilteredMap(test.Calls, func(i int, value *ast.CallExpr) (*Call, bool) {
			functionPrint, err := astw_util.String(value.Fun)
			if err != nil {
				return nil, false
			}

			funcName, _, err := extractTargetFunction(value)
			if err != nil {
				// fmt.Printf("\tFailed to detect target function name and its receiver\n")
				return nil, false
			}

			implFuncFile, implFuncName := discovery.ExpectedTargetFileAndFuncNameFor(test.Path, funcName)

			implFuncDetails, err := discovery.TargetFunctionInFile(implFuncFile, implFuncName)
			if err != nil {
				// fmt.Printf("\t%s(...) File not found in \"%s\": %e\n", funcName, targetFuncFile, err)
				return nil, false
			}

			rel, err := filepath.Rel(root, implFuncDetails.Path)
			if err != nil {
				rel = implFuncDetails.Path
			}

			return &Call{
				RelativePath:  rel,
				LineNumber:    implFuncDetails.LineStart,
				FunctionPrint: functionPrint,
			}, true
		})

		for j, call := range calls {
			printTestFunc(call.RelativePath, call.LineNumber, call.FunctionPrint, j == len(calls)-1, i < len(tests)-1)
		}
	}

	if len(skipped) > 0 {
		if !c.Verbose {
			fmt.Println("Could not inspect some of the files because of syntax errors. Run with -v flag to print details.")
		} else {
			fmt.Println("Could not inspect some of the files because of syntax errors.")
			for file, err := range skipped {
				if err, ok := err.(scanner.ErrorList); ok {
					fmt.Printf("\033[31m%s:\033[0m\n", file)
					for _, err := range err {
						fmt.Println("    ", err)
					}
				} else {
					fmt.Printf("\033[31m%s:\033[0m %s\n", file, err)
				}
			}
		}
	}
}
