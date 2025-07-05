package cmenu

import (
	"context"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/lib/menu"
	"github.com/kyleu/pftest/app/lib/sandbox"
	"github.com/kyleu/pftest/app/lib/telemetry"
	"github.com/kyleu/pftest/app/lib/user"
	"github.com/kyleu/pftest/app/util"
)

func MenuFor(
	ctx context.Context, isAuthed bool, isAdmin bool, profile *user.Profile, params filter.ParamSet, as *app.State, logger util.Logger,
) (menu.Items, any, error) {
	ctx, sp, _ := telemetry.StartSpan(ctx, "menu", logger)
	defer sp.Complete()
	var ret menu.Items
	var data any
	// $PF_SECTION_START(menu)$
	if isAdmin {
		ret = append(ret, processMenu(as.Services.Exec.Execs), harMenu(as.Services.Har), scriptingMenu(as.Services.Script, logger))
	}
	if isAdmin {
		ret = append(ret, generatedMenu()...)
	}
	if isAdmin {
		admin := &menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"}
		ret = append(ret,
			menu.Separator, notebookMenu(ctx, isAdmin), graphQLMenu(ctx, as.GraphQL), sandbox.Menu(ctx),
			menu.Separator, admin, menu.Separator, mcpMenu(), docMenu(logger),
		)
	}
	const aboutDesc = "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(menu)$
	return ret, data, nil
}
