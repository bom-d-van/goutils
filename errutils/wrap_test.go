package errutils

import (
	"errors"
	"go/build"
	"strings"
	"testing"
)

func stringerr() error {
	return Wrap("error break out")
}

func stringerr2() error {
	return Wrap(stringerr())
}

func stringerr3() error {
	return Wrap(stringerr2())
}

func standarderr() error {
	return Wrap(errors.New("error break out"))
}

func standarderr2() error {
	return Wrap(standarderr())
}

func standarderr3() error {
	return Wrap(standarderr2())
}

type fixture struct {
	name   string
	expect string
	err    func() error
}

var fixtures = []fixture{
	{
		name: "string error",
		expect: `[Error Tracks 2006/01/02 15:04:05] error break out
/src/github.com/bom-d-van/goutils/errutils/wrap_test.go:11
/src/github.com/bom-d-van/goutils/errutils/wrap_test.go:15
/src/github.com/bom-d-van/goutils/errutils/wrap_test.go:19`,
		err: stringerr3,
	},
	{
		name: "standard error",
		expect: `[Error Tracks 2006/01/02 15:04:05] error break out
/src/github.com/bom-d-van/goutils/errutils/wrap_test.go:23
/src/github.com/bom-d-van/goutils/errutils/wrap_test.go:27
/src/github.com/bom-d-van/goutils/errutils/wrap_test.go:31`,
		err: standarderr3,
	},
}

func TestWrap(t *testing.T) {
	getNow = func() string { return "2006/01/02 15:04:05" }
	var erri interface{}
	for _, tf := range fixtures {
		erri = tf.err()
		err := erri.(Error)
		details := err.Details()
		details = strings.Replace(details, build.Default.GOPATH, "", -1)
		if details != tf.expect {
			t.Errorf("Test %s:\nExpect:\n%s\nGot:\n%s", tf.name, tf.expect, details)
		}
	}
}
