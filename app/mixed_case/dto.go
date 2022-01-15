package mixed_case

import "strings"

var (
	table         = "mixed_case"
	columns       = []string{"id", "test_field", "another_field"}
	columnsString = strings.Join(columns, ", ")
)

type dto struct {
	ID           string `db:"id"`
	TestField    string `db:"test_field"`
	AnotherField string `db:"another_field"`
}

func (d *dto) ToMixedCase() *MixedCase {
	if d == nil {
		return nil
	}
	return &MixedCase{ID: d.ID, TestField: d.TestField, AnotherField: d.AnotherField}
}

type dtos []*dto

func (x dtos) ToMixedCases() MixedCases {
	ret := make(MixedCases, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToMixedCase())
	}
	return ret
}
