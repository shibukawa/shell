package shell

import (
	. "gopkg.in/check.v1"
)

type EscapeTest struct{}

var _ = Suite(&EscapeTest{})

func (s *EscapeTest) Test_Escape(c *C) {
	// regular string
	c.Check(Escape("abcdef"), Equals, "abcdef")
	// quote required string
	c.Check(Escape("abc def[]"), Equals, "\"abc def[]\"")
	// escape required string
	c.Check(Escape("abc def \"$`\\"), Equals, "\"abc def \\\"\\$\\`\\\\\"")
}

func (s *EscapeTest) Test_Unescape(c *C) {
	// regular string
	c.Check(Unescape("abcdef"), Equals, "abcdef")
	// quote required string
	c.Check(Unescape("\"abc def[]\""), Equals, "abc def[]")
	// escape required string
	c.Check(Unescape("\"abc def \\\"\\$\\`\\\\\""), Equals, "abc def \"$`\\")
}
