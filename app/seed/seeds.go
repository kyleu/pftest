package seed

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

type Seeds []*Seed

func (s Seeds) Get(id uuid.UUID) *Seed {
	return lo.FindOrElse(s, nil, func(x *Seed) bool {
		return x.ID == id
	})
}

func (s Seeds) IDs() []uuid.UUID {
	return lo.Map(s, func(xx *Seed, _ int) uuid.UUID {
		return xx.ID
	})
}

func (s Seeds) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(s)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(s, func(x *Seed, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (s Seeds) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(s)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(s, func(x *Seed, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (s Seeds) GetByID(id uuid.UUID) Seeds {
	return lo.Filter(s, func(xx *Seed, _ int) bool {
		return xx.ID == id
	})
}

func (s Seeds) GetByIDs(ids ...uuid.UUID) Seeds {
	return lo.Filter(s, func(xx *Seed, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (s Seeds) ToMaps() []util.ValueMap {
	return lo.Map(s, func(xx *Seed, _ int) util.ValueMap {
		return xx.ToMap()
	})
}

func (s Seeds) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(s, func(x *Seed, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (s Seeds) ToCSV() ([]string, [][]string) {
	return SeedFieldDescs.Keys(), lo.Map(s, func(x *Seed, _ int) []string {
		return x.Strings()
	})
}

func (s Seeds) Random() *Seed {
	return util.RandomElement(s)
}

func (s Seeds) Clone() Seeds {
	return lo.Map(s, func(xx *Seed, _ int) *Seed {
		return xx.Clone()
	})
}
