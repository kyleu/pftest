package mixedcase

import (
	"net/url"
	"path"

	"github.com/kyleu/pftest/app/lib/svc"
	"github.com/kyleu/pftest/app/util"
)

const DefaultRoute = "/mixedcase"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(paths...)
}

var _ svc.Model = (*MixedCase)(nil)

type MixedCase struct {
	ID           string `json:"id,omitempty"`
	TestField    string `json:"testField,omitempty"`
	AnotherField string `json:"anotherField,omitempty"`
}

func NewMixedCase(id string) *MixedCase {
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

func RandomMixedCase() *MixedCase {
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
	return MixedCaseFieldDescs.Keys(), [][]string{m.Strings()}
}

func (m *MixedCase) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(m.ID))...)
}

func (m *MixedCase) Breadcrumb(extra ...string) string {
	return m.TitleString() + "||" + m.WebPath(extra...) + "**star"
}

func (m *MixedCase) ToData() []any {
	return []any{m.ID, m.TestField, m.AnotherField}
}

var MixedCaseFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "string"},
	{Key: "testField", Title: "Test Field", Description: "", Type: "string"},
	{Key: "anotherField", Title: "Another Field", Description: "", Type: "string"},
}
