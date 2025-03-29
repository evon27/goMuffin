package utils

import (
	"strconv"
	"time"
)

const (
	ShortTime = "t"
	LongTime  = "T"

	ShortDate = "d"
	LongDate  = "D"

	ShortDateTime = "f"
	LongDateTime  = "F"

	RelativeTime = "R"
)

func InlineCode(str string) string {
	return "`" + str + "`"
}

func CodeBlockWithLanguage(language string, content string) string {
	return "```" + language + "\n" + content + "\n" + "```"
}

func CodeBlock(content string) string {
	return "```\n" + content + "\n" + "```"
}

func Time(time *time.Time) string {
	return "<t:" + strconv.FormatInt(time.Unix(), 10) + ">"
}

func TimeWithStyle(time *time.Time, style string) string {
	return "<t:" + strconv.FormatInt(time.Unix(), 10) + ":" + style + ">"
}
