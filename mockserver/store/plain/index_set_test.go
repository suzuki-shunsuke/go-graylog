package plain_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/store"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/plain"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHasIndexSet(t *testing.T) {
	store := plain.NewStore("")
	ok, err := store.HasIndexSet("foo")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("is foo should not exist")
	}
}

func TestGetIndexSet(t *testing.T) {
	store := plain.NewStore("")
	is, err := store.GetIndexSet("foo")
	if err != nil {
		t.Fatal(err)
	}
	if is != nil {
		t.Fatal("is foo should not exist")
	}
}

func TestGetIndexSets(t *testing.T) {
	store := plain.NewStore("")
	iss, _, err := store.GetIndexSets(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if iss != nil && len(iss) != 0 {
		t.Fatal("iss should be nil or empty array")
	}
}

func TestAddIndexSet(t *testing.T) {
	st := plain.NewStore("")
	is := testutil.IndexSet("hoge")
	is.ID = store.NewObjectID()
	if err := st.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	r, err := st.GetIndexSet(is.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("is is nil")
	}
}

func TestUpdateIndexSet(t *testing.T) {
	st := plain.NewStore("")
	is := testutil.IndexSet("hoge")
	is.ID = store.NewObjectID()
	if err := st.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	r, err := st.GetIndexSet(is.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("is is nil")
	}
	is.Title += " changed"
	if _, err := st.UpdateIndexSet(is.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	r, err = st.GetIndexSet(is.ID)
	if err != nil {
		t.Fatal(err)
	}
	if is.Title != r.Title {
		t.Fatalf(`is.Title = "%s", wanted "%s"`, r.Title, is.Title)
	}
}

func TestDeleteIndexSet(t *testing.T) {
	st := plain.NewStore("")
	if err := st.DeleteIndexSet("foo"); err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet("hoge")
	is.ID = store.NewObjectID()
	if err := st.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	if err := st.DeleteIndexSet(is.ID); err != nil {
		t.Fatal(err)
	}
	ok, err := st.HasIndexSet(is.ID)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("is should be deleted")
	}
}

func TestGetDefaultIndexSetID(t *testing.T) {
	store := plain.NewStore("")
	if _, err := store.GetDefaultIndexSetID(); err != nil {
		t.Fatal(err)
	}
}

func TestSetDefaultIndexSetID(t *testing.T) {
	st := plain.NewStore("")
	is := testutil.IndexSet("hoge")
	is.ID = store.NewObjectID()
	if err := st.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	if err := st.SetDefaultIndexSetID(is.ID); err != nil {
		t.Fatal(err)
	}
	id, err := st.GetDefaultIndexSetID()
	if err != nil {
		t.Fatal(err)
	}
	if id != is.ID {
		t.Fatalf("default id is <%s>, wanted <%s>", id, is.ID)
	}
}

func TestIsConflictIndexPrefix(t *testing.T) {
	st := plain.NewStore("")
	is := testutil.IndexSet("hoge")
	is.ID = store.NewObjectID()
	if err := st.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	ok, err := st.IsConflictIndexPrefix(is.ID, is.IndexPrefix)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("id should not conflict")
	}
	ok, err = st.IsConflictIndexPrefix("", is.IndexPrefix)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("id should conflict")
	}
	ok, err = st.IsConflictIndexPrefix("", "ho")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("id should conflict")
	}
	ok, err = st.IsConflictIndexPrefix("", "hogefuga")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("id should conflict")
	}
}
