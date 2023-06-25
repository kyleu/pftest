// Content managed by Project Forge, see [projectforge.md] for details.
package group

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "group"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "child", "data", "created", "updated", "deleted"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID      string          `db:"id"`
	Child   string          `db:"child"`
	Data    json.RawMessage `db:"data"`
	Created time.Time       `db:"created"`
	Updated *time.Time      `db:"updated"`
	Deleted *time.Time      `db:"deleted"`
}

func (r *row) ToGroup() *Group {
	if r == nil {
		return nil
	}
	dataArg := util.ValueMap{}
	_ = util.FromJSON(r.Data, &dataArg)
	return &Group{
		ID:      r.ID,
		Child:   r.Child,
		Data:    dataArg,
		Created: r.Created,
		Updated: r.Updated,
		Deleted: r.Deleted,
	}
}

type rows []*row

func (x rows) ToGroups() Groups {
	return lo.Map(x, func(d *row, _ int) *Group {
		return d.ToGroup()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
