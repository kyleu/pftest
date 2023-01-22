// Content managed by Project Forge, see [projectforge.md] for details.
package capital

import (
	"fmt"
	"strings"
	"time"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "Capital"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"ID", "Name", "Birthday", "Version", "Deathday"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")

	columnsCore    = util.StringArrayQuoted([]string{"ID", "current_Version"})
	columnsVersion = util.StringArrayQuoted([]string{"Capital_ID", "Version", "Name", "Birthday", "Deathday"})

	tableVersion       = table + "_Version"
	tableVersionQuoted = fmt.Sprintf("%q", tableVersion)
	tablesJoined       = fmt.Sprintf(`%q c join %q cr on c."ID" = cr."Capital_ID" and c."current_Version" = cr."Version"`, table, tableVersion) //nolint
)

type row struct {
	ID       string     `db:"ID"`
	Name     string     `db:"Name"`
	Birthday time.Time  `db:"Birthday"`
	Version  int        `db:"Version"`
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
		Version:  r.Version,
		Deathday: r.Deathday,
	}
}

type rows []*row

func (x rows) ToCapitals() Capitals {
	ret := make(Capitals, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToCapital())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"ID\" = $%d", idx+1)
}
