package basic

import "github.com/kyleu/pftest/app/util"

func (b *Basic) ToMap() util.ValueMap {
	return util.ValueMap{"id": b.ID, "name": b.Name, "status": b.Status, "created": b.Created}
}

func BasicFromMap(m util.ValueMap, setPK bool) (*Basic, util.ValueMap, error) {
	ret := &Basic{}
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
