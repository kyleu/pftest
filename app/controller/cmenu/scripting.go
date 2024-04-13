// Package cmenu - Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import (
	"net/url"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/menu"
	"github.com/kyleu/pftest/app/lib/scripting"
	"github.com/kyleu/pftest/app/util"
)

func scriptingMenu(s *scripting.Service, logger util.Logger) *menu.Item {
	scripts := s.ListScripts(logger)
	ret := make(menu.Items, 0, len(scripts))
	lo.ForEach(scripts, func(x string, _ int) {
		ret = append(ret, &menu.Item{Key: x, Title: x, Icon: "file", Route: "/admin/scripting/" + url.QueryEscape(x)})
	})
	desc := "script files managed by this system"
	return &menu.Item{Key: "scripting", Title: "Scripts", Description: desc, Icon: "cog", Route: "/admin/scripting", Children: ret}
}
