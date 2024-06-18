package reference

import "github.com/kyleu/pftest/app/util"

func (r *Reference) Diff(rx *Reference) util.Diffs {
	var diffs util.Diffs
	if r.ID != rx.ID {
		diffs = append(diffs, util.NewDiff("id", r.ID.String(), rx.ID.String()))
	}
	diffs = append(diffs, util.DiffObjects(r.Custom, rx.Custom, "custom")...)
	diffs = append(diffs, util.DiffObjects(r.Self, rx.Self, "self")...)
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}
