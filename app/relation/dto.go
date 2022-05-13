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

type dto struct {
	ID      uuid.UUID `db:"id"`
	BasicID uuid.UUID `db:"basic_id"`
	Name    string    `db:"name"`
	Created time.Time `db:"created"`
}

func (d *dto) ToRelation() *Relation {
	if d == nil {
		return nil
	}
	return &Relation{
		ID:      d.ID,
		BasicID: d.BasicID,
		Name:    d.Name,
		Created: d.Created,
	}
}

type dtos []*dto

func (x dtos) ToRelations() Relations {
	ret := make(Relations, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToRelation())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
