// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vtimestamp/Table.html:1
package vtimestamp

//line views/vtimestamp/Table.html:1
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/timestamp"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/components/view"
)

//line views/vtimestamp/Table.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtimestamp/Table.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtimestamp/Table.html:10
func StreamTable(qw422016 *qt422016.Writer, models timestamp.Timestamps, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vtimestamp/Table.html:10
	qw422016.N().S(`
`)
//line views/vtimestamp/Table.html:11
	prms := params.Sanitized("timestamp", ps.Logger)

//line views/vtimestamp/Table.html:11
	qw422016.N().S(`  <div class="overflow clear">
    <table>
      <thead>
        <tr>
          `)
//line views/vtimestamp/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "timestamp", "id", "ID", "String text", prms, ps.URI, ps)
//line views/vtimestamp/Table.html:16
	qw422016.N().S(`
          `)
//line views/vtimestamp/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "timestamp", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vtimestamp/Table.html:17
	qw422016.N().S(`
          `)
//line views/vtimestamp/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "timestamp", "updated", "Updated", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vtimestamp/Table.html:18
	qw422016.N().S(`
          `)
//line views/vtimestamp/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "timestamp", "deleted", "Deleted", "Date and time, in almost any format (optional)", prms, ps.URI, ps)
//line views/vtimestamp/Table.html:19
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vtimestamp/Table.html:23
	for _, model := range models {
//line views/vtimestamp/Table.html:23
		qw422016.N().S(`        <tr>
          <td><a href="`)
//line views/vtimestamp/Table.html:25
		qw422016.E().S(model.WebPath(paths...))
//line views/vtimestamp/Table.html:25
		qw422016.N().S(`">`)
//line views/vtimestamp/Table.html:25
		view.StreamString(qw422016, model.ID)
//line views/vtimestamp/Table.html:25
		qw422016.N().S(`</a></td>
          <td>`)
//line views/vtimestamp/Table.html:26
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vtimestamp/Table.html:26
		qw422016.N().S(`</td>
          <td>`)
//line views/vtimestamp/Table.html:27
		view.StreamTimestamp(qw422016, model.Updated)
//line views/vtimestamp/Table.html:27
		qw422016.N().S(`</td>
          <td>`)
//line views/vtimestamp/Table.html:28
		view.StreamTimestamp(qw422016, model.Deleted)
//line views/vtimestamp/Table.html:28
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vtimestamp/Table.html:30
	}
//line views/vtimestamp/Table.html:30
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vtimestamp/Table.html:34
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vtimestamp/Table.html:34
		qw422016.N().S(`  <hr />
  `)
//line views/vtimestamp/Table.html:36
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vtimestamp/Table.html:36
		qw422016.N().S(`
  <div class="clear"></div>
`)
//line views/vtimestamp/Table.html:38
	}
//line views/vtimestamp/Table.html:39
}

//line views/vtimestamp/Table.html:39
func WriteTable(qq422016 qtio422016.Writer, models timestamp.Timestamps, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) {
//line views/vtimestamp/Table.html:39
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtimestamp/Table.html:39
	StreamTable(qw422016, models, params, as, ps, paths...)
//line views/vtimestamp/Table.html:39
	qt422016.ReleaseWriter(qw422016)
//line views/vtimestamp/Table.html:39
}

//line views/vtimestamp/Table.html:39
func Table(models timestamp.Timestamps, params filter.ParamSet, as *app.State, ps *cutil.PageState, paths ...string) string {
//line views/vtimestamp/Table.html:39
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtimestamp/Table.html:39
	WriteTable(qb422016, models, params, as, ps, paths...)
//line views/vtimestamp/Table.html:39
	qs422016 := string(qb422016.B)
//line views/vtimestamp/Table.html:39
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtimestamp/Table.html:39
	return qs422016
//line views/vtimestamp/Table.html:39
}
