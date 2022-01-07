package controller

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/sandbox"
	"github.com/kyleu/pftest/views/vsandbox"
)

func SandboxList(rc *fasthttp.RequestCtx) {
	act("sandbox.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Sandboxes"
		ps.Data = sandbox.AllSandboxes
		return render(rc, as, &vsandbox.List{}, ps, "sandbox")
	})
}

func SandboxRun(rc *fasthttp.RequestCtx) {
	act("sandbox.run", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := rcRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		sb := sandbox.AllSandboxes.Get(key)
		if sb == nil {
			return ersp("no sandbox with key [%s]", key)
		}
		ret, err := sb.Run(ps.Context, as, ps.Logger.With(zap.String("sandbox", key)))
		if err != nil {
			return "", err
		}
		ps.Title = sb.Title
		ps.Data = ret
		if sb.Key == "testbed" {
			return render(rc, as, &vsandbox.Testbed{}, ps, "sandbox", sb.Key)
		}
		return render(rc, as, &vsandbox.Run{Key: key, Title: sb.Title, Result: ret}, ps, "sandbox", sb.Key)
	})
}
