package shell

import (
	. "gopkg.in/check.v1"
)

type ParserTest struct{}

var _ = Suite(&ParserTest{})

func (s *ParserTest) Test_Parser_SplitBySpace(c *C) {
	opts := Parse("abc def")
	c.Check(len(opts), Equals, 2)
	c.Check(opts[0], Equals, "abc")
	c.Check(opts[1], Equals, "def")
}

func (s *ParserTest) Test_Parser_QuotedWords(c *C) {
	opts := Parse(`"hello world" "good morning"`)
	c.Check(len(opts), Equals, 2)
	c.Check(opts[0], Equals, "hello world")
	c.Check(opts[1], Equals, "good morning")
}

func (s *ParserTest) Test_Parser_QuotedWords_WithEscape(c *C) {
	opts := Parse("\"hello \\\"world\\\"\" \"good morning\"")
	c.Check(len(opts), Equals, 2)
	c.Check(opts[0], Equals, "hello \"world\"")
	c.Check(opts[1], Equals, "good morning")
}
