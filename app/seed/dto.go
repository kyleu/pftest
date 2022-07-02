// Content managed by Project Forge, see [projectforge.md] for details.
package seed

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "seed"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "name", "size", "obj"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	ID   uuid.UUID       `db:"id"`
	Name string          `db:"name"`
	Size int             `db:"size"`
	Obj  json.RawMessage `db:"obj"`
}

func (d *dto) ToSeed() *Seed {
	if d == nil {
		return nil
	}
	objArg := util.ValueMap{}
	_ = util.FromJSON(d.Obj, &objArg)
	return &Seed{
		ID:   d.ID,
		Name: d.Name,
		Size: d.Size,
		Obj:  objArg,
	}
}

type dtos []*dto

func (x dtos) ToSeeds() Seeds {
	ret := make(Seeds, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToSeed())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
