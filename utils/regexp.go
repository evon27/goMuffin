package utils

import "regexp"

var FlexibleStringParser *regexp.Regexp = regexp.MustCompile("[^\\s\"'「」«»]+|\"([^\"]*)\"|'([^']*)'|「([^」]*)」|«([^»]*)»")
var Decimals *regexp.Regexp = regexp.MustCompile(`\d+`)
