package utilities

func If[N any](cond bool, ifTrue, ifFalse N) N {
	if cond {
		return ifTrue
	} else {
		return ifFalse
	}
}
