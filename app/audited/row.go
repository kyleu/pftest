// Package audited - Content managed by Project Forge, see [projectforge.md] for details.
package audited

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "audited"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "name"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

func (r *row) ToAudited() *Audited {
	if r == nil {
		return nil
	}
	return &Audited{
		ID:   r.ID,
		Name: r.Name,
	}
}

type rows []*row

func (x rows) ToAuditeds() Auditeds {
	return lo.Map(x, func(d *row, _ int) *Audited {
		return d.ToAudited()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
