// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsoftdel/Detail.html:2
package vsoftdel

//line views/vsoftdel/Detail.html:2
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/softdel"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/layout"
)

//line views/vsoftdel/Detail.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsoftdel/Detail.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsoftdel/Detail.html:10
type Detail struct {
	layout.Basic
	Model *softdel.Softdel
}

//line views/vsoftdel/Detail.html:15
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsoftdel/Detail.html:15
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-softdel"><button type="button">JSON</button></a>
      <a href="`)
//line views/vsoftdel/Detail.html:19
	qw422016.E().S(p.Model.WebPath())
//line views/vsoftdel/Detail.html:19
	qw422016.N().S(`/edit"><button>Edit</button></a>
    </div>
    <h3>`)
//line views/vsoftdel/Detail.html:21
	components.StreamSVGRefIcon(qw422016, `star`, ps)
//line views/vsoftdel/Detail.html:21
	qw422016.N().S(` `)
//line views/vsoftdel/Detail.html:21
	qw422016.E().S(p.Model.TitleString())
//line views/vsoftdel/Detail.html:21
	qw422016.N().S(`</h3>
    <div><a href="/softdel"><em>Softdel</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="String text">ID</th>
          <td>`)
//line views/vsoftdel/Detail.html:27
	qw422016.E().S(p.Model.ID)
//line views/vsoftdel/Detail.html:27
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format">Created</th>
          <td>`)
//line views/vsoftdel/Detail.html:31
	components.StreamDisplayTimestamp(qw422016, &p.Model.Created)
//line views/vsoftdel/Detail.html:31
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Updated</th>
          <td>`)
//line views/vsoftdel/Detail.html:35
	components.StreamDisplayTimestamp(qw422016, p.Model.Updated)
//line views/vsoftdel/Detail.html:35
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Deleted</th>
          <td>`)
//line views/vsoftdel/Detail.html:39
	components.StreamDisplayTimestamp(qw422016, p.Model.Deleted)
//line views/vsoftdel/Detail.html:39
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vsoftdel/Detail.html:45
	qw422016.N().S(`  `)
//line views/vsoftdel/Detail.html:46
	components.StreamJSONModal(qw422016, "softdel", "Softdel JSON", p.Model, 1)
//line views/vsoftdel/Detail.html:46
	qw422016.N().S(`
`)
//line views/vsoftdel/Detail.html:47
}

//line views/vsoftdel/Detail.html:47
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsoftdel/Detail.html:47
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsoftdel/Detail.html:47
	p.StreamBody(qw422016, as, ps)
//line views/vsoftdel/Detail.html:47
	qt422016.ReleaseWriter(qw422016)
//line views/vsoftdel/Detail.html:47
}

//line views/vsoftdel/Detail.html:47
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsoftdel/Detail.html:47
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsoftdel/Detail.html:47
	p.WriteBody(qb422016, as, ps)
//line views/vsoftdel/Detail.html:47
	qs422016 := string(qb422016.B)
//line views/vsoftdel/Detail.html:47
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsoftdel/Detail.html:47
	return qs422016
//line views/vsoftdel/Detail.html:47
}