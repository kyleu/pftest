// Code generated by qtc from "Results.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vdatabase/Results.html:2
package vdatabase

//line views/vdatabase/Results.html:2
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/layout"
)

//line views/vdatabase/Results.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vdatabase/Results.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vdatabase/Results.html:11
type Results struct {
	layout.Basic
	Svc     *database.Service
	Schema  string
	Table   string
	Results []util.ValueMap
	Timing  int
	Error   error
}

//line views/vdatabase/Results.html:21
func (p *Results) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Results.html:21
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vdatabase/Results.html:23
	qw422016.E().S(util.MicrosToMillis(p.Timing))
//line views/vdatabase/Results.html:23
	qw422016.N().S(` elapsed</div>
    <h3>`)
//line views/vdatabase/Results.html:24
	components.StreamSVGRefIcon(qw422016, `database`, ps)
//line views/vdatabase/Results.html:24
	qw422016.N().S(`Table [`)
//line views/vdatabase/Results.html:24
	if p.Schema != "default" {
//line views/vdatabase/Results.html:24
		qw422016.E().S(p.Schema)
//line views/vdatabase/Results.html:24
		qw422016.N().S(`:`)
//line views/vdatabase/Results.html:24
	}
//line views/vdatabase/Results.html:24
	qw422016.E().S(p.Table)
//line views/vdatabase/Results.html:24
	qw422016.N().S(`]</h3>
    <div class="right">`)
//line views/vdatabase/Results.html:25
	qw422016.N().D(len(p.Results))
//line views/vdatabase/Results.html:25
	qw422016.N().S(` rows returned</div>
`)
//line views/vdatabase/Results.html:26
	if p.Error != nil {
//line views/vdatabase/Results.html:26
		qw422016.N().S(`    <div class="mt error">`)
//line views/vdatabase/Results.html:27
		qw422016.E().S(p.Error.Error())
//line views/vdatabase/Results.html:27
		qw422016.N().S(`</div>
`)
//line views/vdatabase/Results.html:28
	}
//line views/vdatabase/Results.html:28
	qw422016.N().S(`    <div class="mt">`)
//line views/vdatabase/Results.html:29
	components.StreamDisplayMaps(qw422016, p.Results, true)
//line views/vdatabase/Results.html:29
	qw422016.N().S(`</div>
  </div>
`)
//line views/vdatabase/Results.html:31
}

//line views/vdatabase/Results.html:31
func (p *Results) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Results.html:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vdatabase/Results.html:31
	p.StreamBody(qw422016, as, ps)
//line views/vdatabase/Results.html:31
	qt422016.ReleaseWriter(qw422016)
//line views/vdatabase/Results.html:31
}

//line views/vdatabase/Results.html:31
func (p *Results) Body(as *app.State, ps *cutil.PageState) string {
//line views/vdatabase/Results.html:31
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vdatabase/Results.html:31
	p.WriteBody(qb422016, as, ps)
//line views/vdatabase/Results.html:31
	qs422016 := string(qb422016.B)
//line views/vdatabase/Results.html:31
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vdatabase/Results.html:31
	return qs422016
//line views/vdatabase/Results.html:31
}