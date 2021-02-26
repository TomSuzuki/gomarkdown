package gomarkdown

import "regexp"

// inline / regular expression information list {markdonw, html} !attention to the priority
var listRegInfo = [][2]string{
	{`\*\*([^\*]*)\*\*`, "<strong>$1</strong>"},
	{`!\[(.*?)\]\((.*?)\)`, "<img alt='$1' src='$2'>"},
	{`\[(.*)\]\((.*)\)`, "<a href='$2'>$1</a>"},
	{`\*([^\*]*)\*|_([^_]*)_|__([^_]*)__`, "<em>$1</em>"},
	{`\s\s$`, "<br>"},
	{"~([^~]*)~", "<s>$1</s>"},
	{"`([^`]*)`", "<code>$1</code>"},
}

// inlineConv ...replacement in line (regular expressions)
func (convData *convertedData) inlineConv() {
	for _, v := range listRegInfo {
		var md = v[0]
		var html = v[1]
		line := convData.markdownLines[0]
		line = regexp.MustCompile(md).ReplaceAllString(line, html)
		convData.markdownLines[0] = line
	}
}
