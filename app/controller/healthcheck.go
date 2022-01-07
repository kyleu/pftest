package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app/controller/cutil"
)

func Healthcheck(rc *fasthttp.RequestCtx) {
	x := map[string]string{"status": "OK"}
	_, _ = cutil.RespondJSON(rc, "", x)
}
