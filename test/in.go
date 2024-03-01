package test

import (
	av1 "github.com/ashishmalgawa/mocker/test/a"
	bv1 "github.com/ashishmalgawa/mocker/test/b"
	"github.com/ashishmalgawa/mocker/test/c"
)

type Iface interface {
	One(str string, variadic ...string) (string, []string)
	Two(int, int) int
	Three(av1.Int) bv1.Str
	Four(c.Int)
}
