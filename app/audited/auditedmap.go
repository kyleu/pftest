package audited

import "github.com/kyleu/pftest/app/util"

func (a *Audited) ToMap() util.ValueMap {
	return util.ValueMap{"id": a.ID, "name": a.Name}
}

func AuditedFromMap(m util.ValueMap, setPK bool) (*Audited, util.ValueMap, error) {
	ret := &Audited{}
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

func (a *Audited) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "id", V: a.ID}, {K: "name", V: a.Name}}
	return util.NewOrderedMap(false, 4, pairs...)
}
