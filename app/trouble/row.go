// Content managed by Project Forge, see [projectforge.md] for details.
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

	columnsCore      = util.StringArrayQuoted([]string{"from", "where", "current_selectcol", "limit", "delete"})
	columnsSelectcol = util.StringArrayQuoted([]string{"trouble_from", "trouble_where", "selectcol", "group"})

	tableSelectcol       = table + "_selectcol"
	tableSelectcolQuoted = fmt.Sprintf("%q", tableSelectcol)
	tablesJoined         = fmt.Sprintf(`%q t join %q tr on t."from" = tr."trouble_from" and t."where" = tr."trouble_where" and t."current_selectcol" = tr."selectcol"`, table, tableSelectcol) //nolint
)

type row struct {
	From      string          `db:"from"`
	Where     json.RawMessage `db:"where"`
	Selectcol int             `db:"selectcol"`
	Limit     string          `db:"limit"`
	Group     string          `db:"group"`
	Delete    *time.Time      `db:"delete"`
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
