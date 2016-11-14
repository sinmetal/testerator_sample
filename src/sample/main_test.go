package sample

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/favclip/testerator/datastore"
	_ "github.com/favclip/testerator/memcache"
	_ "github.com/favclip/testerator/search"

	"github.com/favclip/testerator"
)

func TestMain(m *testing.M) {
	_, _, err := testerator.SpinUp()

	fmt.Printf("Test Start")

	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	m.Run()

	err = testerator.SpinDown()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
