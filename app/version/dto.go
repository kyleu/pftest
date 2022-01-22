// Content managed by Project Forge, see [projectforge.md] for details.
package version

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "version"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "revision", "constcol", "varcol", "created", "updated"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
	defaultWC     = "\"id\" = $1"

	columnsCore     = util.StringArrayQuoted([]string{"id", "current_revision", "constcol", "updated"})
	columnsRevision = util.StringArrayQuoted([]string{"version_id", "revision", "varcol", "created"})

	tableRevision       = table + "_revision"
	tableRevisionQuoted = fmt.Sprintf("%q", tableRevision)
	tablesJoined        = fmt.Sprintf(`%q v join %q vr on v."id" = vr."version_id" and v."current_revision" = vr."revision"`, table, tableRevision) // nolint
)

type dto struct {
	ID       string          `db:"id"`
	Revision int             `db:"revision"`
	Constcol string          `db:"constcol"`
	Varcol   json.RawMessage `db:"varcol"`
	Created  time.Time       `db:"created"`
	Updated  *time.Time      `db:"updated"`
}

func (d *dto) ToVersion() *Version {
	if d == nil {
		return nil
	}
	varcolArg := util.ValueMap{}
	_ = util.FromJSON(d.Varcol, &varcolArg)
	return &Version{ID: d.ID, Revision: d.Revision, Constcol: d.Constcol, Varcol: varcolArg, Created: d.Created, Updated: d.Updated}
}

type dtos []*dto

func (x dtos) ToVersions() Versions {
	ret := make(Versions, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToVersion())
	}
	return ret
}
