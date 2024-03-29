// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vg1/vg2/vpath/Table.html:2
package vpath

//line views/vg1/vg2/vpath/Table.html:2
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/g1/g2/path"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/components/view"
)

//line views/vg1/vg2/vpath/Table.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vg1/vg2/vpath/Table.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vg1/vg2/vpath/Table.html:11
func StreamTable(qw422016 *qt422016.Writer, models path.Paths, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vg1/vg2/vpath/Table.html:11
	qw422016.N().S(`
`)
//line views/vg1/vg2/vpath/Table.html:12
	prms := params.Get("path", nil, ps.Logger).Sanitize("path")

//line views/vg1/vg2/vpath/Table.html:12
	qw422016.N().S(`  <table>
    <thead>
      <tr>
        `)
//line views/vg1/vg2/vpath/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "path", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vg1/vg2/vpath/Table.html:16
	qw422016.N().S(`
        `)
//line views/vg1/vg2/vpath/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "path", "name", "Name", "String text", prms, ps.URI, ps)
//line views/vg1/vg2/vpath/Table.html:17
	qw422016.N().S(`
        `)
//line views/vg1/vg2/vpath/Table.html:18
	components.StreamTableHeaderSimple(qw422016, "path", "status", "Status", "String text", prms, ps.URI, ps)
//line views/vg1/vg2/vpath/Table.html:18
	qw422016.N().S(`
        `)
//line views/vg1/vg2/vpath/Table.html:19
	components.StreamTableHeaderSimple(qw422016, "path", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vg1/vg2/vpath/Table.html:19
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vg1/vg2/vpath/Table.html:23
	for _, model := range models {
//line views/vg1/vg2/vpath/Table.html:23
		qw422016.N().S(`      <tr>
        <td><a href="/g1/g2/path/`)
//line views/vg1/vg2/vpath/Table.html:25
		view.StreamUUID(qw422016, &model.ID)
//line views/vg1/vg2/vpath/Table.html:25
		qw422016.N().S(`">`)
//line views/vg1/vg2/vpath/Table.html:25
		view.StreamUUID(qw422016, &model.ID)
//line views/vg1/vg2/vpath/Table.html:25
		qw422016.N().S(`</a></td>
        <td><strong>`)
//line views/vg1/vg2/vpath/Table.html:26
		view.StreamString(qw422016, model.Name)
//line views/vg1/vg2/vpath/Table.html:26
		qw422016.N().S(`</strong></td>
        <td><strong>`)
//line views/vg1/vg2/vpath/Table.html:27
		qw422016.E().S(model.Status)
//line views/vg1/vg2/vpath/Table.html:27
		qw422016.N().S(`</strong></td>
        <td>`)
//line views/vg1/vg2/vpath/Table.html:28
		view.StreamTimestamp(qw422016, &model.Created)
//line views/vg1/vg2/vpath/Table.html:28
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vg1/vg2/vpath/Table.html:30
	}
//line views/vg1/vg2/vpath/Table.html:31
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vg1/vg2/vpath/Table.html:31
		qw422016.N().S(`      <tr>
        <td colspan="4">`)
//line views/vg1/vg2/vpath/Table.html:33
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vg1/vg2/vpath/Table.html:33
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vg1/vg2/vpath/Table.html:35
	}
//line views/vg1/vg2/vpath/Table.html:35
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vg1/vg2/vpath/Table.html:38
}

//line views/vg1/vg2/vpath/Table.html:38
func WriteTable(qq422016 qtio422016.Writer, models path.Paths, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vg1/vg2/vpath/Table.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vg1/vg2/vpath/Table.html:38
	StreamTable(qw422016, models, params, as, ps)
//line views/vg1/vg2/vpath/Table.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/vg1/vg2/vpath/Table.html:38
}

//line views/vg1/vg2/vpath/Table.html:38
func Table(models path.Paths, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vg1/vg2/vpath/Table.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vg1/vg2/vpath/Table.html:38
	WriteTable(qb422016, models, params, as, ps)
//line views/vg1/vg2/vpath/Table.html:38
	qs422016 := string(qb422016.B)
//line views/vg1/vg2/vpath/Table.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vg1/vg2/vpath/Table.html:38
	return qs422016
//line views/vg1/vg2/vpath/Table.html:38
}
