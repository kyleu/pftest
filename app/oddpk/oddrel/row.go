package oddrel

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "oddrel"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "project", "path"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID      uuid.UUID `db:"id" json:"id"`
	Project uuid.UUID `db:"project" json:"project"`
	Path    string    `db:"path" json:"path"`
}

func (r *row) ToOddrel() *Oddrel {
	if r == nil {
		return nil
	}
	return &Oddrel{
		ID:      r.ID,
		Project: r.Project,
		Path:    r.Path,
	}
}

type rows []*row

func (x rows) ToOddrels() Oddrels {
	return lo.Map(x, func(d *row, _ int) *Oddrel {
		return d.ToOddrel()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
