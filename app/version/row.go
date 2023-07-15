// Content managed by Project Forge, see [projectforge.md] for details.
package version

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "version"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "revision", "constcol", "varcol", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")

	columnsCore     = util.StringArrayQuoted([]string{"id", "current_revision", "constcol", "updated"})
	columnsRevision = util.StringArrayQuoted([]string{"version_id", "revision", "varcol", "created"})

	tableRevision       = table + "_revision"
	tableRevisionQuoted = fmt.Sprintf("%q", tableRevision)
	tablesJoined        = fmt.Sprintf(`%q v join %q vr on v."id" = vr."version_id" and v."current_revision" = vr."revision"`, table, tableRevision)
)

type row struct {
	ID       string          `db:"id"`
	Revision int             `db:"revision"`
	Constcol string          `db:"constcol"`
	Varcol   json.RawMessage `db:"varcol"`
	Created  time.Time       `db:"created"`
	Updated  *time.Time      `db:"updated"`
}

func (r *row) ToVersion() *Version {
	if r == nil {
		return nil
	}
	varcolArg := util.ValueMap{}
	_ = util.FromJSON(r.Varcol, &varcolArg)
	return &Version{
		ID:       r.ID,
		Revision: r.Revision,
		Constcol: r.Constcol,
		Varcol:   varcolArg,
		Created:  r.Created,
		Updated:  r.Updated,
	}
}

type rows []*row

func (x rows) ToVersions() Versions {
	return lo.Map(x, func(d *row, _ int) *Version {
		return d.ToVersion()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
