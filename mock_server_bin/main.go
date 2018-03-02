// Run Graylog mock server.
// Usage
// $ graylog-mock-server [--port <port number>] [--log-level debug|info|warn|error|fatal|panic] [--data <data-file-path>]
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/urfave/cli"
)

func action(c *cli.Context) error {
	var (
		server *graylog.MockServer
		err    error
	)
	port := c.Int("port")
	if port == 0 {
		server, err = graylog.NewMockServer("")
	} else {
		server, err = graylog.NewMockServer(fmt.Sprintf(":%d", c.Int("port")))
	}
	if err != nil {
		return cli.NewExitError(
			fmt.Sprintf("Failed to Get Mock Server: %s", err), 1)
	}
	ll := c.String("log-level")
	lvl, err := log.ParseLevel(ll)
	if err != nil {
		return cli.NewExitError(fmt.Sprintf(
			`Invalid log-level %s. log-level must be any of debug|info|warn|error|fatal|panic`, ll), 1)
	}

	server.Logger.SetLevel(lvl)
	server.DataPath = c.String("data")
	if err := server.Load(); err != nil {
		return cli.NewExitError(
			fmt.Sprintf("Failed to load data: %s", err), 1)
	}
	server.Start()
	defer server.Close()
	server.Logger.Infof(
		"Start Graylog mock server: %s\nCtrl + C to stop server", server.Endpoint)
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(
		signal_chan, syscall.SIGHUP, syscall.SIGINT,
		syscall.SIGTERM, syscall.SIGQUIT)
	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			default:
				exit_chan <- 0
			}
		}
	}()

	<-exit_chan
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "graylog-mock-server"
	app.Usage = "Run Graylog mock server."
	app.Action = action
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port",
			Usage: "port number. If you don't set this option, a free port is assigned and the assigned port number is outputed to the console when the mock server runs.",
		},
		cli.StringFlag{
			Name:  "log-level",
			Value: "info",
			Usage: "the log level of logrus which the mock server uses internally.",
		},
		cli.StringFlag{
			Name:  "data",
			Usage: "data file path. When the server runs data of the file is loaded and when data of the server is changed data is saved at the file. If this option is not set, no data is loaded and saved.",
		},
		// cli.StringFlag{
		// 	Name: "log-format",
		// },
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
