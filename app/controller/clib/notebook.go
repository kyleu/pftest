// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/filesystem"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vnotebook"
)

var NotebookBaseURL = fmt.Sprintf("http://localhost:%d", util.AppPort+10)

func Notebook(rc *fasthttp.RequestCtx) {
	controller.Act("notebook.view", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		if !ps.Admin {
			return controller.Unauthorized(rc, "", ps.Accounts)(as, ps)
		}
		status := as.Services.Notebook.Status()
		if status == "running" {
			pathS, _ := cutil.RCRequiredString(rc, "path", false)
			path := util.StringSplitAndTrim(pathS, "/")
			ps.SetTitleAndData("Notebook", path)
			return controller.Render(rc, as, &vnotebook.Notebook{BaseURL: NotebookBaseURL, Path: pathS}, ps, "notebook", pathS)
		}
		ps.SetTitleAndData("Notebook Options", status)
		return controller.Render(rc, as, &vnotebook.Options{}, ps)
	})
}

func NotebookFiles(rc *fasthttp.RequestCtx) {
	controller.Act("notebook.files", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		if !ps.Admin {
			return controller.Unauthorized(rc, "", ps.Accounts)(as, ps)
		}
		pathS, path, bc := notebookGetText(rc)
		fs := as.Services.Notebook.FS
		files, err := fs.ListTree(nil, pathS, []string{"cache"}, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "error listing files")
		}
		ps.SetTitleAndData(fmt.Sprintf("Notebook File /%s", pathS), files)
		return controller.Render(rc, as, &vnotebook.Files{Path: path, FS: as.Services.Notebook.FS}, ps, bc...)
	})
}

func NotebookFileEdit(rc *fasthttp.RequestCtx) {
	controller.Act("notebook.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		if !ps.Admin {
			return controller.Unauthorized(rc, "", ps.Accounts)(as, ps)
		}
		pathS, path, bc := notebookGetText(rc)
		bc = append(bc, "Edit**edit")
		b, err := as.Services.Notebook.FS.ReadFile(pathS)
		if err != nil {
			return "", err
		}
		ret := string(b)
		ps.SetTitleAndData(fmt.Sprintf("Notebook File /%s", pathS), path)
		return controller.Render(rc, as, &vnotebook.FileEdit{Path: path, Content: ret}, ps, bc...)
	})
}

func NotebookFileSave(rc *fasthttp.RequestCtx) {
	controller.Act("notebook.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		if !ps.Admin {
			return controller.Unauthorized(rc, "", ps.Accounts)(as, ps)
		}
		pathS, _, _ := notebookGetText(rc)
		msg := "Notebook file saved"

		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}
		content := frm.GetStringOpt("content")
		if strings.TrimSpace(content) == "" {
			return "", errors.Errorf("file [%s] contents may not be empty", pathS)
		}
		content = strings.ReplaceAll(content, "\r\n", "\n")

		b, err := as.Services.Notebook.FS.ReadFile(pathS)
		if err != nil {
			return "", err
		}
		current := string(b)

		if current == content {
			msg = "No changes required"
		} else {
			err = as.Services.Notebook.FS.WriteFile(pathS, []byte(content), filesystem.DefaultMode, true)
			if err != nil {
				return "", err
			}
		}
		return controller.FlashAndRedir(true, msg, "/notebook/files/"+pathS, rc, ps)
	})
}

func NotebookAction(rc *fasthttp.RequestCtx) {
	controller.Act("notebook.action", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		act, err := cutil.RCRequiredString(rc, "act", false)
		if err != nil {
			return "", err
		}
		switch act {
		case "start":
			err = as.Services.Notebook.Start(as.Services.Exec)
			return controller.FlashAndRedir(true, "Notebook started", "/notebook", rc, ps)
		default:
			return "", errors.Errorf("invalid notebook action [%s]", act)
		}
	})
}

func notebookGetText(rc *fasthttp.RequestCtx) (string, []string, []string) {
	pathS, _ := cutil.RCRequiredString(rc, "path", false)
	path := util.StringSplitAndTrim(pathS, "/")
	bcAppend := "||/notebook/files"
	bc := []string{"notebook", "files"}
	lo.ForEach(path, func(x string, _ int) {
		bcAppend += "/" + x
		b := x + bcAppend
		bc = append(bc, b)
	})
	return pathS, path, bc
}
