package evolution

import (
	"fmt"
	"tde/internal/evolution/models"
	"tde/internal/evolution/pool"
)

// Searches for a AST that compile & run
type programSearch struct {
	*commons
	Subpool *pool.Pool
	Src     models.Sid
	Counter int
	Search  *codeSearch
	Dests   models.Sid
}

func (ps *programSearch) IsEnded() bool {
	return ps.Counter == ps.Params.Program.Generations
}

func (ps *programSearch) Iterate() (models.Subjects, error) {
	// collect findings

	if ps.Search == nil {
		// ps.Search = newCodeSearch(ps.commons, root)
	}

	products, err := ps.Search.Iterate()
	if err != nil {
		return nil, fmt.Errorf("iterating program search: %w", err)
	}

	if len(products) > 0 {
		return products, nil
	}

	if ps.Search.IsEnded() {
		ps.Counter++
	}

	return models.Subjects{}, nil
}

func (ps *programSearch) Terminate() {

}

func newProgramSearch(commons *commons, root *models.Subject) *programSearch {
	return &programSearch{
		commons: commons,
		Subpool: commons.Pool.Sub(root),
		Src:     root.Sid,
		Counter: 0,
		Search:  nil,
		Dests:   "",
	}
}
