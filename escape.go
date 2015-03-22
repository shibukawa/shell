package shell

import (
	"bytes"
	"fmt"
)

var quoteRequiredCharacters map[rune]bool = map[rune]bool{
	'"':  true,
	'$':  true,
	'@':  true,
	'&':  true,
	'\'': true,
	'(':  true,
	')':  true,
	'^':  true,
	'|':  true,
	'[':  true,
	']':  true,
	'{':  true,
	'}':  true,
	';':  true,
	'*':  true,
	'?':  true,
	'<':  true,
	'>':  true,
	'`':  true,
	'\\': true,
}

var escapeRequiredCharacters map[rune]bool = map[rune]bool{
	'"':  true,
	'$':  true,
	'`':  true,
	'\\': true,
}

func Escape(src string) string {
	quoteRequired := false
	var buffer bytes.Buffer
	for _, ch := range src {
		if quoteRequiredCharacters[ch] {
			quoteRequired = true
			if escapeRequiredCharacters[ch] {
				buffer.WriteRune('\\')
			}
		}
		buffer.WriteRune(ch)
	}
	if quoteRequired {
		return fmt.Sprintf(`"%s"`, buffer.String())
	}
	return buffer.String()
}

func Unescape(src string) string {
	if len(src) == 0 || src[0] != '"' {
		return src
	}
	var buffer bytes.Buffer
	bound := len(src) - 1
	escaped := false
	for i, ch := range src[1:bound] {
		// skip
		if escaped {
			buffer.WriteRune(ch)
			escaped = false
		} else if ch == '\\' {
			if i < bound-1 {
				escaped = true
			}
		} else {
			buffer.WriteRune(ch)
		}
	}
	return buffer.String()
}
