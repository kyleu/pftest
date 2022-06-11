// Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import (
	"context"

	"github.com/kyleu/pftest/app/lib/graphql"
	"github.com/kyleu/pftest/app/lib/menu"
)

func graphQLMenu(gqlSvc *graphql.Service, ctx context.Context) *menu.Item {
	l := gqlSvc.Keys()
	kids := make(menu.Items, 0, len(l))
	titles := gqlSvc.Titles()
	if len(l) > 1 {
		for _, x := range l {
			kids = append(kids, &menu.Item{Key: x, Title: titles[x], Description: "A GraphQL schema", Icon: "graph", Route: "/graphql/" + x})
		}
	}
	return &menu.Item{Key: "graphql", Title: "GraphQL", Description: "A graph-based API", Icon: "graph", Route: "/graphql", Children: kids}
}
