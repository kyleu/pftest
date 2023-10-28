// Package trouble - Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "trouble"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"from", "where", "selectcol", "limit", "group", "delete"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	From      string          `db:"from" json:"from"`
	Where     json.RawMessage `db:"where" json:"where"`
	Selectcol int             `db:"selectcol" json:"selectcol"`
	Limit     string          `db:"limit" json:"limit"`
	Group     string          `db:"group" json:"group"`
	Delete    *time.Time      `db:"delete" json:"delete"`
}

func (r *row) ToTrouble() *Trouble {
	if r == nil {
		return nil
	}
	whereArg := []string{}
	_ = util.FromJSON(r.Where, &whereArg)
	return &Trouble{
		From:      r.From,
		Where:     whereArg,
		Selectcol: r.Selectcol,
		Limit:     r.Limit,
		Group:     r.Group,
		Delete:    r.Delete,
	}
}

type rows []*row

func (x rows) ToTroubles() Troubles {
	return lo.Map(x, func(d *row, _ int) *Trouble {
		return d.ToTrouble()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"from\" = $%d and \"where\" = $%d", idx+1, idx+2)
}
