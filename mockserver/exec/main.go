// Run Graylog mock server.
//
// Usage
//   $ graylog-mock-server [--port <port number>] [--log-level debug|info|warn|error|fatal|panic] [--data <data-file-path>]
package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog/mockserver"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/plain"
)

const version = "0.5.0"

var help string

func init() {
	help = fmt.Sprintf(`
graylog-mock-server - Run Graylog mock server.

USAGE:
   graylog-mock-server [options]

VERSION:
   %s

OPTIONS:
   --port value       port number. If you don't set this option, a free port is assigned and the assigned port number is output to the console when the mock server runs.
   --log-level value  the log level of logrus which the mock server uses internally. (default: "info")
   --data value       data file path. When the server runs data of the file is loaded and when data of the server is changed data is saved at the file. If this option is not set, no data is loaded and saved.
   --help, -h         show help
   --version, -v      print the version
`, version)
}

func action(dataPath, logLevel string, port int) error {
	var (
		server *mockserver.Server
		err    error
	)
	if port == 0 {
		server, err = mockserver.NewServer(
			"", plain.NewStore(dataPath))
	} else {
		server, err = mockserver.NewServer(
			fmt.Sprintf(":%d", port), plain.NewStore(dataPath))
	}
	if err != nil {
		return errors.Wrap(err, "failed to create a mock server")
	}
	lvl, err := log.ParseLevel(logLevel)
	if err != nil {
		return fmt.Errorf(
			`invalid log-level %s.
log-level must be any of debug|info|warn|error|fatal|panic`, logLevel)
	}

	server.Logger().SetLevel(lvl)
	if err := server.Load(); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to load data at %s", dataPath))
	}
	server.Start()
	defer server.Close()
	server.Logger().Infof(
		"Start Graylog mock server: %s\nCtrl + C to stop server", server.Endpoint())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan, syscall.SIGHUP, syscall.SIGINT,
		syscall.SIGTERM, syscall.SIGQUIT)
	exitChan := make(chan int)
	go func() {
		for {
			<-signalChan
			exitChan <- 0
		}
	}()

	<-exitChan
	return nil
}

func main() {
	var portFlag = flag.Int(
		"port", 0,
		"port number. If you don't set this option, a free port is assigned and the assigned port number is output to the console when the mock server runs.")
	var dataFlag = flag.String(
		"data", "",
		"data file path. When the server runs data of the file is loaded and when data of the server is changed data is saved at the file. If this option is not set, no data is loaded and saved.")
	var logLevelFlag = flag.String(
		"log-level", "info",
		`the log level of logrus which the mock server uses internally. (default: "info")`)
	var helpFlag = flag.Bool("help", false, "Show help.")
	var versionFlag = flag.Bool("version", false, "Print the version.")
	flag.Parse()

	if *helpFlag {
		fmt.Println(help)
		return
	}
	if *versionFlag {
		fmt.Println(version)
		return
	}

	if err := action(*dataFlag, *logLevelFlag, *portFlag); err != nil {
		log.Fatal(err)
	}
}
