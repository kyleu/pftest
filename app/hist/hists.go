// Package hist - Content managed by Project Forge, see [projectforge.md] for details.
package hist

import (
	"slices"

	"github.com/samber/lo"
)

type Hists []*Hist

func (h Hists) Get(id string) *Hist {
	return lo.FindOrElse(h, nil, func(x *Hist) bool {
		return x.ID == id
	})
}

func (h Hists) GetByIDs(ids ...string) Hists {
	return lo.Filter(h, func(x *Hist, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (h Hists) IDs() []string {
	return lo.Map(h, func(x *Hist, _ int) string {
		return x.ID
	})
}

func (h Hists) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(h)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(h, func(x *Hist, _ int) {
		ret = append(ret, x.ID)
	})
	return ret
}

func (h Hists) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(h)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(h, func(x *Hist, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (h Hists) Clone() Hists {
	return slices.Clone(h)
}
