// Content managed by Project Forge, see [projectforge.md] for details.
package group

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "group"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "group", "data", "created", "updated", "deleted"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	ID      string          `db:"id"`
	Group   string          `db:"group"`
	Data    json.RawMessage `db:"data"`
	Created time.Time       `db:"created"`
	Updated *time.Time      `db:"updated"`
	Deleted *time.Time      `db:"deleted"`
}

func (d *dto) ToGroup() *Group {
	if d == nil {
		return nil
	}
	dataArg := util.ValueMap{}
	_ = util.FromJSON(d.Data, &dataArg)
	return &Group{
		ID:      d.ID,
		Group:   d.Group,
		Data:    dataArg,
		Created: d.Created,
		Updated: d.Updated,
		Deleted: d.Deleted,
	}
}

type dtos []*dto

func (x dtos) ToGroups() Groups {
	ret := make(Groups, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToGroup())
	}
	return ret
}


func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx + 1)
}
