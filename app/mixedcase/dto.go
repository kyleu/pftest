// Content managed by Project Forge, see [projectforge.md] for details.
package mixedcase

import (
	"fmt"
	"strings"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "mixed_case"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "test_field", "another_field"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
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
	return &MixedCase{
		ID:           d.ID,
		TestField:    d.TestField,
		AnotherField: d.AnotherField,
	}
}

type dtos []*dto

func (x dtos) ToMixedCases() MixedCases {
	ret := make(MixedCases, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToMixedCase())
	}
	return ret
}


func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx + 1)
}
