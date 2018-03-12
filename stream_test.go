package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
)

func TestGetStreams(t *testing.T) {
	test.TestGetStreams(t)
}

func TestCreateStream(t *testing.T) {
	test.TestCreateStream(t)
}

func TestGetEnabledStreams(t *testing.T) {
	test.TestGetEnabledStreams(t)
}

func TestGetStream(t *testing.T) {
	test.TestGetStream(t)
}

func TestUpdateStream(t *testing.T) {
	test.TestUpdateStream(t)
}

func TestDeleteStream(t *testing.T) {
	test.TestDeleteStream(t)
}

func TestPauseStream(t *testing.T) {
	test.TestPauseStream(t)
}

func TestResumeStream(t *testing.T) {
	test.TestResumeStream(t)
}
