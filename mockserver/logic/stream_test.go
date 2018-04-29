package logic_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestAddStream(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet("hoge")
	if _, err := lgc.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := lgc.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	if _, err := lgc.AddStream(nil); err == nil {
		t.Fatal("stream is nil")
	}
}

func TestGetStreams(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}

	if _, _, _, err := lgc.GetStreams(); err != nil {
		t.Fatal(err)
	}
}

func TestGetEnabledStreams(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}

	if _, _, _, err := lgc.GetEnabledStreams(); err != nil {
		t.Fatal(err)
	}
}

func TestGetStream(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}

	is := testutil.IndexSet("hoge")
	if _, err := lgc.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := lgc.AddStream(stream); err != nil {
		t.Fatal(err)
	}

	r, _, err := lgc.GetStream(stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("stream is nil")
	}
	if r.ID != stream.ID {
		t.Fatalf(`stream.ID = "%s", wanted "%s"`, r.ID, stream.ID)
	}
	if _, _, err := lgc.GetStream(""); err == nil {
		t.Fatal("id is required")
	}
	if _, _, err := lgc.GetStream("h"); err == nil {
		t.Fatal("stream should not be found")
	}
}

func TestUpdateStream(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}

	is := testutil.IndexSet("hoge")
	if _, err := lgc.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := lgc.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	if _, _, err := lgc.UpdateStream(stream.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	stream.ID = ""
	if _, _, err := lgc.UpdateStream(stream.NewUpdateParams()); err == nil {
		t.Fatal("stream id is required")
	}
}

func TestDeleteStream(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}

	is := testutil.IndexSet("hoge")
	if _, err := lgc.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := lgc.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	if _, err := lgc.DeleteStream(""); err == nil {
		t.Fatal("stream id is required")
	}
	if _, err := lgc.DeleteStream(stream.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := lgc.DeleteStream(stream.ID); err == nil {
		t.Fatal("already deleted")
	}
}

func TestPauseStream(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}

	is := testutil.IndexSet("hoge")
	if _, err := lgc.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := lgc.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	if _, err := lgc.PauseStream(stream.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := lgc.PauseStream(""); err == nil {
		t.Fatal("stream id is required")
	}
}

func TestResumeStream(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}

	is := testutil.IndexSet("hoge")
	if _, err := lgc.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.IndexSetID = is.ID
	if _, err := lgc.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	if _, err := lgc.ResumeStream(stream.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := lgc.ResumeStream(""); err == nil {
		t.Fatal("stream id is required")
	}
}
