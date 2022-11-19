// Content managed by Project Forge, see [projectforge.md] for details.
package user

import (
	"strings"

	"golang.org/x/exp/slices"

	"github.com/kyleu/pftest/app/util"
)

type Account struct {
	Provider string `json:"provider"`
	Email    string `json:"email"`
	Token    string `json:"-"`
}

func accountFromString(s string) *Account {
	p, e := util.StringSplit(s, ':', true)
	var t string
	if strings.Contains(e, "|") {
		e, t = util.StringSplit(e, '|', true)
		if decr, err := util.DecryptMessage(nil, t, nil); err == nil {
			t = decr
		}
	}
	return &Account{Provider: p, Email: e, Token: t}
}

func (a Account) String() string {
	ret := a.Provider + ":" + a.Email
	if a.Token != "" {
		if enc, err := util.EncryptMessage(nil, a.Token, nil); err == nil {
			ret += "|" + enc
		} else {
			ret += "|" + a.Token
		}
	}
	return ret
}

func (a Account) TitleString() string {
	return a.Provider + ":" + a.Email
}

type Accounts []*Account

func (a Accounts) String() string {
	ret := make([]string, 0, len(a))
	for _, x := range a {
		ret = append(ret, x.String())
	}
	return strings.Join(ret, ",")
}

func (a Accounts) TitleString() string {
	ret := make([]string, 0, len(a))
	for _, x := range a {
		ret = append(ret, x.TitleString())
	}
	return strings.Join(ret, ",")
}

func (a Accounts) Sort() {
	slices.SortFunc(a, func(l *Account, r *Account) bool {
		if l.Provider == r.Provider {
			return strings.ToLower(l.Email) < strings.ToLower(r.Email)
		}
		return l.Provider < r.Provider
	})
}

func (a Accounts) GetByProvider(p string) Accounts {
	var ret Accounts
	for _, x := range a {
		if x.Provider == p {
			ret = append(ret, x)
		}
	}
	return ret
}

func (a Accounts) Matches(match string) bool {
	if match == "" || match == "*" {
		return true
	}
	if strings.Contains(match, ",") {
		xs := util.StringSplitAndTrim(match, ",")
		for _, x := range xs {
			if a.Matches(x) {
				return true
			}
		}
		return false
	}
	prv, acct := util.StringSplit(match, ':', true)
	for _, x := range a {
		if x.Provider == prv {
			if acct == "" {
				return true
			}
			return strings.HasSuffix(x.Email, acct)
		}
	}
	return false
}

func (a Accounts) Purge(keys ...string) Accounts {
	ret := make(Accounts, 0, len(a))
	for _, ss := range a {
		hit := false
		for _, key := range keys {
			if ss.Provider == key {
				hit = true
			}
		}
		if !hit {
			ret = append(ret, ss)
		}
	}
	return ret
}

func AccountsFromString(s string) Accounts {
	split := util.StringSplitAndTrim(s, ",")
	ret := make(Accounts, 0, len(split))
	for _, x := range split {
		ret = append(ret, accountFromString(x))
	}
	return ret
}
