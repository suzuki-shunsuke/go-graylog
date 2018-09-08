package usecase

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog/mockserver"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/plain"
)

// Run runs a mock server.
func Run(dataPath, logLevel string, port int) error {
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
