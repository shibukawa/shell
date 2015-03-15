package optstring_parser

import (
	. "gopkg.in/check.v1"
)

type OptStringParserTest struct {}

var _ = Suite(&OptStringParserTest{})

func (s *OptStringParserTest) Test_OptStringParser_SplitBySpace(c *C) {
	opts := Parse("abc def")
	c.Check(len(opts), Equals, 2)
	c.Check(opts[0], Equals, "abc")
	c.Check(opts[1], Equals, "def")
}

func (s *OptStringParserTest) Test_OptStringParser_QuotedWords(c *C) {
	opts := Parse(`"hello world" "good morning"`)
	c.Check(len(opts), Equals, 2)
	c.Check(opts[0], Equals, "hello world")
	c.Check(opts[1], Equals, "good morning")
}