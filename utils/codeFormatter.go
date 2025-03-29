package utils

func InlineCode(str string) string {
	return "`" + str + "`"
}

func CodeBlockWithLanguage(language string, content string) string {
	return "```" + language + "\n" + content + "\n" + "```"
}

func CodeBlock(content string) string {
	return "```\n" + content + "\n" + "```"
}
