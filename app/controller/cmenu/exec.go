package cmenu

import (
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/exec"
	"github.com/kyleu/pftest/app/lib/menu"
)

func processMenu(processes exec.Execs) *menu.Item {
	ret := make(menu.Items, 0, len(processes))
	lo.ForEach(processes, func(p *exec.Exec, _ int) {
		title := p.String()
		if p.Completed != nil {
			title += "*"
		}
		ret = append(ret, &menu.Item{Key: p.String(), Title: title, Icon: "bolt", Description: p.String(), Route: p.WebPath()})
	})
	desc := "process executions managed by this system"
	return &menu.Item{Key: "exec", Title: "Processes", Description: desc, Icon: "desktop", Route: "/admin/exec", Children: ret}
}
