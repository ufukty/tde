package pool

import models "tde/models/program"

type Pool struct {
	p      models.Subjects
	subs   []*Pool
	Depths map[models.Sid]int
}

func (p *Pool) Set(subj *models.Subject) {
	p.p[subj.Sid] = subj
	p.Depths[subj.Sid] = p.Depths[subj.Parent]
}

func (p *Pool) Delete(sid models.Sid) {
	delete(p.p, sid)
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

func (p *Pool) FilterByDepth(le int) models.Subjects {
	results := models.Subjects{}
	for sid, depth := range p.Depths {
		if depth <= le {
			results[sid] = p.p[sid]
		}
	}
	return results
}

func New(root models.Sid) *Pool {
	return &Pool{
		p:      make(models.Subjects),
		Depths: map[models.Sid]int{root: 0},
	}
}
