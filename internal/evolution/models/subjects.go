package models

import (
	"maps"
	"slices"

	"tde/internal/utilities/mapw"
)

type Subjects map[Sid]*Subject // To make Subjects accessible by CIDs

func (s Subjects) Add(subj *Subject) {
	(s)[subj.Sid] = subj
}

func (s Subjects) Join(s2 Subjects) {
	maps.Copy(s, s2)
}

func (s Subjects) Diff(subtract Subjects) Subjects {
	diff := mapw.Diff(map[Sid]*Subject(s), map[Sid]*Subject(subtract))
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
	return slices.Collect(maps.Values(s))
}

func (s Subjects) Keys() []Sid {
	return slices.Collect(maps.Keys(s))
}
