package trouble

import (
	"fmt"

	"github.com/kyleu/pftest/app/util"
)

func (t *Trouble) Diff(tx *Trouble) util.Diffs {
	var diffs util.Diffs
	if t.From != tx.From {
		diffs = append(diffs, util.NewDiff("from", t.From, tx.From))
	}
	diffs = append(diffs, util.DiffObjects(t.Where, tx.Where, "where")...)
	if t.Selectcol != tx.Selectcol {
		diffs = append(diffs, util.NewDiff("selectcol", fmt.Sprint(t.Selectcol), fmt.Sprint(tx.Selectcol)))
	}
	if t.Limit != tx.Limit {
		diffs = append(diffs, util.NewDiff("limit", t.Limit, tx.Limit))
	}
	if t.Group != tx.Group {
		diffs = append(diffs, util.NewDiff("group", t.Group, tx.Group))
	}
	if (t.Delete == nil && tx.Delete != nil) || (t.Delete != nil && tx.Delete == nil) || (t.Delete != nil && tx.Delete != nil && *t.Delete != *tx.Delete) {
		diffs = append(diffs, util.NewDiff("delete", fmt.Sprint(t.Delete), fmt.Sprint(tx.Delete))) //nolint:gocritic // it's nullable
	}
	return diffs
}
