// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/har"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vhar"
	"github.com/kyleu/pftest/views/vpage"
)

func HarList(w http.ResponseWriter, r *http.Request) {
	controller.Act("har.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := as.Services.Har.List(ps.Logger)
		ps.SetTitleAndData("Archives", ret)
		return controller.Render(w, r, as, &vhar.List{Hars: ret}, ps, "har")
	})
}

func HarDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("har.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.PathString(r, "key", true)
		if err != nil {
			return "", err
		}
		ret, err := as.Services.Har.Load(key)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Archive ["+key+"]", ret)
		return controller.Render(w, r, as, &vhar.Detail{Har: ret}, ps, "har", ret.Key)
	})
}

func HarDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("har.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.PathString(r, "key", true)
		if err != nil {
			return "", err
		}
		err = as.Services.Har.Delete(key, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "Archive deleted", "/har", w, ps)
	})
}

func HarTrim(w http.ResponseWriter, r *http.Request) {
	controller.Act("har.trim", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.PathString(r, "key", true)
		if err != nil {
			return "", err
		}
		h, err := as.Services.Har.Load(key)
		if err != nil {
			return "", err
		}
		trimArgs := cutil.Args{
			{Key: "url", Title: "URL", Description: "matches against the URL (add \"*\" on either side to match wildcards)", Type: "string"},
			{Key: "mime", Title: "MIME", Description: "matches against the MIME type of the response", Type: "string", Choices: []string{"application/json"}},
		}
		argRes := cutil.CollectArgs(r, trimArgs)
		if argRes.HasMissing() {
			url := fmt.Sprintf("%s/trim", h.WebPath())
			ps.Data = argRes
			return controller.Render(w, r, as, &vpage.Args{URL: url, Directions: "Select the requests to trim", ArgRes: argRes}, ps, "har", h.Key, "Trim")
		}
		originalCount := len(h.Entries)
		h.Entries, err = h.Entries.Find(&har.Selector{URL: argRes.Values.GetStringOpt("url"), Mime: argRes.Values.GetStringOpt("mime")})
		if err != nil {
			return "", err
		}
		newCount := len(h.Entries)
		if newCount == originalCount {
			return controller.FlashAndRedir(true, "no changes needed", h.WebPath(), w, ps)
		}
		err = as.Services.Har.Save(h)
		if err != nil {
			return "", err
		}
		msg := fmt.Sprintf("Trimmed [%d] entries from archive", originalCount-newCount)
		return controller.FlashAndRedir(true, msg, h.WebPath(), w, ps)
	})
}

func HarUpload(w http.ResponseWriter, r *http.Request) {
	controller.Act("har.upload", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		if err := r.ParseMultipartForm(cutil.MaxBodySize); err != nil {
			return "", err
		}
		mpfrm := r.MultipartForm
		name := strings.Join(mpfrm.Value["n"], "")
		fileHeaders, ok := mpfrm.File["f"]
		if !ok {
			return "", errors.New("no file uploaded")
		}
		if len(fileHeaders) != 1 {
			return "", errors.New("invalid file uploads")
		}
		fileHeader := fileHeaders[0]
		file, err := fileHeader.Open()
		if err != nil {
			return "", err
		}
		if name == "" {
			name = fileHeader.Filename
			if !strings.HasSuffix(name, har.Ext) {
				name += har.Ext
			}
		}

		ps.Logger.Infof("Uploaded File: %+v\n", fileHeader.Filename)
		ps.Logger.Infof("File Size: %+v\n", fileHeader.Size)
		ps.Logger.Infof("MIME Header: %+v\n", fileHeader.Header)

		defer func() { _ = file.Close() }()
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			return "", err
		}
		ret := &har.Wrapper{}
		err = util.FromJSON(fileBytes, ret)
		if err != nil {
			return "", errors.Wrapf(err, "error decoding file [%s]", name)
		}
		ret.Log.Key = name
		err = as.Services.Har.Save(ret.Log)
		if err != nil {
			return "", err
		}
		msg := fmt.Sprintf("Created [%s] (%s)", name, util.ByteSizeSI(fileHeader.Size))
		redir := "/har/" + name
		return controller.FlashAndRedir(true, msg, redir, w, ps)
	})
}
