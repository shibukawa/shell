package shell

import (
	"unicode"
)

type statusCode int

const (
	free statusCode = iota
	inUnquotedWord
	inQuotedWord
	inQuotedWordEscapePrefixed
)

/*
	Splits src string by space. If the chunk is quoted, the spaces between double-quotations are ignored.
*/
func Parse(src string) []string {
	var result []string
	status := free
	currentWordStart := 0
	for i, c := range src {
		switch status {
		case free:
			if !unicode.IsSpace(c) {
				if c == '"' {
					currentWordStart = i
					status = inQuotedWord
				} else {
					currentWordStart = i
					status = inUnquotedWord
				}
			}
		case inUnquotedWord:
			if unicode.IsSpace(c) {
				result = append(result, src[currentWordStart:i])
				status = free
			}
		case inQuotedWord:
			if c == '\\' {
				status = inQuotedWordEscapePrefixed
			} else if c == '"' {
				result = append(result, Unescape(src[currentWordStart:i+1]))
				status = free
			}
		case inQuotedWordEscapePrefixed:
			status = inQuotedWord
		}
	}
	if status != free {
		result = append(result, src[currentWordStart:len(src)])
	}
	return result
}
