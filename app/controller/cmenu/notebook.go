// Package cmenu - Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import (
	"context"

	"github.com/kyleu/pftest/app/lib/menu"
)

func notebookMenu(_ context.Context) *menu.Item {
	return &menu.Item{Key: "notebook", Title: "Notebook", Description: "Notebook", Icon: "notebook", Route: "/notebook"}
}
