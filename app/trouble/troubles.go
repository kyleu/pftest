// Package trouble - Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import (
	"slices"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

type Troubles []*Trouble

func (t Troubles) Get(from string, where []string) *Trouble {
	return lo.FindOrElse(t, nil, func(x *Trouble) bool {
		return x.From == from && slices.Equal(x.Where, where)
	})
}

func (t Troubles) Froms() []string {
	return lo.Map(t, func(xx *Trouble, _ int) string {
		return xx.From
	})
}

func (t Troubles) FromStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(t, func(x *Trouble, _ int) {
		ret = append(ret, x.From)
	})
	return ret
}

func (t Troubles) Wheres() [][]string {
	return lo.Map(t, func(xx *Trouble, _ int) []string {
		return xx.Where
	})
}

func (t Troubles) WhereStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(t, func(x *Trouble, _ int) {
		ret = append(ret, util.ToJSON(x.Where))
	})
	return ret
}

func (t Troubles) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(t, func(x *Trouble, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (t Troubles) ToPKs() []*PK {
	return lo.Map(t, func(x *Trouble, _ int) *PK {
		return x.ToPK()
	})
}

func (t Troubles) GetByFrom(from string) Troubles {
	return lo.Filter(t, func(xx *Trouble, _ int) bool {
		return xx.From == from
	})
}

func (t Troubles) GetByFroms(froms ...string) Troubles {
	return lo.Filter(t, func(xx *Trouble, _ int) bool {
		return lo.Contains(froms, xx.From)
	})
}

func (t Troubles) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), lo.Map(t, func(x *Trouble, _ int) []string {
		return x.Strings()
	})
}

func (t Troubles) Random() *Trouble {
	return util.RandomElement(t)
}

func (t Troubles) Clone() Troubles {
	return slices.Clone(t)
}
