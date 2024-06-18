package capital

import (
	"fmt"

	"github.com/kyleu/pftest/app/util"
)

//nolint:lll
func (c *Capital) Diff(cx *Capital) util.Diffs {
	var diffs util.Diffs
	if c.ID != cx.ID {
		diffs = append(diffs, util.NewDiff("id", c.ID, cx.ID))
	}
	if c.Name != cx.Name {
		diffs = append(diffs, util.NewDiff("name", c.Name, cx.Name))
	}
	if c.Birthday != cx.Birthday {
		diffs = append(diffs, util.NewDiff("birthday", c.Birthday.String(), cx.Birthday.String()))
	}
	if (c.Deathday == nil && cx.Deathday != nil) || (c.Deathday != nil && cx.Deathday == nil) || (c.Deathday != nil && cx.Deathday != nil && *c.Deathday != *cx.Deathday) {
		diffs = append(diffs, util.NewDiff("deathday", fmt.Sprint(c.Deathday), fmt.Sprint(cx.Deathday))) //nolint:gocritic // it's nullable
	}
	return diffs
}
