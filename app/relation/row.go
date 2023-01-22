// Content managed by Project Forge, see [projectforge.md] for details.
package relation

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "relation"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "basic_id", "name", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID      uuid.UUID `db:"id"`
	BasicID uuid.UUID `db:"basic_id"`
	Name    string    `db:"name"`
	Created time.Time `db:"created"`
}

func (r *row) ToRelation() *Relation {
	if r == nil {
		return nil
	}
	return &Relation{
		ID:      r.ID,
		BasicID: r.BasicID,
		Name:    r.Name,
		Created: r.Created,
	}
}

type rows []*row

func (x rows) ToRelations() Relations {
	ret := make(Relations, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToRelation())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
