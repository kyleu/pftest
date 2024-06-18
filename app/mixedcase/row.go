package mixedcase

import (
	"fmt"
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "mixed_case"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "test_field", "another_field"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID           string `db:"id" json:"id"`
	TestField    string `db:"test_field" json:"test_field"`
	AnotherField string `db:"another_field" json:"another_field"`
}

func (r *row) ToMixedCase() *MixedCase {
	if r == nil {
		return nil
	}
	return &MixedCase{
		ID:           r.ID,
		TestField:    r.TestField,
		AnotherField: r.AnotherField,
	}
}

type rows []*row

func (x rows) ToMixedCases() MixedCases {
	return lo.Map(x, func(d *row, _ int) *MixedCase {
		return d.ToMixedCase()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
