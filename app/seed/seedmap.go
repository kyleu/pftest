package seed

import "github.com/kyleu/pftest/app/util"

func (s *Seed) ToMap() util.ValueMap {
	return util.ValueMap{"id": s.ID, "name": s.Name, "size": s.Size, "obj": s.Obj}
}

func FromMap(m util.ValueMap, setPK bool) (*Seed, util.ValueMap, error) {
	ret := &Seed{}
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
		case "size":
			ret.Size, err = m.ParseInt(k, true, true)
		case "obj":
			ret.Obj, err = m.ParseMap(k, true, true)
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
