package softdel

import "github.com/kyleu/pftest/app/util"

func (s *Softdel) ToMap() util.ValueMap {
	return util.ValueMap{"id": s.ID, "created": s.Created, "updated": s.Updated, "deleted": s.Deleted}
}

func SoftdelFromMap(m util.ValueMap, setPK bool) (*Softdel, util.ValueMap, error) {
	ret := &Softdel{}
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

func (s *Softdel) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "id", V: s.ID}, {K: "created", V: s.Created}, {K: "updated", V: s.Updated}, {K: "deleted", V: s.Deleted}}
	return util.NewOrderedMap(false, 4, pairs...)
}
