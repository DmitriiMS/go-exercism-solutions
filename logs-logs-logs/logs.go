package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Handy map for prefixes
var pref = map[string]string{
	"error":   "[ERROR]:",
	"warning": "[WARNING]:",
	"info":    "[INFO]:",
}

func determineLogLevel(line string) string {
	switch {
	case strings.HasPrefix(line, pref["error"]):
		return "error"
	case strings.HasPrefix(line, pref["warning"]):
		return "warning"
	case strings.HasPrefix(line, pref["info"]):
		return "info"
	default:
		return ""
	}
}

func removePrefix(line string) string {
	prefixless := strings.TrimPrefix(line, pref[determineLogLevel(line)])
	return strings.TrimSpace(prefixless)
}

// Message extracts the message from the provided log line.
func Message(line string) string {
	return removePrefix(line)
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	return utf8.RuneCountInString(removePrefix(line))
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	return determineLogLevel(line)
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	return fmt.Sprintf("%s (%s)", removePrefix(line), determineLogLevel(line))
}
