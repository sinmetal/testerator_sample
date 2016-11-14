package sample

import (
	"testing"

	"github.com/favclip/testerator"

	"google.golang.org/appengine/datastore"
)

func TestPut(t *testing.T) {
	_, c, err := testerator.SpinUp()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer testerator.SpinDown()

	const title = "sample"

	stored, err := Put(c, title)
	if err != nil {
		t.Fatalf("Put error. err = %v", err)
	}
	if stored.Title != title {
		t.Fatalf("unexpected Item.Title. v = %s", stored.Title)
	}
}

func TestGetEmpty(t *testing.T) {
	_, c, err := testerator.SpinUp()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer testerator.SpinDown()

	key := datastore.NewKey(c, "Item", "dummy_key", 0, nil)

	_, err = Get(c, key)
	if err != datastore.ErrNoSuchEntity {
		t.Fatalf("unexpected err. err = %v", err)
	}
}

func TestGetEntity(t *testing.T) {
	_, c, err := testerator.SpinUp()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer testerator.SpinDown()

	const title = "sample"

	entity, err := Put(c, title)
	if err != nil {
		t.Fatalf("Put error. err = %v", err)
	}

	stored, err := Get(c, entity.Key)
	if err != nil {
		t.Fatalf("Get error. err = %v", err)
	}
	if stored.Title != title {
		t.Fatalf("unexpected Item.Title. v = %s", stored.Title)
	}
}
