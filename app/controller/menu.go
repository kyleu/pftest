// Content managed by Project Forge, see [projectforge.md] for details.
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
	_, span, _ := telemetry.StartSpan(ctx, "menu:generate", nil)
	defer span.Complete()

	var ret menu.Items
	// $PF_SECTION_START(routes_start)$
	// $PF_SECTION_END(routes_start)$
	// $PF_INJECT_START(codegen)$
	ret = append(ret,
		&menu.Item{Key: "basic", Title: "Basics", Description: "Basic Model", Icon: "star", Route: "/basic"},
		&menu.Item{Key: "reference", Title: "References", Description: "Custom Reference", Icon: "star", Route: "/reference"},
		&menu.Item{Key: "audited", Title: "Auditeds", Description: "Audited Model", Icon: "star", Route: "/audited"},
		&menu.Item{Key: "timestamp", Title: "Timestamps", Description: "Timestamps", Icon: "star", Route: "/timestamp"},
		&menu.Item{Key: "version", Title: "Versions", Description: "Versioned table", Icon: "star", Route: "/version"},
		&menu.Item{Key: "history", Title: "Histories", Description: "History table", Icon: "star", Route: "/history"},
		&menu.Item{Key: "softdel", Title: "Softdels", Description: "Soft-deleted table", Icon: "star", Route: "/softdel"},
		&menu.Item{Key: "group", Title: "Groups", Description: "Grouped table", Icon: "star", Route: "/group", Children: menu.Items{
			&menu.Item{Key: "group", Title: "Groups", Description: "Groups from groups", Icon: "star", Route: "/group/group"},
		}},
		&menu.Item{Key: "capital", Title: "Capitals", Description: "Proper case table", Icon: "star", Route: "/capital"},
		&menu.Item{Key: "mixedcase", Title: "Mixed Cases", Description: "Table and columns with mixed casing", Icon: "star", Route: "/mixedcase"},
		&menu.Item{Key: "trouble", Title: "Troubles", Description: "Columns with scary names", Icon: "star", Route: "/troub/le"},
	)
	// $PF_INJECT_END(codegen)$
	// $PF_SECTION_START(routes_end)$
	if isAdmin {
		ret = append(ret,
			menu.Separator,
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
