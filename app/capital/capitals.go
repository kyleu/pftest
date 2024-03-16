// Package capital - Content managed by Project Forge, see [projectforge.md] for details.
package capital

import (
	"slices"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

type Capitals []*Capital

func (c Capitals) Get(id string) *Capital {
	return lo.FindOrElse(c, nil, func(x *Capital) bool {
		return x.ID == id
	})
}

func (c Capitals) IDs() []string {
	return lo.Map(c, func(xx *Capital, _ int) string {
		return xx.ID
	})
}

func (c Capitals) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(c)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(c, func(x *Capital, _ int) {
		ret = append(ret, x.ID)
	})
	return ret
}

func (c Capitals) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(c)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(c, func(x *Capital, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (c Capitals) GetByID(id string) Capitals {
	return lo.Filter(c, func(xx *Capital, _ int) bool {
		return xx.ID == id
	})
}

func (c Capitals) GetByIDs(ids ...string) Capitals {
	return lo.Filter(c, func(xx *Capital, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (c Capitals) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), lo.Map(c, func(x *Capital, _ int) []string {
		return x.Strings()
	})
}

func (c Capitals) Random() *Capital {
	if len(c) == 0 {
		return nil
	}
	return c[util.RandomInt(len(c))]
}

func (c Capitals) Clone() Capitals {
	return slices.Clone(c)
}
