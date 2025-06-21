package mixedcase

import "github.com/kyleu/pftest/app/util"

func (m *MixedCase) ToMap() util.ValueMap {
	return util.ValueMap{"id": m.ID, "testField": m.TestField, "anotherField": m.AnotherField}
}

func MixedCaseFromMap(m util.ValueMap, setPK bool) (*MixedCase, util.ValueMap, error) {
	ret := &MixedCase{}
	extra := util.ValueMap{}
	for k, v := range m {
		var err error
		switch k {
		case "id":
			if setPK {
				ret.ID, err = m.ParseString(k, true, true)
			}
		case "testField":
			ret.TestField, err = m.ParseString(k, true, true)
		case "anotherField":
			ret.AnotherField, err = m.ParseString(k, true, true)
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

func (m *MixedCase) ToOrderedMap() *util.OrderedMap[any] {
	pairs := util.OrderedPairs[any]{{K: "id", V: m.ID}, {K: "testField", V: m.TestField}, {K: "anotherField", V: m.AnotherField}}
	return util.NewOrderedMap(false, 4, pairs...)
}
