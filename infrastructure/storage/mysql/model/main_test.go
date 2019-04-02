package model_test

import (
	"github.com/marugoshi/gobm/shared/app_testutils"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	app_testutils.Setup()
	result := m.Run()
	app_testutils.Teardown()
	os.Exit(result)
}
