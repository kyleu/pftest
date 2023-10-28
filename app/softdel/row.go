// Package softdel - Content managed by Project Forge, see [projectforge.md] for details.
package softdel

import (
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "softdel"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "created", "updated", "deleted"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID      string     `db:"id" json:"id"`
	Created time.Time  `db:"created" json:"created"`
	Updated *time.Time `db:"updated" json:"updated"`
	Deleted *time.Time `db:"deleted" json:"deleted"`
}

func (r *row) ToSoftdel() *Softdel {
	if r == nil {
		return nil
	}
	return &Softdel{
		ID:      r.ID,
		Created: r.Created,
		Updated: r.Updated,
		Deleted: r.Deleted,
	}
}

type rows []*row

func (x rows) ToSoftdels() Softdels {
	return lo.Map(x, func(d *row, _ int) *Softdel {
		return d.ToSoftdel()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
