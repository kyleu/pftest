// Content managed by Project Forge, see [projectforge.md] for details.
package timestamp

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Timestamps []*Timestamp

func (t Timestamps) Get(id string) *Timestamp {
	return lo.FindOrElse(t, nil, func(x *Timestamp) bool {
		return x.ID == id
	})
}

func (t Timestamps) GetByIDs(ids ...string) Timestamps {
	return lo.Filter(t, func(x *Timestamp, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (t Timestamps) IDs() []string {
	return lo.Map(t, func(x *Timestamp, _ int) string {
		return x.ID
	})
}

func (t Timestamps) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(t, func(x *Timestamp, _ int) {
		ret = append(ret, x.ID)
	})
	return ret
}

func (t Timestamps) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(t, func(x *Timestamp, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (t Timestamps) Clone() Timestamps {
	return slices.Clone(t)
}
