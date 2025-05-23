package relation

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

type Relations []*Relation

func (r Relations) Get(id uuid.UUID) *Relation {
	return lo.FindOrElse(r, nil, func(x *Relation) bool {
		return x.ID == id
	})
}

func (r Relations) IDs() []uuid.UUID {
	return lo.Map(r, func(xx *Relation, _ int) uuid.UUID {
		return xx.ID
	})
}

func (r Relations) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(r)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(r, func(x *Relation, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (r Relations) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(r)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(r, func(x *Relation, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (r Relations) GetByID(id uuid.UUID) Relations {
	return lo.Filter(r, func(xx *Relation, _ int) bool {
		return xx.ID == id
	})
}

func (r Relations) GetByIDs(ids ...uuid.UUID) Relations {
	return lo.Filter(r, func(xx *Relation, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (r Relations) BasicIDs() []uuid.UUID {
	return lo.Map(r, func(xx *Relation, _ int) uuid.UUID {
		return xx.BasicID
	})
}

func (r Relations) GetByBasicID(basicID uuid.UUID) Relations {
	return lo.Filter(r, func(xx *Relation, _ int) bool {
		return xx.BasicID == basicID
	})
}

func (r Relations) GetByBasicIDs(basicIDs ...uuid.UUID) Relations {
	return lo.Filter(r, func(xx *Relation, _ int) bool {
		return lo.Contains(basicIDs, xx.BasicID)
	})
}

func (r Relations) ToMaps() []util.ValueMap {
	return lo.Map(r, func(x *Relation, _ int) util.ValueMap {
		return x.ToMap()
	})
}

func (r Relations) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(r, func(x *Relation, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (r Relations) ToCSV() ([]string, [][]string) {
	return RelationFieldDescs.Keys(), lo.Map(r, func(x *Relation, _ int) []string {
		return x.Strings()
	})
}

func (r Relations) Random() *Relation {
	return util.RandomElement(r)
}

func (r Relations) Clone() Relations {
	return lo.Map(r, func(xx *Relation, _ int) *Relation {
		return xx.Clone()
	})
}
