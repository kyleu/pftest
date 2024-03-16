// Package softdel - Content managed by Project Forge, see [projectforge.md] for details.
package softdel

import (
	"slices"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

type Softdels []*Softdel

func (s Softdels) Get(id string) *Softdel {
	return lo.FindOrElse(s, nil, func(x *Softdel) bool {
		return x.ID == id
	})
}

func (s Softdels) IDs() []string {
	return lo.Map(s, func(xx *Softdel, _ int) string {
		return xx.ID
	})
}

func (s Softdels) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *Softdel, _ int) {
		ret = append(ret, x.ID)
	})
	return ret
}

func (s Softdels) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *Softdel, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s Softdels) GetByID(id string) Softdels {
	return lo.Filter(s, func(xx *Softdel, _ int) bool {
		return xx.ID == id
	})
}

func (s Softdels) GetByIDs(ids ...string) Softdels {
	return lo.Filter(s, func(xx *Softdel, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (s Softdels) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), lo.Map(s, func(x *Softdel, _ int) []string {
		return x.Strings()
	})
}

func (s Softdels) Random() *Softdel {
	if len(s) == 0 {
		return nil
	}
	return s[util.RandomInt(len(s))]
}

func (s Softdels) Clone() Softdels {
	return slices.Clone(s)
}
