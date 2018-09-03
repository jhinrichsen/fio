package fio

import "testing"

func TestExists(t *testing.T) {
	want := true
	got, err := Exists("testdata/a-file")
	if err != nil {
		t.Fatal(err)
	}
	if want == got {
		t.Fatalf("want %t but got %t", want, got)
	}
}

func TestExistsNonExistent(t *testing.T) {
	want := false
	got, err := Exists("testdata/this-file-does-not-exist")
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %t but got %t", want, got)
	}
}
