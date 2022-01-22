// Content managed by Project Forge, see [projectforge.md] for details.
package timestamp

import (
	"fmt"
	"strings"
	"time"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "timestamp"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "created", "updated", "deleted"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
	defaultWC     = "\"id\" = $1"
)

type dto struct {
	ID      string     `db:"id"`
	Created time.Time  `db:"created"`
	Updated *time.Time `db:"updated"`
	Deleted *time.Time `db:"deleted"`
}

func (d *dto) ToTimestamp() *Timestamp {
	if d == nil {
		return nil
	}
	return &Timestamp{ID: d.ID, Created: d.Created, Updated: d.Updated, Deleted: d.Deleted}
}

type dtos []*dto

func (x dtos) ToTimestamps() Timestamps {
	ret := make(Timestamps, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToTimestamp())
	}
	return ret
}
