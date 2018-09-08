package cmd

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/exec/usecase"
)

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
`, graylog.Version)
}

// Run runs a mock server.
func Run() {
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
		fmt.Println(graylog.Version)
		return
	}

	if err := usecase.Run(*dataFlag, *logLevelFlag, *portFlag); err != nil {
		log.Fatal(err)
	}
}
