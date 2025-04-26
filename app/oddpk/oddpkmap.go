package oddpk

import "github.com/kyleu/pftest/app/util"

func (o *OddPK) ToMap() util.ValueMap {
	return util.ValueMap{"project": o.Project, "path": o.Path, "name": o.Name}
}

func OddPKFromMap(m util.ValueMap, setPK bool) (*OddPK, util.ValueMap, error) {
	ret := &OddPK{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "project":
			if setPK {
				retProject, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retProject != nil {
					ret.Project = *retProject
				}
			}
		case "path":
			if setPK {
				ret.Path, err = m.ParseString(k, true, true)
			}
		case "name":
			ret.Name, err = m.ParseString(k, true, true)
		default:
			extra[k] = v
		}
		if err != nil {
			return nil, nil, err
		}
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, extra, nil
}

func (o *OddPK) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "project", V: o.Project}, {K: "path", V: o.Path}, {K: "name", V: o.Name}}
	return util.NewOrderedMap[any](false, 4, pairs...)
}
