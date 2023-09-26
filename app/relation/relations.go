// Package relation - Content managed by Project Forge, see [projectforge.md] for details.
package relation

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Relations []*Relation

func (r Relations) Get(id uuid.UUID) *Relation {
	return lo.FindOrElse(r, nil, func(x *Relation) bool {
		return x.ID == id
	})
}

func (r Relations) GetByIDs(ids ...uuid.UUID) Relations {
	return lo.Filter(r, func(x *Relation, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (r Relations) IDs() []uuid.UUID {
	return lo.Map(r, func(x *Relation, _ int) uuid.UUID {
		return x.ID
	})
}

func (r Relations) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *Relation, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (r Relations) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(r, func(x *Relation, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (r Relations) Clone() Relations {
	return slices.Clone(r)
}
