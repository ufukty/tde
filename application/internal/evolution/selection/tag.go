package selection

import (
	models "tde/models/program"
)

func isIn(sid models.Sid, collection []*models.Subject) bool {
	for _, cand := range collection {
		if sid == cand.Sid {
			return true
		}
	}
	return false
}

// tag first occurence of subjects as to "keep". tag the rest as "mutate"
// TODO: cross-overs?
func Tag(current, next []*models.Subject) (keep, mutate, clear []*models.Subject) {
	for _, cand := range next {
		if isIn(cand.Sid, keep) {
			mutate = append(mutate, cand)
		} else if !isIn(cand.Sid, current) {
			clear = append(clear, cand)
		} else {
			keep = append(keep, cand)
		}
	}
	return
}
