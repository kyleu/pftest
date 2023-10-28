// Package path - Content managed by Project Forge, see [projectforge.md] for details.
package path

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "path"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "name", "status", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID      uuid.UUID `db:"id" json:"id"`
	Name    string    `db:"name" json:"name"`
	Status  string    `db:"status" json:"status"`
	Created time.Time `db:"created" json:"created"`
}

func (r *row) ToPath() *Path {
	if r == nil {
		return nil
	}
	return &Path{
		ID:      r.ID,
		Name:    r.Name,
		Status:  r.Status,
		Created: r.Created,
	}
}

type rows []*row

func (x rows) ToPaths() Paths {
	return lo.Map(x, func(d *row, _ int) *Path {
		return d.ToPath()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
