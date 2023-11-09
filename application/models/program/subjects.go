package models

import (
	"tde/internal/utilities"

	"golang.org/x/exp/maps"
)

type Subjects map[Sid]*Subject // To make Subjects accessible by CIDs

func (s Subjects) Add(subj *Subject) {
	(s)[subj.Sid] = subj
}

func (s Subjects) Join(s2 Subjects) {
	maps.Copy(s, s2)
}

func (s Subjects) Diff(subtract Subjects) Subjects {
	diff := utilities.MapDiff(map[Sid]*Subject(s), map[Sid]*Subject(subtract))
	return (Subjects)(diff)
}

func SubjectsFrom(s []*Subject) Subjects {
	ss := Subjects{}
	for _, i := range s {
		ss[i.Sid] = i
	}
	return ss
}

func (s Subjects) Values() []*Subject {
	return maps.Values(s)
}

func (s Subjects) Keys() []Sid {
	return maps.Keys(s)
}
