package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	buildNo   string
	commit    string
	branch    string
	version   string
	startTime string
)

func main() {
	flag.StringVar(&buildNo, "buildNo", "0", "Jenkin build number")
	flag.Parse()

	commit = runCommand("NA", "git", "rev-parse", "--short", "HEAD")
	if commit == "" {
		commit = "NA"
	}
	branch = runCommand("master", "git", "rev-parse", "--abbrev-ref", "HEAD")
	if branch == "" {
		branch = "master"
	}
	version = os.Getenv("APP_AGENT")
	if version == "" {
		version = "1.0.0"
	}
	now := time.Now()
	startTime = fmt.Sprintf("%d", now.Unix())

	build("./bin/project", "./cmd")
}

func build(binaryName string, packageAddr string) {
	ldFlags := fmt.Sprintf("-w -s -X main.buildNo=%s -X main.version=%s -X main.commit=%s -X main.branch=%s -X main.startTime=%s",
		buildNo, version, commit, branch, startTime)

	args := []string{"build", "-ldflags", ldFlags, "-o", binaryName, packageAddr}

	output := runCommand("", "go", args...)
	fmt.Println("Output of build command : %s", output)
}

func runCommand(defaultResponse string, cmd string, args ...string) string {
	command := exec.Command(cmd, args...)
	fmt.Println(command)
	output, err := command.Output()
	if err != nil || len(output) == 0 {
		fmt.Println(err)
		return defaultResponse
	}
	return string(bytes.Trim(output, "\n"))
}
