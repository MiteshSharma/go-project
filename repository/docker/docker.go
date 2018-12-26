package docker

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	dockerStatusExited   = "exited"
	dockerStatusRunning  = "running"
	dockerStatusStarting = "starting"
)

type Docker struct {
	ContainerID   string
	ContainerName string
	Option        ContainerOption
}

type ContainerOption struct {
	Name              string
	ContainerFileName string
	Options           map[string]string
	MountVolumePath   string
	PortExpose        string
}

func (d *Docker) isInstalled() bool {
	command := exec.Command("docker", "ps")
	err := command.Run()
	if err != nil {
		return false
	}
	return true
}

func (d *Docker) Start(c ContainerOption) (string, error) {
	d.Option = c
	dockerArgs := d.getDockerRunOptions(c)
	command := exec.Command("docker", dockerArgs...)
	command.Stderr = os.Stderr
	result, err := command.Output()
	if err != nil {
		return "", err
	}
	d.ContainerID = strings.TrimSpace(string(result))
	d.ContainerName = c.Name
	command = exec.Command("docker", "inspect", d.ContainerID)
	result, err = command.Output()
	if err != nil {
		d.Stop()
		return "", err
	}
	return string(result), nil
}

func (d *Docker) WaitForStartOrKill(timeout int) error {
	for tick := 0; tick < timeout; tick++ {
		containerStatus := d.getContainerStatus()
		if containerStatus == dockerStatusRunning {
			return nil
		}
		if containerStatus == dockerStatusExited {
			return nil
		}
		time.Sleep(time.Second)
	}
	d.Stop()
	return errors.New("Docker fail to start in given time period so stopped")
}

func (d *Docker) getContainerStatus() string {
	command := exec.Command("docker", "ps", "-a", "--format", "{{.ID}}|{{.Status}}|{{.Ports}}|{{.Names}}")
	output, err := command.CombinedOutput()
	if err != nil {
		d.Stop()
		return dockerStatusExited
	}
	outputString := string(output)
	outputString = strings.TrimSpace(outputString)
	dockerPsResponse := strings.Split(outputString, "\n")
	for _, response := range dockerPsResponse {
		containerStatusData := strings.Split(response, "|")
		containerStatus := containerStatusData[1]
		containerName := containerStatusData[3]
		if containerName == d.ContainerName {
			if strings.HasPrefix(containerStatus, "Up ") {
				return dockerStatusRunning
			}
		}
	}
	return dockerStatusStarting
}

func (d *Docker) WaitForPortOpen(timeout int) error {
	for tick := 0; tick < timeout; tick += 2 {
		err := d.checkIfPortOpen()
		if err == nil {
			return nil
		}
		time.Sleep(time.Second)
	}
	return errors.New("Docker fail to open port in given time period")
}

func (d *Docker) checkIfPortOpen() error {
	conn, err := net.DialTimeout("tcp", "localhost:"+d.Option.PortExpose, time.Second)
	defer conn.Close()
	if err != nil {
		return err
	}
	one := []byte{}
	conn.SetReadDeadline(time.Now().Add(time.Millisecond * 100))
	_, err = conn.Read(one)
	if err != nil {
		return err
	}
	return nil
}

func (d *Docker) getDockerRunOptions(c ContainerOption) []string {
	portExpose := fmt.Sprintf("%s:%s", c.PortExpose, c.PortExpose)
	var args []string
	for key, value := range c.Options {
		args = append(args, []string{"-e", fmt.Sprintf("%s=%s", key, value)}...)
	}
	args = append(args, []string{"--tmpfs", c.MountVolumePath, c.ContainerFileName}...)
	dockerArgs := append([]string{"run", "-d", "--name", c.Name, "-p", portExpose}, args...)
	return dockerArgs
}

func (d *Docker) Stop() {
	exec.Command("docker", "rm", "-f", d.ContainerID).Run()
}
