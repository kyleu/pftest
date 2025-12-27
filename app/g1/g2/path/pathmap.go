package path

import "github.com/kyleu/pftest/app/util"

func (p *Path) ToMap() util.ValueMap {
	return util.ValueMap{"id": p.ID, "name": p.Name, "status": p.Status, "created": p.Created}
}

func PathFromMap(m util.ValueMap, setPK bool) (*Path, util.ValueMap, error) {
	ret := &Path{}
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
		case "name":
			ret.Name, err = m.ParseString(k, true, true)
		case "status":
			ret.Status, err = m.ParseString(k, true, true)
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

func (p *Path) ToOrderedMap() *util.OrderedMap[any] {
	if p == nil {
		return nil
	}
	pairs := util.OrderedPairs[any]{{K: "id", V: p.ID}, {K: "name", V: p.Name}, {K: "status", V: p.Status}, {K: "created", V: p.Created}}
	return util.NewOrderedMap(false, 4, pairs...)
}
