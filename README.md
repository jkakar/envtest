Support for manipulating environment variables in tests.

## Installation

```bash
go get -u github.com/jkakar/envtest
```

## Usage

Setup the test, defer the cleanup function and then do whatever you want to
the environment. Added, modified and deleted environment variables will be
returned to their original state at the end of the test.

```go
func TestSomething(t *testing.T) {
    teardown := envtest.Setup()
    defer teardown()
    os.Setenv("EDITOR", "/usr/bin/sed")
    os.Unsetenv("PATH")
}
```
