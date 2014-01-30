package gocheckutil

import (
	"regexp"

	"launchpad.net/gocheck"
)

type RegexpMatchesChecker struct {
	checkerInfo *gocheck.CheckerInfo
}

var RegexpMatches = RegexpMatchesChecker{&gocheck.CheckerInfo{
	Name:   "RegexpMatches",
	Params: []string{"sample", "regexp"},
}}

func (r RegexpMatchesChecker) Info() *gocheck.CheckerInfo {
	return r.checkerInfo
}

func (r RegexpMatchesChecker) Check(params []interface{}, names []string) (result bool, error string) {
	sample, ok := params[0].(string)
	if !ok {
		return false, "Sample must be a string"
	}
	reg, ok := params[1].(regexp.Regexp)
	if !ok {
		reg, ok := params[1].(*regexp.Regexp)
		if !ok {
			return false, "Regexp must be a regexp.Regexp"
		}

		return reg.MatchString(sample), ""
	}

	return reg.MatchString(sample), ""
}
