// Basic Sample Go Project API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost:3002
//     BasePath: /api/v1
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Mitesh Sharma<godev@goproject.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - AuthKey:
//
//     SecurityDefinitions:
//     AuthKey:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package main

import (
	"net"
	"os"
	"os/signal"

	"github.com/MiteshSharma/project/setting"

	"github.com/MiteshSharma/project/cmd/server"

	"github.com/MiteshSharma/project/logger"
)

var version = "1.0.0"
var commit = ""
var branch = "master"
var startTime = ""
var buildNo = "0"

func main() {
	setting := setting.NewSetting(buildNo, version, commit, branch, startTime)
	server := server.NewServer(setting)
	defer server.StopServer()

	server.StartServer()

	err := sendSystemdNotification()

	if err != nil {
		server.Log.Debug("Systemd notification error.", logger.Error(err))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	server.Log.Debug("Stopped server.")

	server.StopServer()

	os.Exit(0)
}

func sendSystemdNotification() error {
	notifySocket := os.Getenv("NOTIFY_SOCKET")
	if notifySocket != "" {
		state := "READY=1"
		socketAddr := &net.UnixAddr{
			Name: notifySocket,
			Net:  "unixgram",
		}
		conn, err := net.DialUnix(socketAddr.Net, nil, socketAddr)
		if err != nil {
			return err
		}
		defer conn.Close()
		_, err = conn.Write([]byte(state))
		return err
	}
	return nil
}
