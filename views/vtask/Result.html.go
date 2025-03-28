// Code generated by qtc from "Result.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vtask/Result.html:1
package vtask

//line views/vtask/Result.html:1
import (
	"strings"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/task"
	"github.com/kyleu/pftest/app/util"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/components/view"
)

//line views/vtask/Result.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtask/Result.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtask/Result.html:12
func StreamResult(qw422016 *qt422016.Writer, as *app.State, res *task.Result, ps *cutil.PageState) {
//line views/vtask/Result.html:12
	qw422016.N().S(`
`)
//line views/vtask/Result.html:13
	if res.Error != "" {
//line views/vtask/Result.html:13
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vtask/Result.html:15
		components.StreamSVGIcon(qw422016, "error", ps)
//line views/vtask/Result.html:15
		qw422016.N().S(` Error</h3>
    <div class="mt">
      <pre class="error">Error: `)
//line views/vtask/Result.html:17
		qw422016.E().S(res.Error)
//line views/vtask/Result.html:17
		qw422016.N().S(`</pre>
    </div>
  </div>
`)
//line views/vtask/Result.html:20
	}
//line views/vtask/Result.html:20
	qw422016.N().S(`
`)
//line views/vtask/Result.html:22
	if len(res.Logs) > 0 {
//line views/vtask/Result.html:22
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vtask/Result.html:24
		components.StreamSVGIcon(qw422016, "file", ps)
//line views/vtask/Result.html:24
		qw422016.N().S(` `)
//line views/vtask/Result.html:24
		qw422016.E().S(res.Task.TitleSafe())
//line views/vtask/Result.html:24
		qw422016.N().S(` Logs</h3>
    <div class="mt">`)
//line views/vtask/Result.html:25
		components.StreamTerminal(qw422016, "console-list", strings.Join(res.Logs, "\n"))
//line views/vtask/Result.html:25
		qw422016.N().S(`</div>
  </div>
`)
//line views/vtask/Result.html:27
	}
//line views/vtask/Result.html:27
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      `)
//line views/vtask/Result.html:31
	qw422016.E().S(util.MicrosToMillis(res.Elapsed))
//line views/vtask/Result.html:31
	qw422016.N().S(`
      <a href="#modal-result-`)
//line views/vtask/Result.html:32
	qw422016.E().S(res.ID.String())
//line views/vtask/Result.html:32
	qw422016.N().S(`" title="JSON"><button>`)
//line views/vtask/Result.html:32
	components.StreamSVGButton(qw422016, `code`, ps)
//line views/vtask/Result.html:32
	qw422016.N().S(`</button></a>
    </div>
    <h3>`)
//line views/vtask/Result.html:34
	components.StreamSVGIcon(qw422016, "cog", ps)
//line views/vtask/Result.html:34
	qw422016.N().S(` `)
//line views/vtask/Result.html:34
	qw422016.E().S(res.Task.TitleSafe())
//line views/vtask/Result.html:34
	qw422016.N().S(` Result</h3>
`)
//line views/vtask/Result.html:35
	if len(res.Tags) > 0 {
//line views/vtask/Result.html:35
		qw422016.N().S(`    <div class="clear"></div>
    <div class="right mts">`)
//line views/vtask/Result.html:37
		view.StreamTags(qw422016, res.Tags, nil)
//line views/vtask/Result.html:37
		qw422016.N().S(`</div>
`)
//line views/vtask/Result.html:38
	}
//line views/vtask/Result.html:38
	qw422016.N().S(`    <div><em>`)
//line views/vtask/Result.html:39
	view.StreamTimestampRelative(qw422016, &res.Started, false)
//line views/vtask/Result.html:39
	qw422016.N().S(`</em></div>
    <div class="mt">
      `)
//line views/vtask/Result.html:41
	streamrenderResult(qw422016, res, as, ps)
//line views/vtask/Result.html:41
	qw422016.N().S(`
    </div>
  </div>
  `)
//line views/vtask/Result.html:44
	components.StreamJSONModal(qw422016, "result-"+res.ID.String(), "Result ["+res.String()+"] JSON", res, 1)
//line views/vtask/Result.html:44
	qw422016.N().S(`
`)
//line views/vtask/Result.html:45
}

//line views/vtask/Result.html:45
func WriteResult(qq422016 qtio422016.Writer, as *app.State, res *task.Result, ps *cutil.PageState) {
//line views/vtask/Result.html:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtask/Result.html:45
	StreamResult(qw422016, as, res, ps)
//line views/vtask/Result.html:45
	qt422016.ReleaseWriter(qw422016)
//line views/vtask/Result.html:45
}

//line views/vtask/Result.html:45
func Result(as *app.State, res *task.Result, ps *cutil.PageState) string {
//line views/vtask/Result.html:45
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtask/Result.html:45
	WriteResult(qb422016, as, res, ps)
//line views/vtask/Result.html:45
	qs422016 := string(qb422016.B)
//line views/vtask/Result.html:45
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtask/Result.html:45
	return qs422016
//line views/vtask/Result.html:45
}

//line views/vtask/Result.html:47
func StreamResultSummary(qw422016 *qt422016.Writer, as *app.State, res *task.Result, ps *cutil.PageState) {
//line views/vtask/Result.html:47
	qw422016.N().S(`
`)
//line views/vtask/Result.html:48
	if res.Error != "" {
//line views/vtask/Result.html:48
		qw422016.N().S(`  <div class="card">
    <h3>`)
//line views/vtask/Result.html:50
		components.StreamSVGIcon(qw422016, "error", ps)
//line views/vtask/Result.html:50
		qw422016.N().S(` Error</h3>
    <div class="mt">
      <pre class="error">Error: `)
//line views/vtask/Result.html:52
		qw422016.E().S(res.Error)
//line views/vtask/Result.html:52
		qw422016.N().S(`</pre>
    </div>
  </div>
`)
//line views/vtask/Result.html:55
	}
//line views/vtask/Result.html:55
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      `)
//line views/vtask/Result.html:59
	qw422016.E().S(util.MicrosToMillis(res.Elapsed))
//line views/vtask/Result.html:59
	qw422016.N().S(`
      <a href="#modal-result" title="JSON"><button>`)
//line views/vtask/Result.html:60
	components.StreamSVGButton(qw422016, `code`, ps)
//line views/vtask/Result.html:60
	qw422016.N().S(`</button></a>
    </div>
    <h3>`)
//line views/vtask/Result.html:62
	components.StreamSVGIcon(qw422016, "cog", ps)
//line views/vtask/Result.html:62
	qw422016.N().S(` `)
//line views/vtask/Result.html:62
	qw422016.E().S(res.Task.TitleSafe())
//line views/vtask/Result.html:62
	qw422016.N().S(` Result</h3>
`)
//line views/vtask/Result.html:63
	if len(res.Tags) > 0 {
//line views/vtask/Result.html:63
		qw422016.N().S(`    <div class="clear"></div>
    <div class="right mts">`)
//line views/vtask/Result.html:65
		view.StreamTags(qw422016, res.Tags, nil)
//line views/vtask/Result.html:65
		qw422016.N().S(`</div>
`)
//line views/vtask/Result.html:66
	}
//line views/vtask/Result.html:66
	qw422016.N().S(`    <div><em>`)
//line views/vtask/Result.html:67
	view.StreamTimestampRelative(qw422016, &res.Started, false)
//line views/vtask/Result.html:67
	qw422016.N().S(`</em></div>
    <div class="mt">
      `)
//line views/vtask/Result.html:69
	streamrenderResult(qw422016, res, as, ps)
//line views/vtask/Result.html:69
	qw422016.N().S(`
    </div>
  </div>
  `)
//line views/vtask/Result.html:72
	components.StreamJSONModal(qw422016, "result", "Result ["+res.String()+"] JSON", res, 1)
//line views/vtask/Result.html:72
	qw422016.N().S(`
`)
//line views/vtask/Result.html:73
}

//line views/vtask/Result.html:73
func WriteResultSummary(qq422016 qtio422016.Writer, as *app.State, res *task.Result, ps *cutil.PageState) {
//line views/vtask/Result.html:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtask/Result.html:73
	StreamResultSummary(qw422016, as, res, ps)
//line views/vtask/Result.html:73
	qt422016.ReleaseWriter(qw422016)
//line views/vtask/Result.html:73
}

//line views/vtask/Result.html:73
func ResultSummary(as *app.State, res *task.Result, ps *cutil.PageState) string {
//line views/vtask/Result.html:73
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtask/Result.html:73
	WriteResultSummary(qb422016, as, res, ps)
//line views/vtask/Result.html:73
	qs422016 := string(qb422016.B)
//line views/vtask/Result.html:73
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtask/Result.html:73
	return qs422016
//line views/vtask/Result.html:73
}

//line views/vtask/Result.html:75
func streamrenderResult(qw422016 *qt422016.Writer, res *task.Result, as *app.State, ps *cutil.PageState) {
//line views/vtask/Result.html:76
	if res.Data == nil {
//line views/vtask/Result.html:76
		qw422016.N().S(`<em>no data</em>`)
//line views/vtask/Result.html:78
	} else {
//line views/vtask/Result.html:79
		components.StreamJSON(qw422016, res.Data)
//line views/vtask/Result.html:80
	}
//line views/vtask/Result.html:81
}

//line views/vtask/Result.html:81
func writerenderResult(qq422016 qtio422016.Writer, res *task.Result, as *app.State, ps *cutil.PageState) {
//line views/vtask/Result.html:81
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtask/Result.html:81
	streamrenderResult(qw422016, res, as, ps)
//line views/vtask/Result.html:81
	qt422016.ReleaseWriter(qw422016)
//line views/vtask/Result.html:81
}

//line views/vtask/Result.html:81
func renderResult(res *task.Result, as *app.State, ps *cutil.PageState) string {
//line views/vtask/Result.html:81
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtask/Result.html:81
	writerenderResult(qb422016, res, as, ps)
//line views/vtask/Result.html:81
	qs422016 := string(qb422016.B)
//line views/vtask/Result.html:81
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtask/Result.html:81
	return qs422016
//line views/vtask/Result.html:81
}
