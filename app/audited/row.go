// Content managed by Project Forge, see [projectforge.md] for details.
package audited

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

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
	ret := make(Auditeds, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToAudited())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
