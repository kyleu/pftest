package oddpk

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

type OddPKs []*OddPK

func (o OddPKs) Get(project uuid.UUID, path string) *OddPK {
	return lo.FindOrElse(o, nil, func(x *OddPK) bool {
		return x.Project == project && x.Path == path
	})
}

func (o OddPKs) Projects() []uuid.UUID {
	return lo.Map(o, func(xx *OddPK, _ int) uuid.UUID {
		return xx.Project
	})
}

func (o OddPKs) ProjectStrings(includeNil bool) []string {
	ret := make([]string, 0, len(o)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(o, func(x *OddPK, _ int) {
		ret = append(ret, x.Project.String())
	})
	return ret
}

func (o OddPKs) Paths() []string {
	return lo.Map(o, func(xx *OddPK, _ int) string {
		return xx.Path
	})
}

func (o OddPKs) PathStrings(includeNil bool) []string {
	ret := make([]string, 0, len(o)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(o, func(x *OddPK, _ int) {
		ret = append(ret, x.Path)
	})
	return ret
}

func (o OddPKs) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(o)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(o, func(x *OddPK, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (o OddPKs) ToPKs() []*PK {
	return lo.Map(o, func(x *OddPK, _ int) *PK {
		return x.ToPK()
	})
}

func (o OddPKs) GetByProject(project uuid.UUID) OddPKs {
	return lo.Filter(o, func(xx *OddPK, _ int) bool {
		return xx.Project == project
	})
}

func (o OddPKs) GetByProjects(projects ...uuid.UUID) OddPKs {
	return lo.Filter(o, func(xx *OddPK, _ int) bool {
		return lo.Contains(projects, xx.Project)
	})
}

func (o OddPKs) GetByPath(path string) OddPKs {
	return lo.Filter(o, func(xx *OddPK, _ int) bool {
		return xx.Path == path
	})
}

func (o OddPKs) GetByPaths(paths ...string) OddPKs {
	return lo.Filter(o, func(xx *OddPK, _ int) bool {
		return lo.Contains(paths, xx.Path)
	})
}

func (o OddPKs) ToMaps() []util.ValueMap {
	return lo.Map(o, func(x *OddPK, _ int) util.ValueMap {
		return x.ToMap()
	})
}

func (o OddPKs) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(o, func(x *OddPK, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (o OddPKs) ToCSV() ([]string, [][]string) {
	return OddPKFieldDescs.Keys(), lo.Map(o, func(x *OddPK, _ int) []string {
		return x.Strings()
	})
}

func (o OddPKs) Random() *OddPK {
	return util.RandomElement(o)
}

func (o OddPKs) Clone() OddPKs {
	return lo.Map(o, func(xx *OddPK, _ int) *OddPK {
		return xx.Clone()
	})
}
