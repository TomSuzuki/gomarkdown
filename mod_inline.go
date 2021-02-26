package gomarkdown

import (
	"fmt"
	"regexp"
	"strings"
)

// inline / regular expression information list {markdonw, html} !attention to the priority
var listRegInfo = [][2]string{
	{`!\[(.*?)\]\((.*?)\)`, "<img alt='$1' src='$2'>"},
	{`\[(.*)\]\((.*)\)`, "<a href='$2'>$1</a>"},
}

// inlineConv ...replacement in line (regular expressions)
func (convData *convertedData) inlineConv() {

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

	// <img> and <a>
	for _, v := range listRegInfo {
		var md = v[0]
		var html = v[1]
		line := convData.markdownLines[0]
		line = regexp.MustCompile(md).ReplaceAllString(line, html)
		convData.markdownLines[0] = line
	}

	// <br>
	convData.markdownLines[0] = strings.Replace(convData.markdownLines[0], "  ", "<br>", -1)
}

// inlineTage ...md -> html
func (convData *convertedData) inlineTage(md string, html string) {
	var codeList = strings.Split(convData.markdownLines[0], md)
	convData.markdownLines[0] = ""
	var isEven = len(codeList)%2 == 0

	// insert tags
	for i, v := range codeList {
		if isEven && i == len(codeList)-1 {
			convData.markdownLines[0] += (md + v)
		} else if i%2 == 0 {
			convData.markdownLines[0] += v
		} else if isNotBrokenHTML(v) {
			convData.markdownLines[0] += fmt.Sprintf("<%s>%s</%s>", html, v, html)
		} else {
			convData.markdownLines[0] += (md + v)
		}
	}
}

// isNotBrokenHTML ..."<s></s><em></em>" <<< true, "<s><em></em>" <<< false, "<img ...>" <<< false...?
func isNotBrokenHTML(html string) bool {
	var nest = 0
	var open = false
	for _, s := range []byte(html) {
		if open {
			if string(s) == "/" {
				nest--
			} else {
				nest++
			}
		}
		open = string(s) == "<"
		if nest == -1 {
			return false
		}
	}
	return nest == 0
}
