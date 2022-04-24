// Content managed by Project Forge, see [projectforge.md] for details.
package audit

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

var (
	recordTable         = "audit_record"
	recordTableQuoted   = fmt.Sprintf("%q", recordTable)
	recordColumns       = []string{"id", "audit_id", "t", "pk", "changes", "metadata", "occurred"}
	recordColumnsQuoted = util.StringArrayQuoted(recordColumns)
	recordColumnsString = strings.Join(recordColumnsQuoted, ", ")
)

type recordDTO struct {
	ID       uuid.UUID       `db:"id"`
	AuditID  uuid.UUID       `db:"audit_id"`
	T        string          `db:"t"`
	PK       string          `db:"pk"`
	Changes  json.RawMessage `db:"changes"`
	Metadata json.RawMessage `db:"metadata"`
	Occurred time.Time       `db:"occurred"`
}

func (d *recordDTO) ToRecord() *Record {
	if d == nil {
		return nil
	}
	changesArg := util.Diffs{}
	_ = util.FromJSON(d.Changes, &changesArg)
	metadataArg := util.ValueMap{}
	_ = util.FromJSON(d.Metadata, &metadataArg)
	return &Record{ID: d.ID, AuditID: d.AuditID, T: d.T, PK: d.PK, Changes: changesArg, Metadata: metadataArg, Occurred: d.Occurred}
}

type recordDTOs []*recordDTO

func (x recordDTOs) ToRecords() Records {
	ret := make(Records, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToRecord())
	}
	return ret
}
