// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"net/http"
	"strings"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/site"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/verror"
)

func Site(w http.ResponseWriter, r *http.Request) {
	path := util.StringSplitAndTrim(string(r.URL.Path), "/")
	action := "site"
	if len(path) > 0 {
		action += "." + strings.Join(path, ".")
	}
	ActSite(action, w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		redir, page, bc, err := site.Handle(path, as, ps)
		if err != nil {
			return "", err
		}
		if _, ok := page.(*verror.NotFound); ok {
			w.WriteHeader(http.StatusNotFound)
		}
		if redir != "" {
			return redir, nil
		}
		return Render(w, r, as, page, ps, bc...)
	})
}
