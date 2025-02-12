package timestamp

import (
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

type Timestamps []*Timestamp

func (t Timestamps) Get(id string) *Timestamp {
	return lo.FindOrElse(t, nil, func(x *Timestamp) bool {
		return x.ID == id
	})
}

func (t Timestamps) IDs() []string {
	return lo.Map(t, func(xx *Timestamp, _ int) string {
		return xx.ID
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

func (t Timestamps) GetByID(id string) Timestamps {
	return lo.Filter(t, func(xx *Timestamp, _ int) bool {
		return xx.ID == id
	})
}

func (t Timestamps) GetByIDs(ids ...string) Timestamps {
	return lo.Filter(t, func(xx *Timestamp, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (t Timestamps) ToMaps() []util.ValueMap {
	return lo.Map(t, func(x *Timestamp, _ int) util.ValueMap {
		return x.ToMap()
	})
}

func (t Timestamps) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(t, func(x *Timestamp, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (t Timestamps) ToCSV() ([]string, [][]string) {
	return TimestampFieldDescs.Keys(), lo.Map(t, func(x *Timestamp, _ int) []string {
		return x.Strings()
	})
}

func (t Timestamps) Random() *Timestamp {
	return util.RandomElement(t)
}

func (t Timestamps) Clone() Timestamps {
	return lo.Map(t, func(xx *Timestamp, _ int) *Timestamp {
		return xx.Clone()
	})
}
