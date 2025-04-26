package cmenu

import "github.com/kyleu/pftest/app/lib/menu"

//nolint:lll
var (
	menuItemAudited     = &menu.Item{Key: "audited", Title: "Auditeds", Description: "Audited Model", Icon: "star", Route: "/audited"}
	menuItemBasic       = &menu.Item{Key: "basic", Title: "Basics", Description: "Basic Model", Icon: "star", Route: "/basic"}
	menuItemCapital     = &menu.Item{Key: "capital", Title: "Capitals", Description: "Proper case table", Icon: "star", Route: "/capital"}
	menuItemG2Path      = &menu.Item{Key: "path", Title: "Paths", Description: "Path Model", Icon: "star", Route: "/g1/g2/path"}
	menuItemMixedCase   = &menu.Item{Key: "mixedcase", Title: "Mixed Cases", Description: "Table and columns with mixed casing", Icon: "star", Route: "/mixedcase"}
	menuItemOddPK       = &menu.Item{Key: "oddpk", Title: "Odd PKs", Description: "Odd PK", Icon: "star", Route: "/oddpk", Children: menu.Items{menuItemOddPKOddrel}}
	menuItemOddPKOddrel = &menu.Item{Key: "oddrel", Title: "Oddrels", Description: "Odd Rel", Icon: "star", Route: "/oddpk/oddrel"}
	menuItemReference   = &menu.Item{Key: "reference", Title: "References", Description: "Custom Reference", Icon: "star", Route: "/reference"}
	menuItemRelation    = &menu.Item{Key: "relation", Title: "Relations", Description: "Relation Model", Icon: "star", Route: "/relation"}
	menuItemSeed        = &menu.Item{Key: "seed", Title: "Seeds", Description: "Model with seed data", Icon: "star", Route: "/seed"}
	menuItemSoftdel     = &menu.Item{Key: "softdel", Title: "Softdels", Description: "Soft-deleted table", Icon: "star", Route: "/softdel"}
	menuItemTimestamp   = &menu.Item{Key: "timestamp", Title: "Timestamps", Description: "Timestamps", Icon: "star", Route: "/timestamp"}
	menuItemTrouble     = &menu.Item{Key: "trouble", Title: "Troubles", Description: "Columns with scary names", Icon: "star", Route: "/troub/le"}

	menuGroupG1 = &menu.Item{Key: "g1", Title: "g1", Children: menu.Items{menuGroupG2}}
	menuGroupG2 = &menu.Item{Key: "g2", Title: "g2", Children: menu.Items{menuItemG2Path}}
)

func generatedMenu() menu.Items {
	return menu.Items{
		menuGroupG1,
		menuItemCapital,
		menuItemAudited,
		menuItemBasic,
		menuItemMixedCase,
		menuItemOddPK,
		menuItemReference,
		menuItemRelation,
		menuItemSeed,
		menuItemSoftdel,
		menuItemTimestamp,
		menuItemTrouble,
	}
}
