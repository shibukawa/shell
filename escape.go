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
	' ':  true,
	'\t': true,
	'\r': true,
	'\n': true,
}

var escapeRequiredCharacters map[rune]string = map[rune]string{
	'"':  `\"`,
	'$':  `\$`,
	'`':  "\\`",
	'\\': `\\`,
	'\t': `\t`,
	'\r': `\r`,
	'\n': `\n`,
}

func Escape(src string) string {
	quoteRequired := false
	var buffer bytes.Buffer
	for _, ch := range src {
		if quoteRequiredCharacters[ch] {
			quoteRequired = true
			replace, ok := escapeRequiredCharacters[ch]
			if ok {
				buffer.WriteString(replace)
			} else {
				buffer.WriteRune(ch)
			}
		} else {
			buffer.WriteRune(ch)
		}
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
			switch ch {
			case 't':
				buffer.WriteByte('\t')
			case 'r':
				buffer.WriteByte('\r')
			case 'n':
				buffer.WriteByte('\n')
			default:
				buffer.WriteRune(ch)
			}
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
