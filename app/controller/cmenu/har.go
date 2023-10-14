// Package cmenu - Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import (
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/har"
	"github.com/kyleu/pftest/app/lib/menu"
)

func harMenu(s *har.Service) *menu.Item {
	harKids := lo.Map(s.List(nil), func(n string, _ int) *menu.Item {
		n = strings.TrimSuffix(n, har.Ext)
		return &menu.Item{Key: n, Title: n, Icon: "book", Route: "/har/" + n}
	})
	return &menu.Item{Key: "har", Title: "Archives", Description: "HTTP Archive files", Icon: "book", Route: "/har", Children: harKids}
}
