package timestamp

import (
	"strings"
	"time"
)

var (
	table         = "timestamp"
	columns       = []string{"id", "created", "updated", "deleted"}
	columnsString = strings.Join(columns, ", ")
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
