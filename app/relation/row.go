package relation

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "relation"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "basic_id", "name", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = util.StringJoin(columnsQuoted, ", ")
)

type row struct {
	ID      uuid.UUID `db:"id" json:"id"`
	BasicID uuid.UUID `db:"basic_id" json:"basic_id"`
	Name    string    `db:"name" json:"name"`
	Created time.Time `db:"created" json:"created"`
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
	return lo.Map(x, func(d *row, _ int) *Relation {
		return d.ToRelation()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
