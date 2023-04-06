// Content managed by Project Forge, see [projectforge.md] for details.
package cutil

import (
	"context"
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/cmenu"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/lib/menu"
	"github.com/kyleu/pftest/app/lib/telemetry"
	"github.com/kyleu/pftest/app/lib/theme"
	"github.com/kyleu/pftest/app/lib/user"
	"github.com/kyleu/pftest/app/util"
)

const (
	DefaultSearchPath  = "/search"
	DefaultProfilePath = "/profile"
	defaultIcon        = "app"
)

var (
	defaultRootTitleAppend = util.GetEnv("app_display_name_append")
	defaultRootTitle       = func() string {
		if tmp := util.GetEnv("app_display_name"); tmp != "" {
			return tmp
		}
		return util.AppName
	}()
)

type PageState struct {
	Title         string            `json:"title,omitempty"`
	Description   string            `json:"description,omitempty"`
	Method        string            `json:"method,omitempty"`
	URI           *fasthttp.URI     `json:"-"`
	Menu          menu.Items        `json:"menu,omitempty"`
	Breadcrumbs   cmenu.Breadcrumbs `json:"breadcrumbs,omitempty"`
	Flashes       []string          `json:"flashes,omitempty"`
	Session       util.ValueMap     `json:"-"`
	Profile       *user.Profile     `json:"profile,omitempty"`
	Accounts      user.Accounts     `json:"accounts,omitempty"`
	Authed        bool              `json:"authed,omitempty"`
	Admin         bool              `json:"admin,omitempty"`
	Params        filter.ParamSet   `json:"params,omitempty"`
	Icons         []string          `json:"icons,omitempty"`
	RootIcon      string            `json:"rootIcon,omitempty"`
	RootPath      string            `json:"rootPath,omitempty"`
	RootTitle     string            `json:"rootTitle,omitempty"`
	SearchPath    string            `json:"searchPath,omitempty"`
	ProfilePath   string            `json:"profilePath,omitempty"`
	HideMenu      bool              `json:"hideMenu,omitempty"`
	ForceRedirect string            `json:"forceRedirect,omitempty"`
	HeaderContent string            `json:"headerContent,omitempty"`
	Data          any               `json:"data,omitempty"`
	Logger        util.Logger       `json:"-"`
	Context       context.Context   `json:"-"` //nolint:containedctx // properly closed, never directly used
	Span          *telemetry.Span   `json:"-"`
	RenderElapsed float64           `json:"renderElapsed,omitempty"`
}

func (p *PageState) AddIcon(keys ...string) {
	for _, k := range keys {
		var hit bool
		for _, icon := range p.Icons {
			if icon == k {
				hit = true
				break
			}
		}
		if !hit {
			p.Icons = append(p.Icons, k)
		}
	}
}

func (p *PageState) TitleString() string {
	if p.Title == "" {
		return util.AppName
	}
	return fmt.Sprintf("%s - %s", p.Title, util.AppName)
}

func (p *PageState) Username() string {
	return p.Profile.Name
}

func (p *PageState) AuthString() string {
	n := p.Profile.String()
	msg := fmt.Sprintf("signed in as %s", n)
	if len(p.Accounts) == 0 {
		if n == user.DefaultProfile.Name {
			return "click to sign in"
		}
		return msg
	}
	return fmt.Sprintf("%s using [%s]", msg, p.Accounts.TitleString())
}

func (p *PageState) Clean(_ *fasthttp.RequestCtx, as *app.State) error {
	if p.Profile != nil && p.Profile.Theme == "" {
		p.Profile.Theme = theme.ThemeDefault.Key
	}
	if p.RootIcon == "" {
		p.RootIcon = defaultIcon
	}
	if p.RootPath == "" {
		p.RootPath = "/"
	}
	if p.RootTitle == "" {
		p.RootTitle = defaultRootTitle
	}
	if defaultRootTitleAppend != "" {
		p.RootTitle += " " + defaultRootTitleAppend
	}
	if p.SearchPath == "" {
		p.SearchPath = DefaultSearchPath
	}
	if p.ProfilePath == "" {
		p.ProfilePath = DefaultProfilePath
	}
	if len(p.Menu) == 0 {
		m, data, err := cmenu.MenuFor(p.Context, p.Authed, p.Admin, p.Profile, p.Params, as, p.Logger)
		if err != nil {
			return err
		}
		if data != nil && p.Data == nil {
			p.Data = data
		}
		p.Menu = m
	}
	return nil
}

func (p *PageState) Close() {
	if p.Span != nil {
		p.Span.Complete()
	}
}
