// Package timestamp - Content managed by Project Forge, see [projectforge.md] for details.
package timestamp

import (
	"fmt"

	"github.com/kyleu/pftest/app/util"
)

func (t *Timestamp) Diff(tx *Timestamp) util.Diffs {
	var diffs util.Diffs
	if t.ID != tx.ID {
		diffs = append(diffs, util.NewDiff("id", t.ID, tx.ID))
	}
	if t.Created != tx.Created {
		diffs = append(diffs, util.NewDiff("created", t.Created.String(), tx.Created.String()))
	}
	if (t.Deleted == nil && tx.Deleted != nil) || (t.Deleted != nil && tx.Deleted == nil) || (t.Deleted != nil && tx.Deleted != nil && *t.Deleted != *tx.Deleted) {
		diffs = append(diffs, util.NewDiff("deleted", fmt.Sprint(t.Deleted), fmt.Sprint(tx.Deleted))) //nolint:gocritic // it's nullable
	}
	return diffs
}
