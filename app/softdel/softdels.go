package softdel

import (
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
	return SoftdelFieldDescs.Keys(), lo.Map(s, func(x *Softdel, _ int) []string {
		return x.Strings()
	})
}

func (s Softdels) Random() *Softdel {
	return util.RandomElement(s)
}

func (s Softdels) Clone() Softdels {
	return lo.Map(s, func(xx *Softdel, _ int) *Softdel {
		return xx.Clone()
	})
}
