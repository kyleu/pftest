package mixed_case

import "github.com/kyleu/pftest/app/util"

type MixedCase struct {
	ID           string `json:"id"`
	TestField    string `json:"testField"`
	AnotherField string `json:"anotherField"`
}

func New(id string) *MixedCase {
	return &MixedCase{ID: id}
}

func FromMap(m util.ValueMap, setPK bool) (*MixedCase, error) {
	ret := &MixedCase{}
	var err error
	if setPK {
		ret.ID, err = m.ParseString("id", true, true)
		if err != nil {
			return nil, err
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	ret.TestField, err = m.ParseString("testField", true, true)
	if err != nil {
		return nil, err
	}
	ret.AnotherField, err = m.ParseString("anotherField", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (m *MixedCase) String() string {
	return m.ID
}

func (m *MixedCase) WebPath() string {
	return "/mixed_case" + "/" + m.ID
}

func (m *MixedCase) ToData() []interface{} {
	return []interface{}{m.ID, m.TestField, m.AnotherField}
}

type MixedCases []*MixedCase
