package audited

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

type Auditeds []*Audited

func (a Auditeds) Get(id uuid.UUID) *Audited {
	return lo.FindOrElse(a, nil, func(x *Audited) bool {
		return x.ID == id
	})
}

func (a Auditeds) IDs() []uuid.UUID {
	return lo.Map(a, func(xx *Audited, _ int) uuid.UUID {
		return xx.ID
	})
}

func (a Auditeds) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(a)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(a, func(x *Audited, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (a Auditeds) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(a)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(a, func(x *Audited, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (a Auditeds) GetByID(id uuid.UUID) Auditeds {
	return lo.Filter(a, func(xx *Audited, _ int) bool {
		return xx.ID == id
	})
}

func (a Auditeds) GetByIDs(ids ...uuid.UUID) Auditeds {
	return lo.Filter(a, func(xx *Audited, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (a Auditeds) ToMaps() []util.ValueMap {
	return lo.Map(a, func(x *Audited, _ int) util.ValueMap {
		return x.ToMap()
	})
}

func (a Auditeds) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(a, func(x *Audited, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (a Auditeds) ToCSV() ([]string, [][]string) {
	return AuditedFieldDescs.Keys(), lo.Map(a, func(x *Audited, _ int) []string {
		return x.Strings()
	})
}

func (a Auditeds) Random() *Audited {
	return util.RandomElement(a)
}

func (a Auditeds) Clone() Auditeds {
	return lo.Map(a, func(xx *Audited, _ int) *Audited {
		return xx.Clone()
	})
}
