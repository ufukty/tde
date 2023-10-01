package pool

import models "tde/models/program"

type Pool struct {
	p    models.Subjects
	subs []*Pool
}

func (p *Pool) Reproduce(id models.Sid) {}

func New() *Pool {
	return &Pool{
		p: make(models.Subjects),
	}
}

func (p *Pool) All() models.Subjects {
	return p.p
}

func (p *Pool) FilterValidIn(layer models.Layer) models.Subjects {
	results := models.Subjects{}
	for sid, cand := range p.p {
		if cand.IsValidIn(layer) {
			results[sid] = cand
		}
	}
	return results
}

func (p *Pool) Sub() *Pool {
	sub := &Pool{}
	p.subs = append(p.subs, sub)
	return sub
}
