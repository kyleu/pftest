// Code generated by qtc from "Request.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vhar/Request.html:1
package vhar

//line views/vhar/Request.html:1
import (
	"fmt"

	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/har"
	"github.com/kyleu/pftest/app/util"
)

//line views/vhar/Request.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vhar/Request.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vhar/Request.html:9
func StreamRenderRequest(qw422016 *qt422016.Writer, key string, r *har.Request, ps *cutil.PageState) {
//line views/vhar/Request.html:9
	qw422016.N().S(`
  <div class="overflow full-width">
    <table class="min-200 expanded">
      <tbody>
        <tr>
          <th class="shrink">Method</th>
          <td>`)
//line views/vhar/Request.html:15
	qw422016.E().S(r.Method)
//line views/vhar/Request.html:15
	qw422016.N().S(`</td>
        </tr>
        <tr>
          <th class="shrink">URL</th>
          <td><a target="_blank" rel="noopener noreferrer" href="`)
//line views/vhar/Request.html:19
	qw422016.E().S(r.URL)
//line views/vhar/Request.html:19
	qw422016.N().S(`">`)
//line views/vhar/Request.html:19
	qw422016.E().S(r.URL)
//line views/vhar/Request.html:19
	qw422016.N().S(`</a></td>
        </tr>
        <tr>
          <th class="shrink">Headers </th>
          <td>`)
//line views/vhar/Request.html:23
	streamrenderNVPsHidden(qw422016, fmt.Sprintf("request-header-%s", key), "Header", r.Headers, r.HeadersSize, ps)
//line views/vhar/Request.html:23
	qw422016.N().S(`</td>
        </tr>
`)
//line views/vhar/Request.html:25
	if len(r.QueryString) > 0 {
//line views/vhar/Request.html:25
		qw422016.N().S(`        <tr>
          <th class="shrink">Query String</th>
          <td>`)
//line views/vhar/Request.html:28
		streamrenderNVPsHidden(qw422016, fmt.Sprintf("request-query-string-%s", key), "Query String", r.QueryString, 0, ps)
//line views/vhar/Request.html:28
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vhar/Request.html:30
	}
//line views/vhar/Request.html:31
	if len(r.Cookies) > 0 {
//line views/vhar/Request.html:31
		qw422016.N().S(`        <tr>
          <th class="shrink">Cookies</th>
          <td>`)
//line views/vhar/Request.html:34
		streamrenderCookiesHidden(qw422016, fmt.Sprintf("request-cookies-%s", key), r.Cookies, ps)
//line views/vhar/Request.html:34
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vhar/Request.html:36
	}
//line views/vhar/Request.html:37
	if r.PostData != nil {
//line views/vhar/Request.html:37
		qw422016.N().S(`        <tr>
          <th class="shrink">Body</th>
          <td>
            <div class="right"><em>`)
//line views/vhar/Request.html:41
		if r.PostData.MimeType != "" {
//line views/vhar/Request.html:41
			qw422016.E().S(r.PostData.MimeType)
//line views/vhar/Request.html:41
			qw422016.N().S(`, `)
//line views/vhar/Request.html:41
		}
//line views/vhar/Request.html:41
		qw422016.E().S(util.ByteSizeSI(int64(r.BodySize)))
//line views/vhar/Request.html:41
		qw422016.N().S(`</em></div>
`)
//line views/vhar/Request.html:42
		if len(r.PostData.Params) > 0 {
//line views/vhar/Request.html:42
			qw422016.N().S(`            <ul>
`)
//line views/vhar/Request.html:44
			for _, param := range r.PostData.Params {
//line views/vhar/Request.html:44
				qw422016.N().S(`              <li><strong>`)
//line views/vhar/Request.html:45
				qw422016.E().S(param.Name)
//line views/vhar/Request.html:45
				qw422016.N().S(`</strong>: `)
//line views/vhar/Request.html:45
				qw422016.E().S(param.Value)
//line views/vhar/Request.html:45
				qw422016.N().S(`</li>
`)
//line views/vhar/Request.html:46
			}
//line views/vhar/Request.html:46
			qw422016.N().S(`            </ul>
`)
//line views/vhar/Request.html:48
		} else if r.PostData.Text != "" {
//line views/vhar/Request.html:48
			qw422016.N().S(`            <ul class="accordion">
              <li>
                <input id="accordion-request-body-`)
//line views/vhar/Request.html:51
			qw422016.E().S(key)
//line views/vhar/Request.html:51
			qw422016.N().S(`" type="checkbox" hidden="hidden" />
                <label class="no-padding" for="accordion-request-body-`)
//line views/vhar/Request.html:52
			qw422016.E().S(key)
//line views/vhar/Request.html:52
			qw422016.N().S(`"><em>(click to show)</em></label>
                <div class="bd"><div><div>
                  <pre>`)
//line views/vhar/Request.html:54
			qw422016.E().S(r.PostData.Text)
//line views/vhar/Request.html:54
			qw422016.N().S(`</pre>
                </div></div></div>
              </li>
            </ul>
`)
//line views/vhar/Request.html:58
		}
//line views/vhar/Request.html:58
		qw422016.N().S(`          </td>
        </tr>
`)
//line views/vhar/Request.html:61
	}
//line views/vhar/Request.html:61
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vhar/Request.html:65
}

//line views/vhar/Request.html:65
func WriteRenderRequest(qq422016 qtio422016.Writer, key string, r *har.Request, ps *cutil.PageState) {
//line views/vhar/Request.html:65
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Request.html:65
	StreamRenderRequest(qw422016, key, r, ps)
//line views/vhar/Request.html:65
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Request.html:65
}

//line views/vhar/Request.html:65
func RenderRequest(key string, r *har.Request, ps *cutil.PageState) string {
//line views/vhar/Request.html:65
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Request.html:65
	WriteRenderRequest(qb422016, key, r, ps)
//line views/vhar/Request.html:65
	qs422016 := string(qb422016.B)
//line views/vhar/Request.html:65
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Request.html:65
	return qs422016
//line views/vhar/Request.html:65
}

//line views/vhar/Request.html:67
func streamrenderCookiesHidden(qw422016 *qt422016.Writer, key string, cookies har.Cookies, ps *cutil.PageState) {
//line views/vhar/Request.html:67
	qw422016.N().S(`
  <ul class="accordion">
    <li>
      <input id="accordion-cookies-`)
//line views/vhar/Request.html:70
	qw422016.E().S(key)
//line views/vhar/Request.html:70
	qw422016.N().S(`" type="checkbox" hidden="hidden" />
      <label class="no-padding" for="accordion-cookies-`)
//line views/vhar/Request.html:71
	qw422016.E().S(key)
//line views/vhar/Request.html:71
	qw422016.N().S(`">`)
//line views/vhar/Request.html:71
	qw422016.E().S(util.StringPlural(len(cookies), "Cookie"))
//line views/vhar/Request.html:71
	qw422016.N().S(` <em>(click to show)</em></label>
      <div class="bd"><div><div>
        `)
//line views/vhar/Request.html:73
	streamrenderCookies(qw422016, key, cookies, ps)
//line views/vhar/Request.html:73
	qw422016.N().S(`
      </div></div></div>
    </li>
  </ul>
`)
//line views/vhar/Request.html:77
}

//line views/vhar/Request.html:77
func writerenderCookiesHidden(qq422016 qtio422016.Writer, key string, cookies har.Cookies, ps *cutil.PageState) {
//line views/vhar/Request.html:77
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Request.html:77
	streamrenderCookiesHidden(qw422016, key, cookies, ps)
//line views/vhar/Request.html:77
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Request.html:77
}

//line views/vhar/Request.html:77
func renderCookiesHidden(key string, cookies har.Cookies, ps *cutil.PageState) string {
//line views/vhar/Request.html:77
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Request.html:77
	writerenderCookiesHidden(qb422016, key, cookies, ps)
//line views/vhar/Request.html:77
	qs422016 := string(qb422016.B)
//line views/vhar/Request.html:77
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Request.html:77
	return qs422016
//line views/vhar/Request.html:77
}

//line views/vhar/Request.html:79
func streamrenderCookies(qw422016 *qt422016.Writer, key string, cs har.Cookies, ps *cutil.PageState) {
//line views/vhar/Request.html:79
	qw422016.N().S(`
  <div class="overflow full-width">
    <table>
      <thead>
        <tr>
          <th class="shrink">Name</th>
          <th class="shrink">Value</th>
          <th class="shrink">Path</th>
          <th class="shrink">Domain</th>
          <th class="shrink">Expires</th>
          <th>Tags</th>
        </tr>
      </thead>
      <tbody>
`)
//line views/vhar/Request.html:93
	for i, c := range cs {
//line views/vhar/Request.html:93
		qw422016.N().S(`        <tr>
          <td>`)
//line views/vhar/Request.html:95
		qw422016.E().S(c.Name)
//line views/vhar/Request.html:95
		qw422016.N().S(`</td>
          <td>
`)
//line views/vhar/Request.html:97
		if len(c.Value) > 64 {
//line views/vhar/Request.html:97
			qw422016.N().S(`            <ul class="accordion">
              <li>
                <input id="accordion-`)
//line views/vhar/Request.html:100
			qw422016.E().S(key)
//line views/vhar/Request.html:100
			qw422016.N().S(`-cookie-`)
//line views/vhar/Request.html:100
			qw422016.N().D(i)
//line views/vhar/Request.html:100
			qw422016.N().S(`" type="checkbox" hidden="hidden" />
                <label class="no-padding" for="accordion-`)
//line views/vhar/Request.html:101
			qw422016.E().S(key)
//line views/vhar/Request.html:101
			qw422016.N().S(`-cookie-`)
//line views/vhar/Request.html:101
			qw422016.N().D(i)
//line views/vhar/Request.html:101
			qw422016.N().S(`">`)
//line views/vhar/Request.html:101
			qw422016.E().S(c.Value[:64])
//line views/vhar/Request.html:101
			qw422016.N().S(`...</label>
                <div class="bd"><div><div>
                  `)
//line views/vhar/Request.html:103
			qw422016.E().S(c.Value)
//line views/vhar/Request.html:103
			qw422016.N().S(`
                </div></div></div>
              </li>
            </ul>
`)
//line views/vhar/Request.html:107
		} else {
//line views/vhar/Request.html:107
			qw422016.N().S(`            `)
//line views/vhar/Request.html:108
			qw422016.E().S(c.Value)
//line views/vhar/Request.html:108
			qw422016.N().S(`
`)
//line views/vhar/Request.html:109
		}
//line views/vhar/Request.html:109
		qw422016.N().S(`          </td>
          <td>`)
//line views/vhar/Request.html:111
		qw422016.E().S(c.Path)
//line views/vhar/Request.html:111
		qw422016.N().S(`</td>
          <td>`)
//line views/vhar/Request.html:112
		qw422016.E().S(c.Domain)
//line views/vhar/Request.html:112
		qw422016.N().S(`</td>
          <td title="`)
//line views/vhar/Request.html:113
		qw422016.E().S(util.TimeToFull(c.Exp()))
//line views/vhar/Request.html:113
		qw422016.N().S(`">`)
//line views/vhar/Request.html:113
		qw422016.E().S(c.ExpRelative())
//line views/vhar/Request.html:113
		qw422016.N().S(`</td>
          <td>`)
//line views/vhar/Request.html:114
		qw422016.E().S(util.StringJoin(c.Tags(), ", "))
//line views/vhar/Request.html:114
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vhar/Request.html:116
	}
//line views/vhar/Request.html:116
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vhar/Request.html:120
}

//line views/vhar/Request.html:120
func writerenderCookies(qq422016 qtio422016.Writer, key string, cs har.Cookies, ps *cutil.PageState) {
//line views/vhar/Request.html:120
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Request.html:120
	streamrenderCookies(qw422016, key, cs, ps)
//line views/vhar/Request.html:120
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Request.html:120
}

//line views/vhar/Request.html:120
func renderCookies(key string, cs har.Cookies, ps *cutil.PageState) string {
//line views/vhar/Request.html:120
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Request.html:120
	writerenderCookies(qb422016, key, cs, ps)
//line views/vhar/Request.html:120
	qs422016 := string(qb422016.B)
//line views/vhar/Request.html:120
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Request.html:120
	return qs422016
//line views/vhar/Request.html:120
}

//line views/vhar/Request.html:122
func streamrenderNVPsHidden(qw422016 *qt422016.Writer, key string, title string, nvps har.NVPs, size int, ps *cutil.PageState) {
//line views/vhar/Request.html:122
	qw422016.N().S(`
  <ul class="accordion">
    <li>
      <input id="accordion-`)
//line views/vhar/Request.html:125
	qw422016.E().S(key)
//line views/vhar/Request.html:125
	qw422016.N().S(`" type="checkbox" hidden="hidden" />
      <label class="no-padding" for="accordion-`)
//line views/vhar/Request.html:126
	qw422016.E().S(key)
//line views/vhar/Request.html:126
	qw422016.N().S(`">
`)
//line views/vhar/Request.html:127
	if size > 0 {
//line views/vhar/Request.html:127
		qw422016.N().S(`        <div class="right"><em>`)
//line views/vhar/Request.html:128
		qw422016.E().S(util.ByteSizeSI(int64(size)))
//line views/vhar/Request.html:128
		qw422016.N().S(`</em></div>
`)
//line views/vhar/Request.html:129
	}
//line views/vhar/Request.html:129
	qw422016.N().S(`        `)
//line views/vhar/Request.html:130
	qw422016.E().S(util.StringPlural(len(nvps), title))
//line views/vhar/Request.html:130
	qw422016.N().S(`
        <em>(click to show)</em>
      </label>
      <div class="bd"><div><div>
        `)
//line views/vhar/Request.html:134
	streamrenderNVPs(qw422016, nvps, ps)
//line views/vhar/Request.html:134
	qw422016.N().S(`
      </div></div></div>
    </li>
  </ul>
`)
//line views/vhar/Request.html:138
}

//line views/vhar/Request.html:138
func writerenderNVPsHidden(qq422016 qtio422016.Writer, key string, title string, nvps har.NVPs, size int, ps *cutil.PageState) {
//line views/vhar/Request.html:138
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Request.html:138
	streamrenderNVPsHidden(qw422016, key, title, nvps, size, ps)
//line views/vhar/Request.html:138
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Request.html:138
}

//line views/vhar/Request.html:138
func renderNVPsHidden(key string, title string, nvps har.NVPs, size int, ps *cutil.PageState) string {
//line views/vhar/Request.html:138
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Request.html:138
	writerenderNVPsHidden(qb422016, key, title, nvps, size, ps)
//line views/vhar/Request.html:138
	qs422016 := string(qb422016.B)
//line views/vhar/Request.html:138
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Request.html:138
	return qs422016
//line views/vhar/Request.html:138
}

//line views/vhar/Request.html:140
func streamrenderNVPs(qw422016 *qt422016.Writer, nvps har.NVPs, ps *cutil.PageState) {
//line views/vhar/Request.html:140
	qw422016.N().S(`
  <div class="overflow full-width">
    <table>
      <tbody>
`)
//line views/vhar/Request.html:144
	for _, n := range nvps {
//line views/vhar/Request.html:144
		qw422016.N().S(`        <tr title="`)
//line views/vhar/Request.html:145
		qw422016.E().S(n.Comment)
//line views/vhar/Request.html:145
		qw422016.N().S(`">
          <th class="shrink">`)
//line views/vhar/Request.html:146
		qw422016.E().S(n.Name)
//line views/vhar/Request.html:146
		qw422016.N().S(`</th>
          <td>`)
//line views/vhar/Request.html:147
		qw422016.E().S(n.Value)
//line views/vhar/Request.html:147
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vhar/Request.html:149
	}
//line views/vhar/Request.html:149
	qw422016.N().S(`      </tbody>
    </table>
  </div>
`)
//line views/vhar/Request.html:153
}

//line views/vhar/Request.html:153
func writerenderNVPs(qq422016 qtio422016.Writer, nvps har.NVPs, ps *cutil.PageState) {
//line views/vhar/Request.html:153
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vhar/Request.html:153
	streamrenderNVPs(qw422016, nvps, ps)
//line views/vhar/Request.html:153
	qt422016.ReleaseWriter(qw422016)
//line views/vhar/Request.html:153
}

//line views/vhar/Request.html:153
func renderNVPs(nvps har.NVPs, ps *cutil.PageState) string {
//line views/vhar/Request.html:153
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vhar/Request.html:153
	writerenderNVPs(qb422016, nvps, ps)
//line views/vhar/Request.html:153
	qs422016 := string(qb422016.B)
//line views/vhar/Request.html:153
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vhar/Request.html:153
	return qs422016
//line views/vhar/Request.html:153
}
