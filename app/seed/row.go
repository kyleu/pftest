// Package seed - Content managed by Project Forge, see [projectforge.md] for details.
package seed

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "seed"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "name", "size", "obj"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID   uuid.UUID       `db:"id"`
	Name string          `db:"name"`
	Size int             `db:"size"`
	Obj  json.RawMessage `db:"obj"`
}

func (r *row) ToSeed() *Seed {
	if r == nil {
		return nil
	}
	objArg := util.ValueMap{}
	_ = util.FromJSON(r.Obj, &objArg)
	return &Seed{
		ID:   r.ID,
		Name: r.Name,
		Size: r.Size,
		Obj:  objArg,
	}
}

type rows []*row

func (x rows) ToSeeds() Seeds {
	return lo.Map(x, func(d *row, _ int) *Seed {
		return d.ToSeed()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
