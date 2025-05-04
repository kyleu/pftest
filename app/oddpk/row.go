package oddpk

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "oddpk"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"project", "path", "name"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = util.StringJoin(columnsQuoted, ", ")
)

type row struct {
	Project uuid.UUID `db:"project" json:"project"`
	Path    string    `db:"path" json:"path"`
	Name    string    `db:"name" json:"name"`
}

func (r *row) ToOddPK() *OddPK {
	if r == nil {
		return nil
	}
	return &OddPK{
		Project: r.Project,
		Path:    r.Path,
		Name:    r.Name,
	}
}

type rows []*row

func (x rows) ToOddPKs() OddPKs {
	return lo.Map(x, func(d *row, _ int) *OddPK {
		return d.ToOddPK()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"project\" = $%d and \"path\" = $%d", idx+1, idx+2)
}
