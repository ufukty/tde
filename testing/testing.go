package testing

type Run struct {
	totalChecks       int
	totalFailedChecks int
}

type Objective struct {
	Tag  string
	Kind int
	Runs []Run
}

func (r *Run) Assert(left, right int) {
	r.totalChecks += 1
	
}

func (o *Objective) NewRun(test func(r *Run)) {
	run := Run{0, 0}
	o.Runs = append(o.Runs, run)
	test(&run)
}

type Candidate struct {
	
}

type Testing struct {
	Objectives []Objective
}

func CreateTesting() *Testing {
	return &Testing{}
}

func (t *Testing) NewObjective(tag string, test func(o *Objective)) {
	objective := Objective{
		Tag: tag,
	}
	t.Objectives = append(t.Objectives, objective)
	test(&objective)
}


func (t *Testing) Candidate(candidateTesting func(c *Candidate) ) {
	for id, ind := range t.
})