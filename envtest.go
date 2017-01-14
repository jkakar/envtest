// Package envtest lets you safely modify the environment in tests.
//
// Setup the test, defer the cleanup function and then do whatever you want to
// the environment. Added, modified and deleted environment variables will be
// returned to their original state at the end of the test.
//
//     func TestSomething(t *testing.T) {
//         teardown := envtest.Setup()
//         defer teardown()
//         os.Setenv("EDITOR", "/usr/bin/sed")
//         os.Unsetenv("PATH")
//     }
package envtest

import (
	"os"
	"strings"
)

// Setup captures the current state of the environment and returns a cleanup
// function that can be used to restore it.
func Setup() func() {
	orig := mapenv()
	return func() {
		restore(orig)
	}
}

func mapenv() map[string]string {
	var env = make(map[string]string)
	for _, v := range os.Environ() {
		kv := strings.SplitN(v, "=", 2)
		env[kv[0]] = kv[1]
	}
	return env
}

func restore(orig map[string]string) {
	removeAddedVars(orig)
	restoreModifiedVars(orig)
	restoreDeletedVars(orig)
}

func removeAddedVars(orig map[string]string) {
	for k, _ := range mapenv() {
		if _, ok := orig[k]; !ok {
			os.Unsetenv(k)
		}
	}
}

func restoreModifiedVars(orig map[string]string) {
	for k, vn := range mapenv() {
		if v := orig[k]; v != vn {
			os.Setenv(k, v)
		}
	}
}

func restoreDeletedVars(orig map[string]string) {
	cur := mapenv()
	for k, v := range orig {
		if _, ok := cur[k]; !ok {
			os.Setenv(k, v)
		}
	}
}
