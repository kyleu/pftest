package notebook

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/filesystem"
	"github.com/kyleu/pftest/app/util"
)

var FavoritePages = util.NewOrderedMap[string](false, 10)

type Service struct {
	BaseURL string                `json:"baseURL"`
	FS      filesystem.FileLoader `json:"-"`
}

func NewService() *Service {
	baseURL := util.GetEnv("notebook_base_url", fmt.Sprintf("http://localhost:%d", util.AppPort+10))
	baseURL = strings.TrimSuffix(baseURL, "/")
	fs, _ := filesystem.NewFileSystem("notebook/docs", false, "")
	return &Service{BaseURL: baseURL, FS: fs}
}

func (s *Service) Status() string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.BaseURL, http.NoBody)
	if err != nil {
		return "internal-error"
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "not-started"
	}
	if rsp.StatusCode != 200 {
		return "invalid-response"
	}
	return "running"
}

func (s *Service) Start() error {
	if s.Status() == "running" {
		return errors.Errorf("can't start notebook, something is already listening on port [%d]", util.AppPort+10)
	}
	_, err := util.StartProcess("bin/dev.sh", "./notebook", nil, nil, nil)
	if err != nil {
		return err
	}
	return nil
}
