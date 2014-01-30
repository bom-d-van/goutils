package gocheckutil

import (
	"regexp"
	"testing"
	. "launchpad.net/gocheck"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type GoCheckutilSuite struct{}

var _ = Suite(&GoCheckutilSuite{})

func (s *GoCheckutilSuite) TestRegexpMatches(c *C) {
	c.Check("42", RegexpMatches, regexp.MustCompile("42"))
}
