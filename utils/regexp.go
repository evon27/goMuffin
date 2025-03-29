package utils

import "regexp"

var ExtractQuotedText *regexp.Regexp = regexp.MustCompile("[\"'`](.*?)[\"'`]")
var Decimals *regexp.Regexp = regexp.MustCompile(`\d+`)
