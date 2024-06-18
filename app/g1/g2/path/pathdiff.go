package path

import "github.com/kyleu/pftest/app/util"

func (p *Path) Diff(px *Path) util.Diffs {
	var diffs util.Diffs
	if p.ID != px.ID {
		diffs = append(diffs, util.NewDiff("id", p.ID.String(), px.ID.String()))
	}
	if p.Name != px.Name {
		diffs = append(diffs, util.NewDiff("name", p.Name, px.Name))
	}
	if p.Status != px.Status {
		diffs = append(diffs, util.NewDiff("status", p.Status, px.Status))
	}
	if p.Created != px.Created {
		diffs = append(diffs, util.NewDiff("created", p.Created.String(), px.Created.String()))
	}
	return diffs
}
