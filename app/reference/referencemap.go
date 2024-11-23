package reference

import (
	"github.com/kyleu/pftest/app/foo"
	"github.com/kyleu/pftest/app/util"
)

func (r *Reference) ToMap() util.ValueMap {
	return util.ValueMap{"id": r.ID, "custom": r.Custom, "self": r.Self, "created": r.Created}
}

func ReferenceFromMap(m util.ValueMap, setPK bool) (*Reference, util.ValueMap, error) {
	ret := &Reference{}
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
		case "custom":
			tmpCustom, err := m.ParseString("custom", true, true)
			if err != nil {
				return nil, nil, err
			}
			customArg := &foo.Custom{}
			err = util.FromJSON([]byte(tmpCustom), customArg)
			if err != nil {
				return nil, nil, err
			}
			ret.Custom = customArg
		case "self":
			tmpSelf, err := m.ParseString("self", true, true)
			if err != nil {
				return nil, nil, err
			}
			selfArg := &SelfCustom{}
			err = util.FromJSON([]byte(tmpSelf), selfArg)
			if err != nil {
				return nil, nil, err
			}
			ret.Self = selfArg
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

func (r *Reference) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "id", V: r.ID}, {K: "custom", V: r.Custom}, {K: "self", V: r.Self}, {K: "created", V: r.Created}}
	return util.NewOrderedMap[any](false, 4, pairs...)
}
