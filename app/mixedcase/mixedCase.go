// Content managed by Project Forge, see [projectforge.md] for details.
package mixedcase

import "github.com/kyleu/pftest/app/util"

type MixedCase struct {
	ID           string `json:"id"`
	TestField    string `json:"testField"`
	AnotherField string `json:"anotherField"`
}

func New(id string) *MixedCase {
	return &MixedCase{ID: id}
}

func Random() *MixedCase {
	return &MixedCase{
		ID:           util.RandomString(12),
		TestField:    util.RandomString(12),
		AnotherField: util.RandomString(12),
	}
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

func (m *MixedCase) Clone() *MixedCase {
	return &MixedCase{
		ID:           m.ID,
		TestField:    m.TestField,
		AnotherField: m.AnotherField,
	}
}

func (m *MixedCase) String() string {
	return m.ID
}

func (m *MixedCase) WebPath() string {
	return "/mixedcase" + "/" + m.ID
}

func (m *MixedCase) ToData() []interface{} {
	return []interface{}{m.ID, m.TestField, m.AnotherField}
}

type MixedCases []*MixedCase
