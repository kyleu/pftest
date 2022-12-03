// Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import "github.com/kyleu/pftest/app/lib/menu"

//nolint:lll
func generatedMenu() menu.Items {
	return menu.Items{
		&menu.Item{Key: "g1", Title: "G1", Icon: "star", Children: menu.Items{
			&menu.Item{Key: "g2", Title: "G2", Icon: "star", Children: menu.Items{
				&menu.Item{Key: "path", Title: "Paths", Description: "Path Model", Icon: "star", Route: "/g1/g2/path"},
			}},
		}},
		&menu.Item{Key: "capital", Title: "Capitals", Description: "Proper case table", Icon: "star", Route: "/capital"},
		&menu.Item{Key: "audited", Title: "Auditeds", Description: "Audited Model", Icon: "star", Route: "/audited"},
		&menu.Item{Key: "basic", Title: "Basics", Description: "Basic Model", Icon: "star", Route: "/basic"},
		&menu.Item{Key: "group", Title: "Groups", Description: "Grouped table", Icon: "star", Route: "/group", Children: menu.Items{
			&menu.Item{Key: "child", Title: "Children", Description: "Children from groups", Icon: "star", Route: "/group/child"},
		}},
		&menu.Item{Key: "hist", Title: "Hists", Description: "History table", Icon: "star", Route: "/hist"},
		&menu.Item{Key: "mixedcase", Title: "Mixed Cases", Description: "Table and columns with mixed casing", Icon: "star", Route: "/mixedcase"},
		&menu.Item{Key: "reference", Title: "References", Description: "Custom Reference", Icon: "star", Route: "/reference"},
		&menu.Item{Key: "relation", Title: "Relations", Description: "Relation Model", Icon: "star", Route: "/relation"},
		&menu.Item{Key: "seed", Title: "Seeds", Description: "Model with seed data", Icon: "star", Route: "/seed"},
		&menu.Item{Key: "softdel", Title: "Softdels", Description: "Soft-deleted table", Icon: "star", Route: "/softdel"},
		&menu.Item{Key: "timestamp", Title: "Timestamps", Description: "Timestamps", Icon: "star", Route: "/timestamp"},
		&menu.Item{Key: "trouble", Title: "Troubles", Description: "Columns with scary names", Icon: "star", Route: "/troub/le"},
		&menu.Item{Key: "version", Title: "Versions", Description: "Versioned table", Icon: "star", Route: "/version"},
	}
}
