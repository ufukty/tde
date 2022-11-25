package testing

type Run struct {
	expected
}

const (
	HitOrMiss           = iota
	MeasurableErrorRate = iota
)

type Objective struct {
	Tag  string
	Kind int
	Runs []Run
}

func (r *Run) Assert(left, right int) {
	r.TotalCalls += 1
	if o.Kind == HitOrMiss {
		//
	}
}

func (o *Objective) NewRun(test func(run *Run)) {
	run := Run{
		Tag:  tag,
		Kind: kind,
	}
	o.Runs = append(o.Runs, run)
	test(&run)
}

// type TestFunction func(t *Testing)

type Testing struct {
	Objectives []Objective
}

func CreateTesting() *Testing {
	return &Testing{}
}

func (t *Testing) NewObjective(tag string, isErrorRateMeasurable bool, test func(objective *Objective)) {
	var kind int
	if isErrorRateMeasurable {
		kind = MeasurableErrorRate
	} else {
		kind = HitOrMiss
	}

	objective := Objective{
		Tag:  tag,
		Kind: kind,
	}
	t.Objectives = append(t.Objectives, objective)
	test(&objective)
}
