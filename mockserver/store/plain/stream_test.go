package plain_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/store"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/plain"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHasStream(t *testing.T) {
	store := plain.NewStore("")
	ok, err := store.HasStream("foo")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("stream foo should not exist")
	}
}

func TestGetStream(t *testing.T) {
	store := plain.NewStore("")
	is, err := store.GetStream("01")
	if err != nil {
		t.Fatal(err)
	}
	if is != nil {
		t.Fatal("stream foo should not exist")
	}
}

func TestGetStreams(t *testing.T) {
	store := plain.NewStore("")
	streams, _, err := store.GetStreams()
	if err != nil {
		t.Fatal(err)
	}
	if len(streams) != 0 {
		t.Fatal("streams should be nil or empty array")
	}
}

func TestAddStream(t *testing.T) {
	st := plain.NewStore("")
	is := testutil.IndexSet("hoge")
	is.ID = store.NewObjectID()
	if err := st.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.ID = store.NewObjectID()
	stream.IndexSetID = is.ID
	if err := st.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	r, err := st.GetStream(stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("is is nil")
	}
}

func TestUpdateStream(t *testing.T) {
	st := plain.NewStore("")
	is := testutil.IndexSet("hoge")
	is.ID = store.NewObjectID()
	if err := st.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream := testutil.Stream()
	stream.ID = store.NewObjectID()
	stream.IndexSetID = is.ID
	if err := st.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	stream.Title += " changed"
	if _, err := st.UpdateStream(stream.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	r, err := st.GetStream(stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	if stream.Title != r.Title {
		t.Fatalf(`stream.Title = "%s", wanted "%s"`, r.Title, stream.Title)
	}
}

func TestDeleteStream(t *testing.T) {
	st := plain.NewStore("")
	stream := testutil.Stream()
	stream.ID = store.NewObjectID()
	if err := st.DeleteStream(stream.ID); err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet("hoge")
	is.ID = store.NewObjectID()
	if err := st.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	stream.IndexSetID = is.ID
	if err := st.AddStream(stream); err != nil {
		t.Fatal(err)
	}
	if err := st.DeleteStream(stream.ID); err != nil {
		t.Fatal(err)
	}
	ok, err := st.HasStream(stream.ID)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("stream should be deleted")
	}
}

func TestGetEnabledStreams(t *testing.T) {
	store := plain.NewStore("")
	streams, _, err := store.GetEnabledStreams()
	if err != nil {
		t.Fatal(err)
	}
	if len(streams) != 0 {
		t.Fatal("streams should be nil or empty array")
	}
}
