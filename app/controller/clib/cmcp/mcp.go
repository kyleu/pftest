package cmcp

import (
	"net/http"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/views/vmcp"
)

const mcpBreadcrumb = "mcp"

func MCPIndex(w http.ResponseWriter, r *http.Request) {
	controller.Act("mcp.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		mcpx, _, err := mcpTool(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("MCP", mcpx)
		return controller.Render(r, as, &vmcp.MCPList{Server: mcpx}, ps, mcpBreadcrumb)
	})
}

func MCPServe(w http.ResponseWriter, r *http.Request) {
	controller.Act("mcp.serve.streamable", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		mcpx, _, err := mcpTool(r, as, ps)
		if err != nil {
			return "", err
		}
		mcpx.ServeHTTP(ps.Context, w, r, ps.Logger)
		return "", nil
	})
}
