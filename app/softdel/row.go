// Content managed by Project Forge, see [projectforge.md] for details.
package softdel

import (
	"fmt"
	"strings"
	"time"

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
	ID      string     `db:"id"`
	Created time.Time  `db:"created"`
	Updated *time.Time `db:"updated"`
	Deleted *time.Time `db:"deleted"`
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
	ret := make(Softdels, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToSoftdel())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}