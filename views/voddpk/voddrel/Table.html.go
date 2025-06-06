// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/voddpk/voddrel/Table.html:1
package voddrel

//line views/voddpk/voddrel/Table.html:1
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/oddpk/oddrel"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/components/view"
)

//line views/voddpk/voddrel/Table.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/voddpk/voddrel/Table.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/voddpk/voddrel/Table.html:10
func StreamTable(qw422016 *qt422016.Writer, models oddrel.Oddrels, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/voddpk/voddrel/Table.html:10
	qw422016.N().S(`
`)
//line views/voddpk/voddrel/Table.html:11
	prms := params.Sanitized("oddrel", ps.Logger)

//line views/voddpk/voddrel/Table.html:11
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/voddpk/voddrel/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "oddrel", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/voddpk/voddrel/Table.html:16
	qw422016.N().S(`
          `)
//line views/voddpk/voddrel/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "oddrel", "project", "Project", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/voddpk/voddrel/Table.html:17
	qw422016.N().S(`
          `)
//line views/voddpk/voddrel/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "oddrel", "path", "Path", "String text", prms, ps.URI, ps)
//line views/voddpk/voddrel/Table.html:18
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/voddpk/voddrel/Table.html:22
	for _, model := range models {
//line views/voddpk/voddrel/Table.html:22
		qw422016.N().S(`        <tr>
          <td><a href="`)
//line views/voddpk/voddrel/Table.html:24
		qw422016.E().S(model.WebPath(paths...))
//line views/voddpk/voddrel/Table.html:24
		qw422016.N().S(`">`)
//line views/voddpk/voddrel/Table.html:24
		view.StreamUUID(qw422016, &model.ID)
//line views/voddpk/voddrel/Table.html:24
		qw422016.N().S(`</a></td>
          <td>`)
//line views/voddpk/voddrel/Table.html:25
		view.StreamUUID(qw422016, &model.Project)
//line views/voddpk/voddrel/Table.html:25
		qw422016.N().S(`</td>
          <td>`)
//line views/voddpk/voddrel/Table.html:26
		view.StreamString(qw422016, model.Path)
//line views/voddpk/voddrel/Table.html:26
		qw422016.N().S(`</td>
        </tr>
`)
//line views/voddpk/voddrel/Table.html:28
	}
//line views/voddpk/voddrel/Table.html:28
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/voddpk/voddrel/Table.html:32
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/voddpk/voddrel/Table.html:32
		qw422016.N().S(`  <hr />
  `)
//line views/voddpk/voddrel/Table.html:34
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/voddpk/voddrel/Table.html:34
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/voddpk/voddrel/Table.html:36
	}
//line views/voddpk/voddrel/Table.html:37
}

//line views/voddpk/voddrel/Table.html:37
func WriteTable(qq422016 qtio422016.Writer, models oddrel.Oddrels, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/voddpk/voddrel/Table.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/voddpk/voddrel/Table.html:37
	StreamTable(qw422016, models, params, as, ps, paths...)
//line views/voddpk/voddrel/Table.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/voddpk/voddrel/Table.html:37
}

//line views/voddpk/voddrel/Table.html:37
func Table(models oddrel.Oddrels, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) string {
//line views/voddpk/voddrel/Table.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/voddpk/voddrel/Table.html:37
	WriteTable(qb422016, models, params, as, ps, paths...)
//line views/voddpk/voddrel/Table.html:37
	qs422016 := string(qb422016.B)
//line views/voddpk/voddrel/Table.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/voddpk/voddrel/Table.html:37
	return qs422016
//line views/voddpk/voddrel/Table.html:37
}
