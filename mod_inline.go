package gomarkdown

import (
	"fmt"
	"regexp"
	"strings"
)

// inline / regular expression information list {markdonw, html} !attention to the priority
var listRegInfo = [][2]string{
	//{`\*\*([^\*]*)\*\*`, "<strong>$1</strong>"},
	{`!\[(.*?)\]\((.*?)\)`, "<img alt='$1' src='$2'>"},
	{`\[(.*)\]\((.*)\)`, "<a href='$2'>$1</a>"},
	//{`\*([^\*]*)\*|_([^_]*)_|__([^_]*)__`, "<em>$1</em>"},
	//{"~([^~]*)~", "<s>$1</s>"},
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

	// inline text
	for _, v := range [][]string{
		{"**", "strong"},
		{"`", "code"},
		{"~", "s"},
		{"__", "em"},
		{"_", "em"},
		{"*", "em"},
	} {
		convData.inlineTage(v[0], v[1])
	}

	// <br>
	convData.markdownLines[0] = strings.Replace(convData.markdownLines[0], "  ", "<br>", -1)
}

// [! Needs correction !] inlineTage ...md -> html / "* text - text * text - " <- ????
func (convData *convertedData) inlineTage(md string, html string) {
	var codeList = strings.Split(convData.markdownLines[0], md)
	var lastText = ""

	convData.markdownLines[0] = codeList[0]
	codeList = codeList[1:]
	if len(codeList)%2 != 0 {
		lastText = md + codeList[len(codeList)-1]
		codeList = codeList[:len(codeList)-1]
	}

	// insert tags
	for i, v := range codeList {
		if i%2 == 0 {
			convData.markdownLines[0] += fmt.Sprintf("<%s>%s", html, v)
		} else {
			convData.markdownLines[0] += fmt.Sprintf("</%s>%s", html, v)
		}
	}

	convData.markdownLines[0] += lastText
}
