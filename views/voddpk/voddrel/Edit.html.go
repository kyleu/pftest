// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/voddpk/voddrel/Edit.html:1
package voddrel

//line views/voddpk/voddrel/Edit.html:1
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/oddpk/oddrel"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/components/edit"
	"github.com/kyleu/pftest/views/layout"
)

//line views/voddpk/voddrel/Edit.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/voddpk/voddrel/Edit.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/voddpk/voddrel/Edit.html:11
type Edit struct {
	layout.Basic
	Model *oddrel.Oddrel
	Paths []string
	IsNew bool
}

//line views/voddpk/voddrel/Edit.html:18
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/voddpk/voddrel/Edit.html:18
	qw422016.N().S(`
  <div class="card">
`)
//line views/voddpk/voddrel/Edit.html:20
	if p.IsNew {
//line views/voddpk/voddrel/Edit.html:20
		qw422016.N().S(`    <div class="right"><a href="?prototype=random"><button>Random</button></a></div>
    <h3>`)
//line views/voddpk/voddrel/Edit.html:22
		components.StreamSVGIcon(qw422016, `star`, ps)
//line views/voddpk/voddrel/Edit.html:22
		qw422016.N().S(` New Oddrel</h3>
`)
//line views/voddpk/voddrel/Edit.html:23
	} else {
//line views/voddpk/voddrel/Edit.html:23
		qw422016.N().S(`    <div class="right"><a class="link-confirm" href="`)
//line views/voddpk/voddrel/Edit.html:24
		qw422016.E().S(p.Model.WebPath(p.Paths...))
//line views/voddpk/voddrel/Edit.html:24
		qw422016.N().S(`/delete" data-message="Are you sure you wish to delete oddrel [`)
//line views/voddpk/voddrel/Edit.html:24
		qw422016.E().S(p.Model.String())
//line views/voddpk/voddrel/Edit.html:24
		qw422016.N().S(`]?"><button>`)
//line views/voddpk/voddrel/Edit.html:24
		components.StreamSVGButton(qw422016, "times", ps)
//line views/voddpk/voddrel/Edit.html:24
		qw422016.N().S(` Delete</button></a></div>
    <h3>`)
//line views/voddpk/voddrel/Edit.html:25
		components.StreamSVGIcon(qw422016, `star`, ps)
//line views/voddpk/voddrel/Edit.html:25
		qw422016.N().S(` Edit Oddrel [`)
//line views/voddpk/voddrel/Edit.html:25
		qw422016.E().S(p.Model.String())
//line views/voddpk/voddrel/Edit.html:25
		qw422016.N().S(`]</h3>
`)
//line views/voddpk/voddrel/Edit.html:26
	}
//line views/voddpk/voddrel/Edit.html:26
	qw422016.N().S(`    <form action="`)
//line views/voddpk/voddrel/Edit.html:27
	qw422016.E().S(util.Choose(p.IsNew, oddrel.Route(p.Paths...)+`/_new`, p.Model.WebPath(p.Paths...)+`/edit`))
//line views/voddpk/voddrel/Edit.html:27
	qw422016.N().S(`" class="mt" method="post">
      <table class="mt expanded">
        <tbody>
          `)
//line views/voddpk/voddrel/Edit.html:30
	if p.IsNew {
//line views/voddpk/voddrel/Edit.html:30
		edit.StreamUUIDTable(qw422016, "id", "", "ID", &p.Model.ID, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/voddpk/voddrel/Edit.html:30
	}
//line views/voddpk/voddrel/Edit.html:30
	qw422016.N().S(`
          `)
//line views/voddpk/voddrel/Edit.html:31
	edit.StreamUUIDTable(qw422016, "project", "", "Project", &p.Model.Project, 5, "UUID in format (00000000-0000-0000-0000-000000000000)")
//line views/voddpk/voddrel/Edit.html:31
	qw422016.N().S(`
          `)
//line views/voddpk/voddrel/Edit.html:32
	edit.StreamStringTable(qw422016, "path", "", "Path", p.Model.Path, 5, "String text")
//line views/voddpk/voddrel/Edit.html:32
	qw422016.N().S(`
          <tr><td colspan="2"><button type="submit">Save Changes</button></td></tr>
        </tbody>
      </table>
    </form>
  </div>
`)
//line views/voddpk/voddrel/Edit.html:38
}

//line views/voddpk/voddrel/Edit.html:38
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/voddpk/voddrel/Edit.html:38
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/voddpk/voddrel/Edit.html:38
	p.StreamBody(qw422016, as, ps)
//line views/voddpk/voddrel/Edit.html:38
	qt422016.ReleaseWriter(qw422016)
//line views/voddpk/voddrel/Edit.html:38
}

//line views/voddpk/voddrel/Edit.html:38
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/voddpk/voddrel/Edit.html:38
	qb422016 := qt422016.AcquireByteBuffer()
//line views/voddpk/voddrel/Edit.html:38
	p.WriteBody(qb422016, as, ps)
//line views/voddpk/voddrel/Edit.html:38
	qs422016 := string(qb422016.B)
//line views/voddpk/voddrel/Edit.html:38
	qt422016.ReleaseByteBuffer(qb422016)
//line views/voddpk/voddrel/Edit.html:38
	return qs422016
//line views/voddpk/voddrel/Edit.html:38
}
