// Code generated by qtc from "Profile.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vprofile/Profile.html:2
package vprofile

//line views/vprofile/Profile.html:2
import (
	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cutil"
	"github.com/kyleu/pftest/app/lib/auth"
	"github.com/kyleu/pftest/app/lib/theme"
	"github.com/kyleu/pftest/app/lib/user"
	"github.com/kyleu/pftest/views/components"
	"github.com/kyleu/pftest/views/layout"
	"github.com/kyleu/pftest/views/vauth"
	"github.com/kyleu/pftest/views/vtheme"
)

//line views/vprofile/Profile.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vprofile/Profile.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vprofile/Profile.html:14
type Profile struct {
	layout.Basic
	Profile   *user.Profile
	Theme     *theme.Theme
	Providers auth.Providers
	Referrer  string
}

//line views/vprofile/Profile.html:22
func (p *Profile) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vprofile/Profile.html:22
	qw422016.N().S(`
  `)
//line views/vprofile/Profile.html:23
	vauth.StreamSigninTable(qw422016, p.Providers, p.Referrer, as, ps)
//line views/vprofile/Profile.html:23
	qw422016.N().S(`
  <form action="" method="post">
    <input name="referrer" type="hidden" value="`)
//line views/vprofile/Profile.html:25
	qw422016.E().S(p.Referrer)
//line views/vprofile/Profile.html:25
	qw422016.N().S(`" />
    <div class="card">
      <div class="right">
        <a href="#modal-profile"><button type="button">Profile</button></a>
        <a href="#modal-theme"><button type="button">Theme</button></a>
      </div>
      <h3>Profile</h3>
      <table class="mt expanded">
        <tbody>
          `)
//line views/vprofile/Profile.html:34
	components.StreamTableInput(qw422016, "name", "Name", p.Profile.Name, 5)
//line views/vprofile/Profile.html:34
	qw422016.N().S(`
          <tr>
            <th class="shrink"><label>Mode</label></th>
            <td>
              <label>
`)
//line views/vprofile/Profile.html:39
	if p.Profile.Mode == "" {
//line views/vprofile/Profile.html:39
		qw422016.N().S(`                <input type="radio" class="mode-input" name="mode" value="" checked="checked" />
`)
//line views/vprofile/Profile.html:41
	} else {
//line views/vprofile/Profile.html:41
		qw422016.N().S(`                <input type="radio" class="mode-input" name="mode" value="" />
`)
//line views/vprofile/Profile.html:43
	}
//line views/vprofile/Profile.html:43
	qw422016.N().S(`                System Default
              </label>
              <label>
`)
//line views/vprofile/Profile.html:47
	if p.Profile.Mode == "light" {
//line views/vprofile/Profile.html:47
		qw422016.N().S(`                <input type="radio" class="mode-input" name="mode" value="light" checked="checked" />
`)
//line views/vprofile/Profile.html:49
	} else {
//line views/vprofile/Profile.html:49
		qw422016.N().S(`                <input type="radio" class="mode-input" name="mode" value="light" />
`)
//line views/vprofile/Profile.html:51
	}
//line views/vprofile/Profile.html:51
	qw422016.N().S(`                Light
              </label>
              <label>
`)
//line views/vprofile/Profile.html:55
	if p.Profile.Mode == "dark" {
//line views/vprofile/Profile.html:55
		qw422016.N().S(`                <input type="radio" class="mode-input" name="mode" value="dark" checked="checked" />
`)
//line views/vprofile/Profile.html:57
	} else {
//line views/vprofile/Profile.html:57
		qw422016.N().S(`                <input type="radio" class="mode-input" name="mode" value="dark" />
`)
//line views/vprofile/Profile.html:59
	}
//line views/vprofile/Profile.html:59
	qw422016.N().S(`                Dark
              </label>
            </td>
          </tr>
          `)
//line views/vprofile/Profile.html:64
	vtheme.StreamChoicePanel(qw422016, as.Themes.All(ps.Logger), 5, ps)
//line views/vprofile/Profile.html:64
	qw422016.N().S(`
        </tbody>
      </table>
      <div class="mt">
        <button type="submit">Save All Changes</button>
        <button type="reset">Reset</button>
      </div>
    </div>
  </form>

  `)
//line views/vprofile/Profile.html:74
	components.StreamJSONModal(qw422016, "profile", "Profile JSON", p.Profile, 1)
//line views/vprofile/Profile.html:74
	qw422016.N().S(`
  `)
//line views/vprofile/Profile.html:75
	components.StreamJSONModal(qw422016, "theme", "Theme JSON", p.Theme, 1)
//line views/vprofile/Profile.html:75
	qw422016.N().S(`
`)
//line views/vprofile/Profile.html:76
}

//line views/vprofile/Profile.html:76
func (p *Profile) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vprofile/Profile.html:76
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vprofile/Profile.html:76
	p.StreamBody(qw422016, as, ps)
//line views/vprofile/Profile.html:76
	qt422016.ReleaseWriter(qw422016)
//line views/vprofile/Profile.html:76
}

//line views/vprofile/Profile.html:76
func (p *Profile) Body(as *app.State, ps *cutil.PageState) string {
//line views/vprofile/Profile.html:76
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vprofile/Profile.html:76
	p.WriteBody(qb422016, as, ps)
//line views/vprofile/Profile.html:76
	qs422016 := string(qb422016.B)
//line views/vprofile/Profile.html:76
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vprofile/Profile.html:76
	return qs422016
//line views/vprofile/Profile.html:76
}