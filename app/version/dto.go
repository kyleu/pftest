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
	columns       = []string{"id", "revision", "const", "var", "created", "updated", "deleted"}
	columnsString = strings.Join(columns, ", ")

	columnsCore     = []string{"id", "current_revision", "const", "updated", "deleted"}
	columnsRevision = []string{"version_id", "revision", "var", "created"}

	tableRevision = table + "_revision"
	tablesJoined  = fmt.Sprintf("%q v join %q vr on v.id = vr.version_id and v.current_revision = vr.revision", table, tableRevision)
)

type dto struct {
	ID       string          `db:"id"`
	Revision int             `db:"revision"`
	Const    string          `db:"const"`
	Var      json.RawMessage `db:"var"`
	Created  time.Time       `db:"created"`
	Updated  *time.Time      `db:"updated"`
	Deleted  *time.Time      `db:"deleted"`
}

func (d *dto) ToVersion() *Version {
	if d == nil {
		return nil
	}
	varArg := util.ValueMap{}
	_ = util.FromJSON(d.Var, &varArg)
	return &Version{ID: d.ID, Revision: d.Revision, Const: d.Const, Var: varArg, Created: d.Created, Updated: d.Updated, Deleted: d.Deleted}
}

type dtos []*dto

func (x dtos) ToVersions() Versions {
	ret := make(Versions, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToVersion())
	}
	return ret
}
