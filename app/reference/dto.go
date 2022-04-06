// Content managed by Project Forge, see [projectforge.md] for details.
package reference

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/foo"
	"github.com/kyleu/pftest/app/util"
)

var (
	table         = "reference"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "custom", "self", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type dto struct {
	ID      uuid.UUID       `db:"id"`
	Custom  json.RawMessage `db:"custom"`
	Self    json.RawMessage `db:"self"`
	Created time.Time       `db:"created"`
}

func (d *dto) ToReference() *Reference {
	if d == nil {
		return nil
	}
	customArg := &foo.Custom{}
	_ = util.FromJSON(d.Custom, customArg)
	selfArg := &SelfCustom{}
	_ = util.FromJSON(d.Self, selfArg)
	return &Reference{
		ID:      d.ID,
		Custom:  customArg,
		Self:    selfArg,
		Created: d.Created,
	}
}

type dtos []*dto

func (x dtos) ToReferences() References {
	ret := make(References, 0, len(x))
	for _, d := range x {
		ret = append(ret, d.ToReference())
	}
	return ret
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
