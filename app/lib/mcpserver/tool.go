package mcpserver

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/util"
)

type ToolHandler func(ctx context.Context, as *app.State, req mcp.CallToolRequest, args util.ValueMap, logger util.Logger) (string, error)

type Tool struct {
	Name        string
	Description string
	Args        util.FieldDescs
	Fn          ToolHandler
}

func (t Tool) ToMCP() (mcp.Tool, error) {
	opts := []mcp.ToolOption{mcp.WithDescription(t.Description)}
	for _, x := range t.Args {
		switch x.Type {
		case "string", "":
			opts = append(opts, mcp.WithString(x.Key, mcp.Required(), mcp.Description(x.Description), mcp.DefaultString(x.Default)))
		default:
			return mcp.Tool{}, errors.Errorf("unable to parse tool argument [%s] as type [%s]", x.Key, x.Type)
		}
	}
	return mcp.NewTool(t.Name, opts...), nil
}

func (t Tool) Handler(as *app.State, logger util.Logger) server.ToolHandlerFunc {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := util.ValueMapFrom(req.GetArguments())
		if t.Fn == nil {
			return nil, errors.Errorf("no handler for tool [%s]", t.Name)
		}
		ret, err := t.Fn(ctx, as, req, args, logger)
		if err != nil {
			return nil, errors.Errorf("errors running [%s] with arguments %s: %+v", t.Name, util.ToJSONCompact(args), err)
		}
		return mcp.NewToolResultText(ret), nil
	}
}

type Tools []*Tool
