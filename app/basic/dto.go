// Content managed by Project Forge, see [projectforge.md] for details.
package basic

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "basic"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "name", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
	defaultWC     = "\"id\" = $1"
)

type dto struct {
	ID      uuid.UUID `db:"id"`
	Name    string    `db:"name"`
	Created time.Time `db:"created"`
}

func (d *dto) ToBasic() *Basic {
	if d == nil {
		return nil
	}
	return &Basic{ID: d.ID, Name: d.Name, Created: d.Created}
}

type dtos []*dto

func (x dtos) ToBasics() Basics {
	ret := make(Basics, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToBasic())
	}
	return ret
}
