package filter

import (
	"strings"

	"go.uber.org/zap"
)

type ParamSet map[string]*Params

func (s ParamSet) Get(key string, allowed []string, logger *zap.SugaredLogger) *Params {
	x, ok := s[key]
	if !ok {
		return &Params{Key: key}
	}
	return x.Filtered(allowed, logger).Sanitize(key)
}

func (s ParamSet) String() string {
	ret := make([]string, 0, len(s))
	for _, p := range s {
		ret = append(ret, p.String())
	}

	return strings.Join(ret, ", ")
}
