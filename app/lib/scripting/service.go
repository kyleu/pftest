// Content managed by Project Forge, see [projectforge.md] for details.
package scripting

import (
	"path/filepath"
	"strings"

	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/filesystem"
	"github.com/kyleu/pftest/app/util"
)

type Service struct {
	FS   filesystem.FileLoader `json:"-"`
	Path string                `json:"path,omitempty"`
}

func NewService(fs filesystem.FileLoader, pth string) *Service {
	return &Service{FS: fs, Path: pth}
}

func (s *Service) ListScripts(logger util.Logger) []string {
	files, _ := s.FS.ListFilesRecursive(s.Path, nil, logger)
	return lo.FilterMap(files, func(x string, _ int) (string, bool) {
		return x, strings.HasSuffix(x, ".js")
	})
}

func (s *Service) LoadScript(pth string, logger util.Logger) (string, error) {
	logger.Infof("loading script [%s]", pth)
	filePath := filepath.Join(s.Path, pth)
	b, err := s.FS.ReadFile(filePath)
	if err != nil {
		b, err = s.FS.ReadFile(filePath + ".js")
		if err != nil {
			return "", err
		}
	}
	sc := string(b)
	return sc, nil
}

func (s *Service) SaveScript(pth string, content string, logger util.Logger) error {
	logger.Infof("saving script [%s]", pth)
	filePath := filepath.Join(s.Path, pth)
	return s.FS.WriteFile(filePath, []byte(content), filesystem.DefaultMode, true)
}

func (s *Service) DeleteScript(pth string, logger util.Logger) error {
	filePath := filepath.Join(s.Path, pth)
	return s.FS.Remove(filePath, logger)
}
