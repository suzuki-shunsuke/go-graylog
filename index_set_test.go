package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/v9"
	"github.com/suzuki-shunsuke/go-graylog/v9/testdata"
)

func TestIndexSetNewUpdateParams(t *testing.T) {
	is := testdata.IndexSet()
	prms := is.NewUpdateParams()
	if is.Title != prms.Title {
		t.Fatalf(`prms.Title = "%s", wanted "%s"`, prms.Title, is.Title)
	}
}

func TestSetCreateDefaultValues(t *testing.T) {
	is := &graylog.IndexSet{}
	is.SetCreateDefaultValues()
	if is.CreationDate == "" {
		t.Fatal("is.CreationDate must be set")
	}
	if is.Shards == 0 {
		t.Fatal("is.Shards must be set")
	}
	if is.IndexAnalyzer != "standard" {
		t.Fatalf(`is.IndexAnalyzer = "%s", wanted "standard"`, is.IndexAnalyzer)
	}
}

func TestCreationTime(t *testing.T) {
	is := &graylog.IndexSet{}
	is.SetCreateDefaultValues()
	if _, err := is.CreationTime(); err != nil {
		t.Fatal(err)
	}
}

func TestNewMessageCountRotationStrategy(t *testing.T) {
	s := graylog.NewMessageCountRotationStrategy(0)
	if s.MaxDocsPerIndex == 0 {
		t.Fatal("s.MaxDocsPerIndex must not be 0")
	}
	s = graylog.NewMessageCountRotationStrategy(10)
	if s.MaxDocsPerIndex != 10 {
		t.Fatalf("s.MaxDocsPerIndex = %d, wanted 10", s.MaxDocsPerIndex)
	}
}

func TestNewSizeBasedRotationStrategy(t *testing.T) {
	s := graylog.NewSizeBasedRotationStrategy(0)
	if s.MaxSize == 0 {
		t.Fatal("s.MaxSize must not be 0")
	}
	s = graylog.NewSizeBasedRotationStrategy(10)
	if s.MaxSize != 10 {
		t.Fatalf("s.MaxSize = %d, wanted 10", s.MaxSize)
	}
}

func TestNewTimeBasedRotationStrategy(t *testing.T) {
	s := graylog.NewTimeBasedRotationStrategy("")
	if s.RotationPeriod == "" {
		t.Fatal(`s.RotationPeriod must not be ""`)
	}
	s = graylog.NewTimeBasedRotationStrategy("a")
	if s.RotationPeriod != "a" {
		t.Fatalf(`s.RotationPeriod = "%s", wanted "a"`, s.RotationPeriod)
	}
}

func TestNewDeletionRetentionStrategy(t *testing.T) {
	s := graylog.NewDeletionRetentionStrategy(0)
	if s.MaxNumberOfIndices == 0 {
		t.Fatal("s.MaxNumberOfIndices must not be 0")
	}
	s = graylog.NewDeletionRetentionStrategy(10)
	if s.MaxNumberOfIndices != 10 {
		t.Fatalf(`s.MaxNumberOfIndices = %d, wanted 10`, s.MaxNumberOfIndices)
	}
}

func TestNewClosingRetentionStrategy(t *testing.T) {
	s := graylog.NewClosingRetentionStrategy(0)
	if s.MaxNumberOfIndices == 0 {
		t.Fatal("s.MaxNumberOfIndices must not be 0")
	}
	s = graylog.NewClosingRetentionStrategy(10)
	if s.MaxNumberOfIndices != 10 {
		t.Fatalf(`s.MaxNumberOfIndices = %d, wanted 10`, s.MaxNumberOfIndices)
	}
}

func TestNewNoopRetentionStrategy(t *testing.T) {
	s := graylog.NewNoopRetentionStrategy(0)
	if s.MaxNumberOfIndices == 0 {
		t.Fatal("s.MaxNumberOfIndices must not be 0")
	}
	s = graylog.NewNoopRetentionStrategy(10)
	if s.MaxNumberOfIndices != 10 {
		t.Fatalf(`s.MaxNumberOfIndices = %d, wanted 10`, s.MaxNumberOfIndices)
	}
}
