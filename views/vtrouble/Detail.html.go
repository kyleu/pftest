// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vtrouble/Detail.html:2
package vtrouble

//line views/vtrouble/Detail.html:2
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/trouble"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/layout"
)

//line views/vtrouble/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtrouble/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtrouble/Detail.html:11
type Detail struct {
	layout.Basic
	Model      *trouble.Trouble
	Params     filter.ParamSet
	Selectcols trouble.Troubles
}

//line views/vtrouble/Detail.html:18
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtrouble/Detail.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-trouble"><button type="button">JSON</button></a>
      <a href="`)
//line views/vtrouble/Detail.html:22
	qw422016.E().S(p.Model.WebPath())
//line views/vtrouble/Detail.html:22
	qw422016.N().S(`/edit"><button>Edit</button></a>
    </div>
    <h3>`)
//line views/vtrouble/Detail.html:24
	components.StreamSVGRefIcon(qw422016, `star`, ps)
//line views/vtrouble/Detail.html:24
	qw422016.N().S(` `)
//line views/vtrouble/Detail.html:24
	qw422016.E().S(p.Model.TitleString())
//line views/vtrouble/Detail.html:24
	qw422016.N().S(`</h3>
    <div><a href="/troub/le"><em>Trouble</em></a></div>
    <table class="mt">
      <tbody>
        <tr>
          <th class="shrink" title="String text">From</th>
          <td>`)
//line views/vtrouble/Detail.html:30
	qw422016.E().S(p.Model.From)
//line views/vtrouble/Detail.html:30
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Integer">Where</th>
          <td>`)
//line views/vtrouble/Detail.html:34
	qw422016.N().D(p.Model.Where)
//line views/vtrouble/Detail.html:34
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Integer">Selectcol</th>
          <td>`)
//line views/vtrouble/Detail.html:38
	qw422016.N().D(p.Model.Selectcol)
//line views/vtrouble/Detail.html:38
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Limit</th>
          <td>`)
//line views/vtrouble/Detail.html:42
	qw422016.E().S(p.Model.Limit)
//line views/vtrouble/Detail.html:42
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="String text">Group</th>
          <td>`)
//line views/vtrouble/Detail.html:46
	qw422016.E().S(p.Model.Group)
//line views/vtrouble/Detail.html:46
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink" title="Date and time, in almost any format (optional)">Delete</th>
          <td>`)
//line views/vtrouble/Detail.html:50
	components.StreamDisplayTimestamp(qw422016, p.Model.Delete)
//line views/vtrouble/Detail.html:50
	qw422016.N().S(`</td>
        </tr>
      </tbody>
    </table>
  </div>
`)
//line views/vtrouble/Detail.html:55
	if len(p.Selectcols) > 1 {
//line views/vtrouble/Detail.html:55
		qw422016.N().S(`  <div class="card">
    <h3>Selectcols</h3>
`)
//line views/vtrouble/Detail.html:58
		prms := p.Params.Get("trouble", nil, ps.Logger).Sanitize("trouble")

//line views/vtrouble/Detail.html:58
		qw422016.N().S(`    <table class="mt">
      <thead>
        <tr>
          `)
//line views/vtrouble/Detail.html:62
		components.StreamTableHeaderSimple(qw422016, "trouble", "from", "From", "String text", prms, ps.URI, ps)
//line views/vtrouble/Detail.html:62
		qw422016.N().S(`
          `)
//line views/vtrouble/Detail.html:63
		components.StreamTableHeaderSimple(qw422016, "trouble", "where", "Where", "Integer", prms, ps.URI, ps)
//line views/vtrouble/Detail.html:63
		qw422016.N().S(`
          `)
//line views/vtrouble/Detail.html:64
		components.StreamTableHeaderSimple(qw422016, "trouble", "selectcol", "Selectcol", "Integer", prms, ps.URI, ps)
//line views/vtrouble/Detail.html:64
		qw422016.N().S(`
        </tr>
      </thead>
      <tbody>
`)
//line views/vtrouble/Detail.html:68
		for _, model := range p.Selectcols {
//line views/vtrouble/Detail.html:68
			qw422016.N().S(`        <tr>
          <td><a href="/troub/le/`)
//line views/vtrouble/Detail.html:70
			qw422016.E().S(model.From)
//line views/vtrouble/Detail.html:70
			qw422016.N().S(`/`)
//line views/vtrouble/Detail.html:70
			qw422016.N().D(model.Where)
//line views/vtrouble/Detail.html:70
			qw422016.N().S(`/selectcol/`)
//line views/vtrouble/Detail.html:70
			qw422016.N().D(model.Selectcol)
//line views/vtrouble/Detail.html:70
			qw422016.N().S(`">`)
//line views/vtrouble/Detail.html:70
			qw422016.E().S(model.From)
//line views/vtrouble/Detail.html:70
			qw422016.N().S(`</a></td>
          <td><a href="/troub/le/`)
//line views/vtrouble/Detail.html:71
			qw422016.E().S(model.From)
//line views/vtrouble/Detail.html:71
			qw422016.N().S(`/`)
//line views/vtrouble/Detail.html:71
			qw422016.N().D(model.Where)
//line views/vtrouble/Detail.html:71
			qw422016.N().S(`/selectcol/`)
//line views/vtrouble/Detail.html:71
			qw422016.N().D(model.Selectcol)
//line views/vtrouble/Detail.html:71
			qw422016.N().S(`">`)
//line views/vtrouble/Detail.html:71
			qw422016.N().D(model.Where)
//line views/vtrouble/Detail.html:71
			qw422016.N().S(`</a></td>
          <td><a href="/troub/le/`)
//line views/vtrouble/Detail.html:72
			qw422016.E().S(model.From)
//line views/vtrouble/Detail.html:72
			qw422016.N().S(`/`)
//line views/vtrouble/Detail.html:72
			qw422016.N().D(model.Where)
//line views/vtrouble/Detail.html:72
			qw422016.N().S(`/selectcol/`)
//line views/vtrouble/Detail.html:72
			qw422016.N().D(model.Selectcol)
//line views/vtrouble/Detail.html:72
			qw422016.N().S(`">`)
//line views/vtrouble/Detail.html:72
			qw422016.N().D(model.Selectcol)
//line views/vtrouble/Detail.html:72
			qw422016.N().S(`</a></td>
        </tr>
`)
//line views/vtrouble/Detail.html:74
		}
//line views/vtrouble/Detail.html:74
		qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vtrouble/Detail.html:78
	}
//line views/vtrouble/Detail.html:80
	qw422016.N().S(`  `)
//line views/vtrouble/Detail.html:81
	components.StreamJSONModal(qw422016, "trouble", "Trouble JSON", p.Model, 1)
//line views/vtrouble/Detail.html:81
	qw422016.N().S(`
`)
//line views/vtrouble/Detail.html:82
}

//line views/vtrouble/Detail.html:82
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtrouble/Detail.html:82
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtrouble/Detail.html:82
	p.StreamBody(qw422016, as, ps)
//line views/vtrouble/Detail.html:82
	qt422016.ReleaseWriter(qw422016)
//line views/vtrouble/Detail.html:82
}

//line views/vtrouble/Detail.html:82
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vtrouble/Detail.html:82
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtrouble/Detail.html:82
	p.WriteBody(qb422016, as, ps)
//line views/vtrouble/Detail.html:82
	qs422016 := string(qb422016.B)
//line views/vtrouble/Detail.html:82
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtrouble/Detail.html:82
	return qs422016
//line views/vtrouble/Detail.html:82
}