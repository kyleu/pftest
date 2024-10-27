package timestamp

import "github.com/kyleu/pftest/app/util"

func (t *Timestamp) ToMap() util.ValueMap {
	return util.ValueMap{"id": t.ID, "created": t.Created, "updated": t.Updated, "deleted": t.Deleted}
}

func TimestampFromMap(m util.ValueMap, setPK bool) (*Timestamp, util.ValueMap, error) {
	ret := &Timestamp{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "id":
			if setPK {
				ret.ID, err = m.ParseString(k, true, true)
			}
		case "deleted":
			ret.Deleted, err = m.ParseTime(k, true, true)
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
