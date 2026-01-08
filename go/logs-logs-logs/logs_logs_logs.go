package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var d = map[string]string{
	"U+2757":  "recommendation",
	"U+1F50D": "search",
	"U+2600":  "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, c := range log {
		char := fmt.Sprintf("%U", c)
		if desc, ok := d[char]; ok {
			return desc
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var builder strings.Builder
	for _, c := range log {
		if c == oldRune {
			builder.WriteRune(newRune)
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
	// return strings.ReplaceAll(log, string(oldRune), string(newRune)) // better
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
