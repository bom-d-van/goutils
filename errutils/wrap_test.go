package errutil

import (
	"errors"
	"testing"
)

func stringerr() error {
	return Wrap("error burnt out from here")
}

func stringerr2() error {
	return Wrap(stringerr())
}

func stringerr3() error {
	return Wrap(stringerr2())
}

func standarderr() error {
	return Wrap(errors.New("error burnt out from here"))
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
		expect: `[Error Tracks 2006/01/02 15:04:05] error burnt out from here
/Users/bom_d_van/Code/go/workspace/src/github.com/bom-d-van/goutils/errutils/wrap_test.go:9
/Users/bom_d_van/Code/go/workspace/src/github.com/bom-d-van/goutils/errutils/wrap_test.go:13
/Users/bom_d_van/Code/go/workspace/src/github.com/bom-d-van/goutils/errutils/wrap_test.go:17`,
		err: stringerr3,
	},
	{
		name: "standard error",
		expect: `[Error Tracks 2006/01/02 15:04:05] error burnt out from here
/Users/bom_d_van/Code/go/workspace/src/github.com/bom-d-van/goutils/errutils/wrap_test.go:21
/Users/bom_d_van/Code/go/workspace/src/github.com/bom-d-van/goutils/errutils/wrap_test.go:25
/Users/bom_d_van/Code/go/workspace/src/github.com/bom-d-van/goutils/errutils/wrap_test.go:29`,
		err: standarderr3,
	},
}

func TestWrap(t *testing.T) {
	getNow = func() string { return "2006/01/02 15:04:05" }
	var erri interface{}
	for _, tf := range fixtures {
		erri = tf.err()
		err := erri.(Err)
		if err.Details() != tf.expect {
			t.Errorf("Test %s:\nExpect:\n%s\nGot:\n%s", tf.name, tf.expect, err.Details())
		}
	}
}
