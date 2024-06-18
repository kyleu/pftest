package mixedcase

import (
	"net/url"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

var _ svc.Model = (*MixedCase)(nil)

type MixedCase struct {
	ID           string `json:"id,omitempty"`
	TestField    string `json:"testField,omitempty"`
	AnotherField string `json:"anotherField,omitempty"`
}

func New(id string) *MixedCase {
	return &MixedCase{ID: id}
}

func (m *MixedCase) Clone() *MixedCase {
	return &MixedCase{m.ID, m.TestField, m.AnotherField}
}

func (m *MixedCase) String() string {
	return m.ID
}

func (m *MixedCase) TitleString() string {
	return m.String()
}

func Random() *MixedCase {
	return &MixedCase{
		ID:           util.RandomString(12),
		TestField:    util.RandomString(12),
		AnotherField: util.RandomString(12),
	}
}

func (m *MixedCase) Strings() []string {
	return []string{m.ID, m.TestField, m.AnotherField}
}

func (m *MixedCase) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{m.Strings()}
}

func (m *MixedCase) WebPath() string {
	return "/mixedcase/" + url.QueryEscape(m.ID)
}

func (m *MixedCase) ToData() []any {
	return []any{m.ID, m.TestField, m.AnotherField}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "string"},
	{Key: "testField", Title: "Test Field", Description: "", Type: "string"},
	{Key: "anotherField", Title: "Another Field", Description: "", Type: "string"},
}
