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
	c.Check(Escape("abc def"), Equals, "\"abc def\"")
	c.Check(Escape("abc\rdef"), Equals, "\"abc\\rdef\"")
	c.Check(Escape("abc\ndef"), Equals, "\"abc\\ndef\"")
	c.Check(Escape("abc\tdef"), Equals, "\"abc\\tdef\"")
	c.Check(Escape("abc def[]"), Equals, "\"abc def[]\"")
	// escape required string
	c.Check(Escape("abc def \"$`\\"), Equals, "\"abc def \\\"\\$\\`\\\\\"")
}

func (s *EscapeTest) Test_Unescape(c *C) {
	// regular string
	c.Check(Unescape("abcdef"), Equals, "abcdef")
	// quote required string
	c.Check(Unescape("\"abc def[]\""), Equals, "abc def[]")
	c.Check(Unescape("\"abc\\tdef[]\""), Equals, "abc\tdef[]")
	c.Check(Unescape("\"abc\\rdef[]\""), Equals, "abc\rdef[]")
	c.Check(Unescape("\"abc\\ndef[]\""), Equals, "abc\ndef[]")
	// escape required string
	c.Check(Unescape("\"abc def \\\"\\$\\`\\\\\""), Equals, "abc def \"$`\\")
}
