// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/vsql"
)

func SQLEditor(rc *fasthttp.RequestCtx) {
	act("sql.editor", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "SQL Editor"
		ps.Data = "Post to this action with [sql] in the body"
		return render(rc, as, &vsql.SQLEditor{}, ps, "sql")
	})
}

func SQLRun(rc *fasthttp.RequestCtx) {
	act("sql.run", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		sql := "select 1;"
		commit := false
		meth := string(rc.Method())
		if strings.EqualFold(meth, "post") {
			f := rc.PostArgs()
			sql = string(f.Peek("sql"))
			c := string(f.Peek("commit"))
			commit = c == util.BoolTrue
			action := string(f.Peek("action"))
			if action == "analyze" {
				sql = "explain analyze " + sql
			}
		}

		tx, err := as.DB.StartTransaction()
		if err != nil {
			return "", errors.Wrap(err, "unable to start transaction")
		}

		var columns []string
		results := [][]any{}

		timer := util.TimerStart()
		result, err := as.DB.Query(ps.Context, sql, tx)
		if err != nil {
			return "", err
		}
		defer func() { _ = result.Close() }()

		elapsed := timer.End()

		if result != nil {
			for result.Next() {
				if columns == nil {
					columns, _ = result.Columns()
				}
				row, e := result.SliceScan()
				if e != nil {
					return "", errors.Wrap(e, "unable to read row")
				}
				results = append(results, row)
			}
		}

		if commit {
			err = tx.Commit()
			if err != nil {
				return "", errors.Wrap(err, "unable to commit transaction")
			}
		} else {
			_ = tx.Rollback()
		}

		ps.Title = "SQL Results"
		ps.Data = results
		return render(rc, as, &vsql.SQLEditor{SQL: sql, Columns: columns, Results: results, Timing: elapsed, Commit: commit}, ps, "sql")
	})
}
