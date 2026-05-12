package testing

// these only exist to comply with testing library interface.
// doing nothing for Fatalx functions is especially important;
// we don't want to early-exit tests,
// to collect every assertion result

func (t *T) Log(args ...any) {
	// do nothing
}

func (t *T) Logf(format string, args ...any) {
	// do nothing
}

func (t *T) Error(args ...any) {
	// do nothing
}

func (t *T) Errorf(format string, args ...any) {
	// do nothing
}

func (t *T) Fatal(args ...any) {
	// do nothing
}

func (t *T) Fatalf(format string, args ...any) {
	// do nothing
}
