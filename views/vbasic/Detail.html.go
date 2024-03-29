// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vbasic/Detail.html:2
package vbasic

//line views/vbasic/Detail.html:2
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/basic"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/relation"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/components/view"
	"github.com/kyleu/pftest/views/layout"
	"github.com/kyleu/pftest/views/vrelation"
)

//line views/vbasic/Detail.html:15
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vbasic/Detail.html:15
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vbasic/Detail.html:15
type Detail struct {
	layout.Basic
	Model                 *basic.Basic
	Params                filter.ParamSet
	RelRelationsByBasicID relation.Relations
}

//line views/vbasic/Detail.html:22
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vbasic/Detail.html:22
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-basic"><button type="button">JSON</button></a>
      <a href="`)
//line views/vbasic/Detail.html:26
	qw422016.E().S(p.Model.WebPath())
//line views/vbasic/Detail.html:26
	qw422016.N().S(`/edit"><button>`)
//line views/vbasic/Detail.html:26
	components.StreamSVGRef(qw422016, "edit", 15, 15, "icon", ps)
//line views/vbasic/Detail.html:26
	qw422016.N().S(`Edit</button></a>
    </div>
    <h3>`)
//line views/vbasic/Detail.html:28
	components.StreamSVGRefIcon(qw422016, `star`, ps)
//line views/vbasic/Detail.html:28
	qw422016.N().S(` `)
//line views/vbasic/Detail.html:28
	qw422016.E().S(p.Model.TitleString())
//line views/vbasic/Detail.html:28
	qw422016.N().S(`</h3>
    <div><a href="/basic"><em>Basic</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
            <td>`)
//line views/vbasic/Detail.html:35
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vbasic/Detail.html:35
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Name</th>
            <td><strong>`)
//line views/vbasic/Detail.html:39
	view.StreamString(qw422016, p.Model.Name)
//line views/vbasic/Detail.html:39
	qw422016.N().S(`</strong></td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Status</th>
            <td><strong>`)
//line views/vbasic/Detail.html:43
	qw422016.E().S(p.Model.Status)
//line views/vbasic/Detail.html:43
	qw422016.N().S(`</strong></td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vbasic/Detail.html:47
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vbasic/Detail.html:47
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vbasic/Detail.html:54
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-RelationsByBasicID" type="checkbox" hidden="hidden"`)
//line views/vbasic/Detail.html:59
	if p.Params.Specifies(`relation`) {
//line views/vbasic/Detail.html:59
		qw422016.N().S(` checked="checked"`)
//line views/vbasic/Detail.html:59
	}
//line views/vbasic/Detail.html:59
	qw422016.N().S(` />
        <label for="accordion-RelationsByBasicID">
          `)
//line views/vbasic/Detail.html:61
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vbasic/Detail.html:61
	qw422016.N().S(`
          `)
//line views/vbasic/Detail.html:62
	components.StreamSVGRef(qw422016, `star`, 16, 16, `icon`, ps)
//line views/vbasic/Detail.html:62
	qw422016.N().S(`
          `)
//line views/vbasic/Detail.html:63
	qw422016.E().S(util.StringPlural(len(p.RelRelationsByBasicID), "Relation"))
//line views/vbasic/Detail.html:63
	qw422016.N().S(` by [Basic ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vbasic/Detail.html:66
	if len(p.RelRelationsByBasicID) == 0 {
//line views/vbasic/Detail.html:66
		qw422016.N().S(`          <em>no related Relations</em>
`)
//line views/vbasic/Detail.html:68
	} else {
//line views/vbasic/Detail.html:68
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vbasic/Detail.html:70
		vrelation.StreamTable(qw422016, p.RelRelationsByBasicID, nil, p.Params, as, ps)
//line views/vbasic/Detail.html:70
		qw422016.N().S(`
          </div>
`)
//line views/vbasic/Detail.html:72
	}
//line views/vbasic/Detail.html:72
	qw422016.N().S(`        </div></div></div>
      </li>
    </ul>
  </div>
  `)
//line views/vbasic/Detail.html:77
	components.StreamJSONModal(qw422016, "basic", "Basic JSON", p.Model, 1)
//line views/vbasic/Detail.html:77
	qw422016.N().S(`
`)
//line views/vbasic/Detail.html:78
}

//line views/vbasic/Detail.html:78
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vbasic/Detail.html:78
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vbasic/Detail.html:78
	p.StreamBody(qw422016, as, ps)
//line views/vbasic/Detail.html:78
	qt422016.ReleaseWriter(qw422016)
//line views/vbasic/Detail.html:78
}

//line views/vbasic/Detail.html:78
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vbasic/Detail.html:78
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vbasic/Detail.html:78
	p.WriteBody(qb422016, as, ps)
//line views/vbasic/Detail.html:78
	qs422016 := string(qb422016.B)
//line views/vbasic/Detail.html:78
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vbasic/Detail.html:78
	return qs422016
//line views/vbasic/Detail.html:78
}
