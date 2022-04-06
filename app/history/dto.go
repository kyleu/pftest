// Content managed by Project Forge, see [projectforge.md] for details.
package history

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "history"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "data", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	ID      string          `db:"id"`
	Data    json.RawMessage `db:"data"`
	Created time.Time       `db:"created"`
	Updated *time.Time      `db:"updated"`
}

func (d *dto) ToHistory() *History {
	if d == nil {
		return nil
	}
	dataArg := util.ValueMap{}
	_ = util.FromJSON(d.Data, &dataArg)
	return &History{
		ID:      d.ID,
		Data:    dataArg,
		Created: d.Created,
		Updated: d.Updated,
	}
}

type dtos []*dto

func (x dtos) ToHistories() Histories {
	ret := make(Histories, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToHistory())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
