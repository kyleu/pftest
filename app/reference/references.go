package reference

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

type References []*Reference

func (r References) Get(id uuid.UUID) *Reference {
	return lo.FindOrElse(r, nil, func(x *Reference) bool {
		return x.ID == id
	})
}

func (r References) IDs() []uuid.UUID {
	return lo.Map(r, func(xx *Reference, _ int) uuid.UUID {
		return xx.ID
	})
}

func (r References) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *Reference, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (r References) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(r, func(x *Reference, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (r References) GetByID(id uuid.UUID) References {
	return lo.Filter(r, func(xx *Reference, _ int) bool {
		return xx.ID == id
	})
}

func (r References) GetByIDs(ids ...uuid.UUID) References {
	return lo.Filter(r, func(xx *Reference, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (r References) ToMaps() []util.ValueMap {
	return lo.Map(r, func(x *Reference, _ int) util.ValueMap {
		return x.ToMap()
	})
}

func (r References) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(r, func(x *Reference, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (r References) ToCSV() ([]string, [][]string) {
	return ReferenceFieldDescs.Keys(), lo.Map(r, func(x *Reference, _ int) []string {
		return x.Strings()
	})
}

func (r References) Random() *Reference {
	return util.RandomElement(r)
}

func (r References) Clone() References {
	return lo.Map(r, func(xx *Reference, _ int) *Reference {
		return xx.Clone()
	})
}
