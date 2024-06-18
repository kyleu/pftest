package reference

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

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

type row struct {
	ID      uuid.UUID       `db:"id" json:"id"`
	Custom  json.RawMessage `db:"custom" json:"custom"`
	Self    json.RawMessage `db:"self" json:"self"`
	Created time.Time       `db:"created" json:"created"`
}

func (r *row) ToReference() *Reference {
	if r == nil {
		return nil
	}
	customArg := &foo.Custom{}
	_ = util.FromJSON(r.Custom, customArg)
	selfArg := &SelfCustom{}
	_ = util.FromJSON(r.Self, selfArg)
	return &Reference{
		ID:      r.ID,
		Custom:  customArg,
		Self:    selfArg,
		Created: r.Created,
	}
}

type rows []*row

func (x rows) ToReferences() References {
	return lo.Map(x, func(d *row, _ int) *Reference {
		return d.ToReference()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
