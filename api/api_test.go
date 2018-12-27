package api

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting main test case")
	apiTest := SetupApiTest()
	fmt.Println("Setup api test complete")

	status := m.Run()

	defer func() {
		fmt.Println("Cleanup api test")
		apiTest.CleanUpApiTest()
		os.Exit(status)
	}()
}
