package clib

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/doc"
	"github.com/kyleu/pftest/views/vdoc"
)

func Docs(w http.ResponseWriter, r *http.Request) {
	controller.Act("docs", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		pth, _ := cutil.PathString(r, "path", false)
		if pth == "" {
			return "", errors.New("invalid path")
		}

		bc := []string{"docs"}
		bc = append(bc, util.StringSplitAndTrim(pth, "/")...)

		title, x, err := doc.HTML("doc:"+pth, pth+util.ExtMarkdown, func(s string) (string, string, error) {
			return cutil.FormatMarkdownClean(s, "file")
		})
		if err != nil {
			return "", errors.Wrapf(err, "unable to load documentation from [%s]", pth)
		}
		c, _ := doc.Content(pth + util.ExtMarkdown)
		ps.SetTitleAndData(title, c)
		return controller.Render(r, as, &vdoc.MarkdownPage{Title: pth, HTML: x}, ps, bc...)
	})
}
