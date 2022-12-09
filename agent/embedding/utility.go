package embedding

import (
	"fmt"
	"go/parser"

	"github.com/pkg/errors"
)

func StampFunctionPrototype(original, stamp string) (string, error) {
	originalClosedCurlyBracet := fmt.Sprintf("%s}", original)
	ast_, err := parser.ParseExpr(originalClosedCurlyBracet)
	if err != nil {
		fmt.Printf("%+#v", ast_)
		return "", errors.Wrap(err, "Could not parse original function")
	}

	fmt.Println(ast_)
	return "", nil
}
