package optstring_parser

import (
	"unicode"
)

type statusCode int

const (
	free statusCode = iota
	inUnquotedWord
	inQuotedWord
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
					currentWordStart = i + 1
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
			if c == '"' {
				result = append(result, src[currentWordStart:i])
				status = free
			}
		}
	}
	if status != free {
		result = append(result, src[currentWordStart:len(src)])
	}
	return result
}