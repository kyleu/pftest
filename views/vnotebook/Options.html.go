// Code generated by qtc from "Options.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vnotebook/Options.html:2
package vnotebook

//line views/vnotebook/Options.html:2
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/layout"
)

//line views/vnotebook/Options.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vnotebook/Options.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vnotebook/Options.html:9
type Options struct {
	layout.Basic
}

//line views/vnotebook/Options.html:13
func (p *Options) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vnotebook/Options.html:13
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vnotebook/Options.html:15
	components.StreamSVGRefIcon(qw422016, `notebook`, ps)
//line views/vnotebook/Options.html:15
	qw422016.N().S(`Notebook</h3>
    <p>The notebook service isn't running. Here are your options:</p>
    <div>
      <a href="/notebook/action/start"><button>Start Server</button></a>
    </div>
  </div>
`)
//line views/vnotebook/Options.html:21
}

//line views/vnotebook/Options.html:21
func (p *Options) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vnotebook/Options.html:21
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vnotebook/Options.html:21
	p.StreamBody(qw422016, as, ps)
//line views/vnotebook/Options.html:21
	qt422016.ReleaseWriter(qw422016)
//line views/vnotebook/Options.html:21
}

//line views/vnotebook/Options.html:21
func (p *Options) Body(as *app.State, ps *cutil.PageState) string {
//line views/vnotebook/Options.html:21
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vnotebook/Options.html:21
	p.WriteBody(qb422016, as, ps)
//line views/vnotebook/Options.html:21
	qs422016 := string(qb422016.B)
//line views/vnotebook/Options.html:21
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vnotebook/Options.html:21
	return qs422016
//line views/vnotebook/Options.html:21
}