package gomarkdown

import "regexp"

// regular expression information
type regList struct {
	regexp   string
	html     string
	isInline bool
}

// regular expression information list
func listRegInfo() []regList {
	return []regList{
		{`\*\*([^\*]*)\*\*`, "<strong>$1</strong>", true},
		{`!\[(.*?)\]\((.*?)\)`, "<img alt='$1' src='$2'>", true},
		{`\[(.*)\]\((.*)\)`, "<a href='$2'>$1</a>", true},
		{`\*([^\*]*)\*|_([^_]*)_|__([^_]*)__`, "<em>$1</em>", true},
		{`^#\s([^#]*?$)`, "<h1>$1</h1>", false},
		{`^##\s([^#]*?$)`, "<h2>$1</h2>", false},
		{`^###\s([^#]*?$)`, "<h3>$1</h3>", false},
		{`^####\s([^#]*?$)`, "<h4>$1</h4>", false},
		{`^#####\s([^#]*?$)`, "<h5>$1</h5>", false},
		{`^######\s([^#]*?$)`, "<h6>$1</h6>", false},
		{`^>\s(.*$)`, "<blockquote><p>$1</p></blockquote>", false},
		{`\s\s$`, "<br>", true},
		{`^(\* ){3,}$|^\*.$|^(- ){3,}|^-{3,}$|^(_ ){3,}$|^_{3,}$`, "<hr>", false},
		{"~([^~]*)~", "<s>$1</s>", true},
		{"`([^`]*)`", "<code>$1</code>", true},
	}
}

// inlineConv ...replacement in line (regular expressions)
func (convData *convertedData) inlineConv() bool {
	var inline = true
	var regexpInfo = listRegInfo()

	for i := range regexpInfo {
		line := convData.markdownLines[0]
		line = regexp.MustCompile(regexpInfo[i].regexp).ReplaceAllString(line, regexpInfo[i].html)
		if convData.markdownLines[0] != line && !regexpInfo[i].isInline {
			inline = false
		}
		convData.markdownLines[0] = line
	}

	return inline
}
