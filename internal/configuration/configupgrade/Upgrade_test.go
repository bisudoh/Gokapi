package configupgrade

import (
	"os"
	"testing"

	"github.com/bisudoh/gokapi/internal/environment"
	"github.com/bisudoh/gokapi/internal/models"
	"github.com/bisudoh/gokapi/internal/test"
	"github.com/bisudoh/gokapi/internal/test/testconfiguration"
)

func TestMain(m *testing.M) {
	testconfiguration.Create(false)
	exitVal := m.Run()
	testconfiguration.Delete()
	os.Exit(exitVal)
}

var oldConfigFile = models.Configuration{
	Authentication: models.AuthenticationConfig{},
	Port:           "127.0.0.1:53844",
	ServerUrl:      "https://gokapi.url/",
	RedirectUrl:    "https://github.com/Bisudoh/Gokapi/",
}

func TestUpgradeDb(t *testing.T) {
	exitCode := 0
	osExit = func(code int) {
		exitCode = code
	}
	env := environment.New()
	// Too old to update
	oldConfigFile.ConfigVersion = minConfigVersion - 1
	upgradeDone := DoUpgrade(&oldConfigFile, &env)
	test.IsEqualBool(t, upgradeDone, true)
	test.IsEqualInt(t, exitCode, 1)

	// Updatable version
	exitCode = 0
	oldConfigFile.ConfigVersion = 21
	upgradeDone = DoUpgrade(&oldConfigFile, &env)
	test.IsEqualBool(t, upgradeDone, true)
	// TODO
	test.IsEqualInt(t, exitCode, 0)

	// Current Version
	exitCode = 0
	oldConfigFile.ConfigVersion = CurrentConfigVersion
	upgradeDone = DoUpgrade(&oldConfigFile, &env)
	test.IsEqualBool(t, upgradeDone, false)
	test.IsEqualInt(t, exitCode, 0)

}
