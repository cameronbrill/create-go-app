package cli

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

func TestClone(t *testing.T) {
	var a app
	a.name = "test-app"
	a.template = "main"
	err := a.clone()
	dir := a.directory
	defer func() {
		if dir == "" {
			panic("tried to remove active directory")
		}
		err := os.RemoveAll(fmt.Sprintf("./%s", dir))
		if err != nil {
			t.Error(err)
		}
	}()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(fmt.Sprintf("./%s", a.name)); err != nil && !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	err = a.clone()
	defer func() {
		if a.directory == "" {
			panic("tried to remove active directory")
		}
		err := os.RemoveAll(fmt.Sprintf("./%s-1", dir))
		if err != nil {
			t.Error(err)
		}
	}()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(fmt.Sprintf("./%s-1", a.name)); err != nil && !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
}
