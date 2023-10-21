// Package capital - Content managed by Project Forge, see [projectforge.md] for details.
package capital

import (
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "Capital"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"ID", "Name", "Birthday", "Deathday"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID       string     `db:"ID"`
	Name     string     `db:"Name"`
	Birthday time.Time  `db:"Birthday"`
	Deathday *time.Time `db:"Deathday"`
}

func (r *row) ToCapital() *Capital {
	if r == nil {
		return nil
	}
	return &Capital{
		ID:       r.ID,
		Name:     r.Name,
		Birthday: r.Birthday,
		Deathday: r.Deathday,
	}
}

type rows []*row

func (x rows) ToCapitals() Capitals {
	return lo.Map(x, func(d *row, _ int) *Capital {
		return d.ToCapital()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"ID\" = $%d", idx+1)
}
