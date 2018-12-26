package app

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	appTestOption := SetupAppTestOption()

	status := m.Run()

	defer func() {
		appTestOption.Cleanup()
		os.Exit(status)
	}()
}
