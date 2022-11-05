// Content managed by Project Forge, see [projectforge.md] for details.
package cutil

import (
	"context"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/pftest/app"
	"github.com/kyleu/pftest/app/controller/csession"
	"github.com/kyleu/pftest/app/lib/telemetry"
	"github.com/kyleu/pftest/app/lib/telemetry/httpmetrics"
	"github.com/kyleu/pftest/app/lib/user"
	"github.com/kyleu/pftest/app/util"
)

var initialIcons = []string{"searchbox"}

func LoadPageState(as *app.State, rc *fasthttp.RequestCtx, key string, logger util.Logger) *PageState {
	ctx, logger := httpmetrics.ExtractHeaders(rc, logger)
	traceCtx, span, logger := telemetry.StartSpan(ctx, "http:"+key, logger)
	span.Attribute("path", string(rc.Request.URI().Path()))
	httpmetrics.InjectHTTP(rc, span)

	session, flashes, prof, accts := loadSession(ctx, as, rc, logger)
	params := ParamSetFromRequest(rc)

	isAuthed, _ := user.Check("/", accts)
	isAdmin, _ := user.Check("/admin", accts)

	return &PageState{
		Method: string(rc.Method()), URI: rc.Request.URI(), Flashes: flashes, Session: session,
		Profile: prof, Accounts: accts, Authed: isAuthed, Admin: isAdmin, Params: params,
		Icons: initialIcons, Context: traceCtx, Span: span, Logger: logger,
	}
}

func loadSession(ctx context.Context, as *app.State, rc *fasthttp.RequestCtx, logger util.Logger) (util.ValueMap, []string, *user.Profile, user.Accounts) {
	sessionBytes := rc.Request.Header.Cookie(util.AppKey)
	session := util.ValueMap{}
	if len(sessionBytes) > 0 {
		dec, err := util.DecryptMessage(nil, string(sessionBytes), logger)
		if err != nil {
			logger.Warnf("error decrypting session: %+v", err)
		}
		err = util.FromJSON([]byte(dec), &session)
		if err != nil {
			session = util.ValueMap{}
		}
	}

	flashes := util.StringSplitAndTrim(session.GetStringOpt(csession.WebFlashKey), ";")
	if len(flashes) > 0 {
		delete(session, csession.WebFlashKey)
		err := csession.SaveSession(rc, session, logger)
		if err != nil {
			logger.Warnf("can't save session: %+v", err)
		}
	}

	prof, err := loadProfile(session)
	if err != nil {
		logger.Warnf("can't load profile: %+v", err)
	}

	var accts user.Accounts
	authX, ok := session[csession.WebAuthKey]
	if ok {
		authS, ok := authX.(string)
		if ok {
			accts = user.AccountsFromString(authS)
		}
	}

	return session, flashes, prof, accts
}

func loadProfile(session util.ValueMap) (*user.Profile, error) {
	x, ok := session["profile"]
	if !ok {
		return user.DefaultProfile.Clone(), nil
	}
	s, ok := x.(string)
	if !ok {
		m, ok := x.(map[string]any)
		if !ok {
			return user.DefaultProfile.Clone(), nil
		}
		s = util.ToJSON(m)
	}
	p := &user.Profile{}
	err := util.FromJSON([]byte(s), p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
