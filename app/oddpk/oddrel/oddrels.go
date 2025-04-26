package oddrel

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

type Oddrels []*Oddrel

func (o Oddrels) Get(id uuid.UUID) *Oddrel {
	return lo.FindOrElse(o, nil, func(x *Oddrel) bool {
		return x.ID == id
	})
}

func (o Oddrels) IDs() []uuid.UUID {
	return lo.Map(o, func(xx *Oddrel, _ int) uuid.UUID {
		return xx.ID
	})
}

func (o Oddrels) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(o)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(o, func(x *Oddrel, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (o Oddrels) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(o)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(o, func(x *Oddrel, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (o Oddrels) GetByID(id uuid.UUID) Oddrels {
	return lo.Filter(o, func(xx *Oddrel, _ int) bool {
		return xx.ID == id
	})
}

func (o Oddrels) GetByIDs(ids ...uuid.UUID) Oddrels {
	return lo.Filter(o, func(xx *Oddrel, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (o Oddrels) ToMaps() []util.ValueMap {
	return lo.Map(o, func(x *Oddrel, _ int) util.ValueMap {
		return x.ToMap()
	})
}

func (o Oddrels) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(o, func(x *Oddrel, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (o Oddrels) ToCSV() ([]string, [][]string) {
	return OddrelFieldDescs.Keys(), lo.Map(o, func(x *Oddrel, _ int) []string {
		return x.Strings()
	})
}

func (o Oddrels) Random() *Oddrel {
	return util.RandomElement(o)
}

func (o Oddrels) Clone() Oddrels {
	return lo.Map(o, func(xx *Oddrel, _ int) *Oddrel {
		return xx.Clone()
	})
}
