package oddrel

import "github.com/kyleu/pftest/app/util"

func (o *Oddrel) ToMap() util.ValueMap {
	return util.ValueMap{"id": o.ID, "project": o.Project, "path": o.Path}
}

func OddrelFromMap(m util.ValueMap, setPK bool) (*Oddrel, util.ValueMap, error) {
	ret := &Oddrel{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "id":
			if setPK {
				retID, e := m.ParseUUID(k, true, true)
				if e != nil {
					return nil, nil, e
				}
				if retID != nil {
					ret.ID = *retID
				}
			}
		case "project":
			retProject, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retProject != nil {
				ret.Project = *retProject
			}
		case "path":
			ret.Path, err = m.ParseString(k, true, true)
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

func (o *Oddrel) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "id", V: o.ID}, {K: "project", V: o.Project}, {K: "path", V: o.Path}}
	return util.NewOrderedMap[any](false, 4, pairs...)
}
