package controller

import (
	"context"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/lib/menu"
	"github.com/kyleu/pftest/app/lib/sandbox"
	"github.com/kyleu/pftest/app/lib/telemetry"
	"github.com/kyleu/pftest/app/util"
)

func MenuFor(ctx context.Context, isAuthed bool, isAdmin bool, as *app.State) (menu.Items, error) {
	ctx, span := telemetry.StartSpan(ctx, "menu", "menu:generate")
	defer span.End()

	var ret menu.Items
	// $PF_SECTION_START(routes_start)$
	// $PF_SECTION_END(routes_start)$
	// $PF_INJECT_START(codegen)$
	ret = append(ret,
		&menu.Item{Key: "basic", Title: "Basics", Description: "Basic Model", Icon: "star", Route: "/basic"},
		&menu.Item{Key: "timestamp", Title: "Timestamps", Description: "Timestamps", Icon: "star", Route: "/timestamp"},
		&menu.Item{Key: "version", Title: "Versions", Description: "Versioned table", Icon: "star", Route: "/version"},
		&menu.Item{Key: "softdel", Title: "Softdels", Description: "Soft-deleted table", Icon: "star", Route: "/softdel"},
		&menu.Item{Key: "group", Title: "Groups", Description: "Grouped table", Icon: "star", Route: "/group", Children: menu.Items{
			&menu.Item{Key: "group", Title: "Groups", Description: "Groups from groups", Icon: "star", Route: "/group/group"},
		}},
		&menu.Item{Key: "mixed_case", Title: "MixedCases", Description: "Table and columns with mixed casing", Icon: "star", Route: "/mixed_case"},
	)
	// $PF_INJECT_END(codegen)$
	// $PF_SECTION_START(routes_end)$
	if isAdmin {
		ret = append(ret,
			sandbox.Menu(ctx),
			menu.Separator,
			&menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"},
		)
	}
	aboutDesc := "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(routes_end)$
	return ret, nil
}
