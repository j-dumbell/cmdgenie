package cli

import (
	"path"
	"testing"

	"github.com/j-dumbell/cmdgenie/internal/config"
)

func TestApp_configure(t *testing.T) {
	configService := newTestConfigService(t)
	NewApp(configService)

}

func newTestConfigService(t *testing.T) config.Service {
	dir := t.TempDir()
	filePath := path.Join(dir, ".cmdgenie.json")
	return config.NewService(filePath)
}
