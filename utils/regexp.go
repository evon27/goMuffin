package utils

import "regexp"

var ExtractQuotedText *regexp.Regexp = regexp.MustCompile("[\"'`](.*?)[\"'`]")
