// Content managed by Project Forge, see [projectforge.md] for details.
package path

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

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
	ID      uuid.UUID `db:"id"`
	Name    string    `db:"name"`
	Status  string    `db:"status"`
	Created time.Time `db:"created"`
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
	ret := make(Paths, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToPath())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
