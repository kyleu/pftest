// Code generated by qtc from "RecordDetail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vaudit/RecordDetail.html:2
package vaudit

//line views/vaudit/RecordDetail.html:2
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/audit"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/layout"
)

//line views/vaudit/RecordDetail.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaudit/RecordDetail.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaudit/RecordDetail.html:10
type RecordDetail struct {
	layout.Basic
	Model *audit.Record
}

//line views/vaudit/RecordDetail.html:15
func (p *RecordDetail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaudit/RecordDetail.html:15
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-record"><button type="button">JSON</button></a>
      <a href="`)
//line views/vaudit/RecordDetail.html:19
	qw422016.E().S(p.Model.WebPath())
//line views/vaudit/RecordDetail.html:19
	qw422016.N().S(`/edit"><button>Edit</button></a>
    </div>
    <h3>`)
//line views/vaudit/RecordDetail.html:21
	components.StreamSVGRefIcon(qw422016, `cog`, ps)
//line views/vaudit/RecordDetail.html:21
	qw422016.N().S(` Audit Record [`)
//line views/vaudit/RecordDetail.html:21
	qw422016.E().S(p.Model.String())
//line views/vaudit/RecordDetail.html:21
	qw422016.N().S(`]</h3>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
          <td>`)
//line views/vaudit/RecordDetail.html:26
	components.StreamDisplayUUID(qw422016, &p.Model.ID)
//line views/vaudit/RecordDetail.html:26
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Audit ID</th>
          <td>
            <div class="icon">`)
//line views/vaudit/RecordDetail.html:31
	components.StreamDisplayUUID(qw422016, &p.Model.AuditID)
//line views/vaudit/RecordDetail.html:31
	qw422016.N().S(`</div>
            <a title="Audit" href="`)
//line views/vaudit/RecordDetail.html:32
	qw422016.E().S(`/audit` + `/` + p.Model.AuditID.String())
//line views/vaudit/RecordDetail.html:32
	qw422016.N().S(`">`)
//line views/vaudit/RecordDetail.html:32
	components.StreamSVGRefIcon(qw422016, "cog", ps)
//line views/vaudit/RecordDetail.html:32
	qw422016.N().S(`</a>
          </td>
        </tr>
        <tr>
          <th class="shrink" title="String text">T</th>
          <td>`)
//line views/vaudit/RecordDetail.html:37
	qw422016.E().S(p.Model.T)
//line views/vaudit/RecordDetail.html:37
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Pk</th>
          <td>`)
//line views/vaudit/RecordDetail.html:41
	qw422016.E().S(p.Model.PK)
//line views/vaudit/RecordDetail.html:41
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="JSON object">Changes</th>
          <td>`)
//line views/vaudit/RecordDetail.html:45
	components.StreamDisplayDiffs(qw422016, p.Model.Changes)
//line views/vaudit/RecordDetail.html:45
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="JSON object">Metadata</th>
          <td>`)
//line views/vaudit/RecordDetail.html:49
	components.StreamJSON(qw422016, p.Model.Metadata)
//line views/vaudit/RecordDetail.html:49
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Occurred</th>
          <td>`)
//line views/vaudit/RecordDetail.html:53
	components.StreamDisplayTimestamp(qw422016, &p.Model.Occurred)
//line views/vaudit/RecordDetail.html:53
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vaudit/RecordDetail.html:59
	qw422016.N().S(`  `)
//line views/vaudit/RecordDetail.html:60
	components.StreamJSONModal(qw422016, "record", "Audit Record JSON", p.Model, 1)
//line views/vaudit/RecordDetail.html:60
	qw422016.N().S(`
`)
//line views/vaudit/RecordDetail.html:61
}

//line views/vaudit/RecordDetail.html:61
func (p *RecordDetail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaudit/RecordDetail.html:61
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaudit/RecordDetail.html:61
	p.StreamBody(qw422016, as, ps)
//line views/vaudit/RecordDetail.html:61
	qt422016.ReleaseWriter(qw422016)
//line views/vaudit/RecordDetail.html:61
}

//line views/vaudit/RecordDetail.html:61
func (p *RecordDetail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaudit/RecordDetail.html:61
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaudit/RecordDetail.html:61
	p.WriteBody(qb422016, as, ps)
//line views/vaudit/RecordDetail.html:61
	qs422016 := string(qb422016.B)
//line views/vaudit/RecordDetail.html:61
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaudit/RecordDetail.html:61
	return qs422016
//line views/vaudit/RecordDetail.html:61
}