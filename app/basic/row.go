// Package basic - Content managed by Project Forge, see [projectforge.md] for details.
package basic

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "basic"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "name", "status", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID      uuid.UUID `db:"id"`
	Name    string    `db:"name"`
	Status  string    `db:"status"`
	Created time.Time `db:"created"`
}

func (r *row) ToBasic() *Basic {
	if r == nil {
		return nil
	}
	return &Basic{
		ID:      r.ID,
		Name:    r.Name,
		Status:  r.Status,
		Created: r.Created,
	}
}

type rows []*row

func (x rows) ToBasics() Basics {
	return lo.Map(x, func(d *row, _ int) *Basic {
		return d.ToBasic()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
