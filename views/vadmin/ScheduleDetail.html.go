// Code generated by qtc from "ScheduleDetail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/ScheduleDetail.html:2
package vadmin

//line views/vadmin/ScheduleDetail.html:2
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/schedule"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/layout"
)

//line views/vadmin/ScheduleDetail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/ScheduleDetail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/ScheduleDetail.html:11
type ScheduleDetail struct {
	layout.Basic
	Job    *schedule.Job
	Result *schedule.Result
}

//line views/vadmin/ScheduleDetail.html:17
func (p *ScheduleDetail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/ScheduleDetail.html:17
	qw422016.N().S(`
  <div class="card">
    <h3>Scheduled Job [`)
//line views/vadmin/ScheduleDetail.html:19
	qw422016.E().S(p.Job.String())
//line views/vadmin/ScheduleDetail.html:19
	qw422016.N().S(`]</h3>
    `)
//line views/vadmin/ScheduleDetail.html:20
	streamjobTable(qw422016, schedule.Jobs{p.Job})
//line views/vadmin/ScheduleDetail.html:20
	qw422016.N().S(`
  </div>
`)
//line views/vadmin/ScheduleDetail.html:22
	if p.Result != nil {
//line views/vadmin/ScheduleDetail.html:22
		qw422016.N().S(`  <div class="card">
    <div class="right">`)
//line views/vadmin/ScheduleDetail.html:24
		qw422016.E().S(util.MicrosToMillis(p.Result.DurationMicro))
//line views/vadmin/ScheduleDetail.html:24
		qw422016.N().S(`</div>
    <h3>Most Recent Result</h3>
    <em>`)
//line views/vadmin/ScheduleDetail.html:26
		qw422016.E().S(util.TimeToFull(&p.Result.Occurred))
//line views/vadmin/ScheduleDetail.html:26
		qw422016.N().S(`</em>
    `)
//line views/vadmin/ScheduleDetail.html:27
		components.StreamJSON(qw422016, p.Result.Returned)
//line views/vadmin/ScheduleDetail.html:27
		qw422016.N().S(`
  </div>
`)
//line views/vadmin/ScheduleDetail.html:29
	}
//line views/vadmin/ScheduleDetail.html:30
}

//line views/vadmin/ScheduleDetail.html:30
func (p *ScheduleDetail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/ScheduleDetail.html:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/ScheduleDetail.html:30
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/ScheduleDetail.html:30
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/ScheduleDetail.html:30
}

//line views/vadmin/ScheduleDetail.html:30
func (p *ScheduleDetail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/ScheduleDetail.html:30
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/ScheduleDetail.html:30
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/ScheduleDetail.html:30
	qs422016 := string(qb422016.B)
//line views/vadmin/ScheduleDetail.html:30
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/ScheduleDetail.html:30
	return qs422016
//line views/vadmin/ScheduleDetail.html:30
}