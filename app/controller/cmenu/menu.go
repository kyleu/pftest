// Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import (
	"context"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/lib/menu"
	"github.com/kyleu/pftest/app/lib/sandbox"
	"github.com/kyleu/pftest/app/lib/telemetry"
	"github.com/kyleu/pftest/app/util"
)

func MenuFor(ctx context.Context, isAuthed bool, isAdmin bool, as *app.State, logger util.Logger) (menu.Items, error) {
	ctx, span, logger := telemetry.StartSpan(ctx, "menu:generate", logger)
	defer span.Complete()
	_ = logger

	var ret menu.Items
	// $PF_SECTION_START(routes_start)$
	// $PF_SECTION_END(routes_start)$
	if isAdmin {
		ret = append(ret, generatedMenu()...)
	}
	// $PF_SECTION_START(routes_end)$
	if isAdmin {
		admin := &menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"}
		ret = append(ret, menu.Separator, graphQLMenu(as.GraphQL, ctx), sandbox.Menu(ctx), menu.Separator, admin, menu.Separator, docMenu(ctx, as, logger))
	}
	const aboutDesc = "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(routes_end)$
	return ret, nil
}
