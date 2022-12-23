package cache

import (
	"go/ast"
	models "tde/models/in_program_models"
)

// type Cache[T any] interface {
// 	Set(models.CandidateID, T)
// 	Get(models.CandidateID) (T, bool)
// 	Delete(models.CandidateID)
// }

type ASTCache struct {
	store map[models.CandidateID]*ast.Node
}

func (c ASTCache) Set(id models.CandidateID, n *ast.Node) {
	c.store[id] = n
}

func (c ASTCache) Get(id models.CandidateID) (*ast.Node, bool) {
	if n, ok := c.store[id]; ok {
		return n, true
	} else {
		return nil, false
	}
}

func (c ASTCache) Delete(id models.CandidateID) {
	delete(c.store, id)
}