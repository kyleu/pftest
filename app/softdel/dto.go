package softdel

import (
	"fmt"
	"strings"
	"time"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "softdel"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "created", "updated", "deleted"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	ID      string     `db:"id"`
	Created time.Time  `db:"created"`
	Updated *time.Time `db:"updated"`
	Deleted *time.Time `db:"deleted"`
}

func (d *dto) ToSoftdel() *Softdel {
	if d == nil {
		return nil
	}
	return &Softdel{ID: d.ID, Created: d.Created, Updated: d.Updated, Deleted: d.Deleted}
}

type dtos []*dto

func (x dtos) ToSoftdels() Softdels {
	ret := make(Softdels, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToSoftdel())
	}
	return ret
}
