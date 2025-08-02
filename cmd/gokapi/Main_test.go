//go:build !integration && test

package main

import (
	"os"
	"testing"

	"github.com/bisudoh/gokapi/internal/environment/flagparser"
	"github.com/bisudoh/gokapi/internal/test"
	"github.com/bisudoh/gokapi/internal/test/testconfiguration"
)

func TestMain(m *testing.M) {
	testconfiguration.Create(false)
	exitVal := m.Run()
	testconfiguration.Delete()
	os.Exit(exitVal)
}

func TestShowVersion(t *testing.T) {
	showVersion(flagparser.MainFlags{})
	osExit = test.ExitCode(t, 0)
	showVersion(flagparser.MainFlags{ShowVersion: true})
}

func TestNoResetPw(t *testing.T) {
	reconfigureServer(flagparser.MainFlags{})
}

func TestCreateSsl(t *testing.T) {
	test.FileDoesNotExist(t, "test/ssl.key")
	createSsl(flagparser.MainFlags{})
	test.FileDoesNotExist(t, "test/ssl.key")
	createSsl(flagparser.MainFlags{CreateSsl: true})
	test.FileExists(t, "test/ssl.key")
}
