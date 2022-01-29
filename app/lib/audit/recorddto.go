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
	recordTable         = "audit"
	recordTableQuoted   = fmt.Sprintf("%q", table)
	recordColumns       = []string{"id", "app", "act", "client", "server", "user", "metadata", "message", "started", "completed"}
	recordColumnsQuoted = util.StringArrayQuoted(columns)
	recordColumnsString = strings.Join(columnsQuoted, ", ")
	recordDefaultWC     = "\"id\" = $1"
)

type recordDTO struct {
	ID        uuid.UUID       `db:"id"`
	AuditID   uuid.UUID       `db:"audit_id"`
	App       string          `db:"app"`
	Act       string          `db:"act"`
	Client    string          `db:"client"`
	Server    string          `db:"server"`
	User      string          `db:"user"`
	Metadata  json.RawMessage `db:"metadata"`
	Message   string          `db:"message"`
	Started   time.Time       `db:"started"`
	Completed time.Time       `db:"completed"`
}

func (d *recordDTO) ToRecord() *Record {
	if d == nil {
		return nil
	}
	metadataArg := util.ValueMap{}
	_ = util.FromJSON(d.Metadata, &metadataArg)
	return &Record{ID: d.ID, AuditID: d.AuditID, App: d.App, Act: d.Act, Client: d.Client, Server: d.Server, User: d.User, Metadata: metadataArg, Message: d.Message, Started: d.Started, Completed: d.Completed}
}

type recordDTOs []*recordDTO

func (x recordDTOs) ToRecords() Records {
	ret := make(Records, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToRecord())
	}
	return ret
}
