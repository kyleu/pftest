// Content managed by Project Forge, see [projectforge.md] for details.
package group

import (
	"slices"

	"github.com/samber/lo"
)

type Groups []*Group

func (g Groups) Get(id string) *Group {
	return lo.FindOrElse(g, nil, func(x *Group) bool {
		return x.ID == id
	})
}

func (g Groups) GetByIDs(ids ...string) Groups {
	return lo.Filter(g, func(x *Group, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (g Groups) IDs() []string {
	return lo.Map(g, func(x *Group, _ int) string {
		return x.ID
	})
}

func (g Groups) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(g)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(g, func(x *Group, _ int) {
		ret = append(ret, x.ID)
	})
	return ret
}

func (g Groups) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(g)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(g, func(x *Group, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (g Groups) Clone() Groups {
	return slices.Clone(g)
}
