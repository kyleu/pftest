// Package hist - Content managed by Project Forge, see [projectforge.md] for details.
package hist

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "hist"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "data", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID      string          `db:"id"`
	Data    json.RawMessage `db:"data"`
	Created time.Time       `db:"created"`
	Updated *time.Time      `db:"updated"`
}

func (r *row) ToHist() *Hist {
	if r == nil {
		return nil
	}
	dataArg := util.ValueMap{}
	_ = util.FromJSON(r.Data, &dataArg)
	return &Hist{
		ID:      r.ID,
		Data:    dataArg,
		Created: r.Created,
		Updated: r.Updated,
	}
}

type rows []*row

func (x rows) ToHists() Hists {
	return lo.Map(x, func(d *row, _ int) *Hist {
		return d.ToHist()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
