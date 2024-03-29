// Code generated by qtc from "RecordTable.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vaudit/RecordTable.html:2
package vaudit

//line views/vaudit/RecordTable.html:2
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/audit"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/components/view"
)

//line views/vaudit/RecordTable.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaudit/RecordTable.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaudit/RecordTable.html:11
func StreamRecordTable(qw422016 *qt422016.Writer, models audit.Records, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vaudit/RecordTable.html:11
	qw422016.N().S(`
`)
//line views/vaudit/RecordTable.html:12
	prms := params.Get("audit_record", nil, ps.Logger)

//line views/vaudit/RecordTable.html:12
	qw422016.N().S(`  <div class="overflow full-width">
    <table class="mt">
      <thead>
        <tr>
          `)
//line views/vaudit/RecordTable.html:17
	components.StreamTableHeaderSimple(qw422016, "audit_record", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vaudit/RecordTable.html:17
	qw422016.N().S(`
          `)
//line views/vaudit/RecordTable.html:18
	components.StreamTableHeaderSimple(qw422016, "audit_record", "audit_id", "Audit ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vaudit/RecordTable.html:18
	qw422016.N().S(`
          `)
//line views/vaudit/RecordTable.html:19
	components.StreamTableHeaderSimple(qw422016, "audit_record", "t", "T", "Type of the target object", prms, ps.URI, ps)
//line views/vaudit/RecordTable.html:19
	qw422016.N().S(`
          `)
//line views/vaudit/RecordTable.html:20
	components.StreamTableHeaderSimple(qw422016, "audit_record", "pk", "Pk", "Primary key of the target object", prms, ps.URI, ps)
//line views/vaudit/RecordTable.html:20
	qw422016.N().S(`
          `)
//line views/vaudit/RecordTable.html:21
	components.StreamTableHeaderSimple(qw422016, "audit_record", "changes", "Changes", "Count of change", prms, ps.URI, ps)
//line views/vaudit/RecordTable.html:21
	qw422016.N().S(`
          `)
//line views/vaudit/RecordTable.html:22
	components.StreamTableHeaderSimple(qw422016, "audit_record", "metadata", "Metadata", "Count of metadata fields", prms, ps.URI, ps)
//line views/vaudit/RecordTable.html:22
	qw422016.N().S(`
          `)
//line views/vaudit/RecordTable.html:23
	components.StreamTableHeaderSimple(qw422016, "audit_record", "occurred", "Occurred", "Timestamp representing the time the event occurred", prms, ps.URI, ps)
//line views/vaudit/RecordTable.html:23
	qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vaudit/RecordTable.html:27
	for _, model := range models {
//line views/vaudit/RecordTable.html:27
		qw422016.N().S(`        <tr>
          <td><a href="/admin/audit/record/`)
//line views/vaudit/RecordTable.html:29
		view.StreamUUID(qw422016, &model.ID)
//line views/vaudit/RecordTable.html:29
		qw422016.N().S(`">`)
//line views/vaudit/RecordTable.html:29
		view.StreamUUID(qw422016, &model.ID)
//line views/vaudit/RecordTable.html:29
		qw422016.N().S(`</a></td>
          <td>
            <div class="icon">`)
//line views/vaudit/RecordTable.html:31
		view.StreamUUID(qw422016, &model.AuditID)
//line views/vaudit/RecordTable.html:31
		qw422016.N().S(`</div>
            <a title="Audit" href="`)
//line views/vaudit/RecordTable.html:32
		qw422016.E().S(`/admin/audit/` + model.AuditID.String())
//line views/vaudit/RecordTable.html:32
		qw422016.N().S(`">`)
//line views/vaudit/RecordTable.html:32
		components.StreamSVGRefIcon(qw422016, "cog", ps)
//line views/vaudit/RecordTable.html:32
		qw422016.N().S(`</a>
          </td>
          <td>`)
//line views/vaudit/RecordTable.html:34
		qw422016.E().S(model.T)
//line views/vaudit/RecordTable.html:34
		qw422016.N().S(`</td>
          <td>`)
//line views/vaudit/RecordTable.html:35
		qw422016.E().S(model.PK)
//line views/vaudit/RecordTable.html:35
		qw422016.N().S(`</td>
          <td>`)
//line views/vaudit/RecordTable.html:36
		view.StreamDiffs(qw422016, model.Changes)
//line views/vaudit/RecordTable.html:36
		qw422016.N().S(`</td>
          <td>`)
//line views/vaudit/RecordTable.html:37
		qw422016.N().D(len(model.Metadata))
//line views/vaudit/RecordTable.html:37
		qw422016.N().S(`</td>
          <td>`)
//line views/vaudit/RecordTable.html:38
		view.StreamTimestamp(qw422016, &model.Occurred)
//line views/vaudit/RecordTable.html:38
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vaudit/RecordTable.html:40
	}
//line views/vaudit/RecordTable.html:40
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vaudit/RecordTable.html:44
}

//line views/vaudit/RecordTable.html:44
func WriteRecordTable(qq422016 qtio422016.Writer, models audit.Records, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vaudit/RecordTable.html:44
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaudit/RecordTable.html:44
	StreamRecordTable(qw422016, models, params, as, ps)
//line views/vaudit/RecordTable.html:44
	qt422016.ReleaseWriter(qw422016)
//line views/vaudit/RecordTable.html:44
}

//line views/vaudit/RecordTable.html:44
func RecordTable(models audit.Records, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vaudit/RecordTable.html:44
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaudit/RecordTable.html:44
	WriteRecordTable(qb422016, models, params, as, ps)
//line views/vaudit/RecordTable.html:44
	qs422016 := string(qb422016.B)
//line views/vaudit/RecordTable.html:44
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaudit/RecordTable.html:44
	return qs422016
//line views/vaudit/RecordTable.html:44
}
