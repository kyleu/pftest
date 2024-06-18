package relation

import "github.com/kyleu/pftest/app/util"

func (r *Relation) Diff(rx *Relation) util.Diffs {
	var diffs util.Diffs
	if r.ID != rx.ID {
		diffs = append(diffs, util.NewDiff("id", r.ID.String(), rx.ID.String()))
	}
	if r.BasicID != rx.BasicID {
		diffs = append(diffs, util.NewDiff("basicID", r.BasicID.String(), rx.BasicID.String()))
	}
	if r.Name != rx.Name {
		diffs = append(diffs, util.NewDiff("name", r.Name, rx.Name))
	}
	if r.Created != rx.Created {
		diffs = append(diffs, util.NewDiff("created", r.Created.String(), rx.Created.String()))
	}
	return diffs
}
