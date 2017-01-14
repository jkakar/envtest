package envtest

import (
	"os"
	"reflect"
	"testing"
)

func TestSetupWithoutChanges(t *testing.T) {
	expected := os.Environ()
	teardown := Setup()
	teardown()
	env := os.Environ()
	if !reflect.DeepEqual(expected, env) {
		t.Fatalf("expected %v but got %v", expected, env)
	}
}

func TestSetupWithAddedVar(t *testing.T) {
	teardown := Setup()
	os.Setenv("FOOBARBAZ", "added")
	teardown()
	actual := os.Getenv("FOOBARBAZ")
	if actual != "" {
		t.Fatalf("expected \"\" but got %s", actual)
	}
}

func TestSetupWithModifiedVar(t *testing.T) {
	os.Setenv("FOOBARBAZ", "original")
	defer os.Unsetenv("FOOBARBAZ")
	teardown := Setup()
	os.Setenv("FOOBARBAZ", "modified")
	teardown()
	actual := os.Getenv("FOOBARBAZ")
	if actual != "original" {
		t.Fatalf("expected original but got %s", actual)
	}
}

func TestSetupWithDeletedVar(t *testing.T) {
	os.Setenv("FOOBARBAZ", "original")
	defer os.Unsetenv("FOOBARBAZ")
	teardown := Setup()
	os.Unsetenv("FOOBARBAZ")
	teardown()
	actual := os.Getenv("FOOBARBAZ")
	if actual != "original" {
		t.Fatalf("expected original but got %s", actual)
	}
}
