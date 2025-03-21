// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vaudit/Edit.html:1
package vaudit

//line views/vaudit/Edit.html:1
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/audit"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/components/edit"
	"github.com/kyleu/pftest/views/layout"
)

//line views/vaudit/Edit.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vaudit/Edit.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vaudit/Edit.html:11
type Edit struct {
	layout.Basic
	Model *audit.Audit
	IsNew bool
}

//line views/vaudit/Edit.html:17
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaudit/Edit.html:17
	qw422016.N().S(`
  <div class="card">
`)
//line views/vaudit/Edit.html:19
	if p.IsNew {
//line views/vaudit/Edit.html:19
		qw422016.N().S(`    <div class="right"><a href="/admin/audit/random"><button>Random</button></a></div>
    <h3>`)
//line views/vaudit/Edit.html:21
		components.StreamSVGIcon(qw422016, `cog`, ps)
//line views/vaudit/Edit.html:21
		qw422016.N().S(` New Audit</h3>
    <form action="/admin/audit/new" class="mt" method="post">
`)
//line views/vaudit/Edit.html:23
	} else {
//line views/vaudit/Edit.html:23
		qw422016.N().S(`    <div class="right"><a href="`)
//line views/vaudit/Edit.html:24
		qw422016.E().S(p.Model.WebPath())
//line views/vaudit/Edit.html:24
		qw422016.N().S(`/delete" onclick="return confirm('Are you sure you wish to delete audit [`)
//line views/vaudit/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/vaudit/Edit.html:24
		qw422016.N().S(`]?')"><button>`)
//line views/vaudit/Edit.html:24
		components.StreamSVGButton(qw422016, "times", ps)
//line views/vaudit/Edit.html:24
		qw422016.N().S(` Delete</button></a></div>
    <h3>`)
//line views/vaudit/Edit.html:25
		components.StreamSVGIcon(qw422016, `cog`, ps)
//line views/vaudit/Edit.html:25
		qw422016.N().S(` Edit Audit [`)
//line views/vaudit/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/vaudit/Edit.html:25
		qw422016.N().S(`]</h3>
    <form action="" class="mt" method="post">
`)
//line views/vaudit/Edit.html:27
	}
//line views/vaudit/Edit.html:27
	qw422016.N().S(`      <div class="overflow full-width">
        <table class="mt expanded">
          <tbody>
            `)
//line views/vaudit/Edit.html:31
	if p.IsNew {
//line views/vaudit/Edit.html:31
		edit.StreamStringTable(qw422016, "id", "", "ID", p.Model.ID.String(), 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/vaudit/Edit.html:31
	}
//line views/vaudit/Edit.html:31
	qw422016.N().S(`
            `)
//line views/vaudit/Edit.html:32
	edit.StreamStringTable(qw422016, "app", "", "App", p.Model.App, 5, "String text")
//line views/vaudit/Edit.html:32
	qw422016.N().S(`
            `)
//line views/vaudit/Edit.html:33
	edit.StreamStringTable(qw422016, "act", "", "Act", p.Model.Act, 5, "String text")
//line views/vaudit/Edit.html:33
	qw422016.N().S(`
            `)
//line views/vaudit/Edit.html:34
	edit.StreamStringTable(qw422016, "client", "", "Client", p.Model.Client, 5, "String text")
//line views/vaudit/Edit.html:34
	qw422016.N().S(`
            `)
//line views/vaudit/Edit.html:35
	edit.StreamStringTable(qw422016, "server", "", "Server", p.Model.Server, 5, "String text")
//line views/vaudit/Edit.html:35
	qw422016.N().S(`
            `)
//line views/vaudit/Edit.html:36
	edit.StreamStringTable(qw422016, "user", "", "User", p.Model.User, 5, "String text")
//line views/vaudit/Edit.html:36
	qw422016.N().S(`
            `)
//line views/vaudit/Edit.html:37
	edit.StreamTextareaTable(qw422016, "metadata", "", "Metadata", 8, util.ToJSON(p.Model.Metadata), 5, "JSON object")
//line views/vaudit/Edit.html:37
	qw422016.N().S(`
            `)
//line views/vaudit/Edit.html:38
	edit.StreamStringTable(qw422016, "message", "", "Message", p.Model.Message, 5, "String text")
//line views/vaudit/Edit.html:38
	qw422016.N().S(`
            `)
//line views/vaudit/Edit.html:39
	edit.StreamTimestampTable(qw422016, "started", "", "Started", &p.Model.Started, 5, "Date and time, in almost any format")
//line views/vaudit/Edit.html:39
	qw422016.N().S(`
            `)
//line views/vaudit/Edit.html:40
	edit.StreamTimestampTable(qw422016, "completed", "", "Completed", &p.Model.Completed, 5, "Date and time, in almost any format")
//line views/vaudit/Edit.html:40
	qw422016.N().S(`
            <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
          </tbody>
        </table>
      </div>
    </form>
  </div>
`)
//line views/vaudit/Edit.html:47
}

//line views/vaudit/Edit.html:47
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vaudit/Edit.html:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vaudit/Edit.html:47
	p.StreamBody(qw422016, as, ps)
//line views/vaudit/Edit.html:47
	qt422016.ReleaseWriter(qw422016)
//line views/vaudit/Edit.html:47
}

//line views/vaudit/Edit.html:47
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vaudit/Edit.html:47
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vaudit/Edit.html:47
	p.WriteBody(qb422016, as, ps)
//line views/vaudit/Edit.html:47
	qs422016 := string(qb422016.B)
//line views/vaudit/Edit.html:47
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vaudit/Edit.html:47
	return qs422016
//line views/vaudit/Edit.html:47
}
