package relation

import "github.com/kyleu/pftest/app/util"

func (r *Relation) ToMap() util.ValueMap {
	return util.ValueMap{"id": r.ID, "basicID": r.BasicID, "name": r.Name, "created": r.Created}
}

func RelationFromMap(m util.ValueMap, setPK bool) (*Relation, util.ValueMap, error) {
	ret := &Relation{}
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
		case "basicID":
			retBasicID, e := m.ParseUUID(k, true, true)
			if e != nil {
				return nil, nil, e
			}
			if retBasicID != nil {
				ret.BasicID = *retBasicID
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

func (r *Relation) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "id", V: r.ID}, {K: "basicID", V: r.BasicID}, {K: "name", V: r.Name}, {K: "created", V: r.Created}}
	return util.NewOrderedMap[any](false, 4, pairs...)
}
