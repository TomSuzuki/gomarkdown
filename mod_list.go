package gomarkdown

import (
	"fmt"
	"strings"
)

// listConv ...list generation
func listConv(convData convertedData) convertedData {
	var line = convData.markdownLines[0]
	var text = ""
	var tag = ""
	var nest = 0
	var oldNest = len(convData.listNest)
	var info = []struct {
		html     string
		markdown string
	}{
		{"ul", "- "},
		{"ol", "1. "},
	}

	// list type
	for i := range info {
		if strings.Index(strings.Trim(line, " "), info[i].markdown) == 0 {
			tag = info[i].html
			nest = 1 + strings.Index(line, info[i].markdown)/2
			line = line[strings.Index(line, info[i].markdown)+len(info[i].markdown):]
		}
	}

	// open <ul> or <ol>
	for nest > len(convData.listNest) {
		convData.listNest = append(convData.listNest, tag)
		text = fmt.Sprintf("<%s>", tag)
	}

	// open <li>
	text = fmt.Sprintf("%s<li>%s", text, line)

	// close </ul> or </ol>
	var closeTags = ""
	for nest < len(convData.listNest) {
		closeTags = fmt.Sprintf("</%s>%s", convData.listNest[len(convData.listNest)-1], closeTags)
		if len(convData.listNest) > 0 {
			closeTags = fmt.Sprintf("%s</li>", closeTags)
		}
		convData.listNest = convData.listNest[:len(convData.listNest)-1]
	}
	text = fmt.Sprintf("%s%s", closeTags, text)

	// close </li>
	if oldNest >= nest {
		text = fmt.Sprintf("</li>%s", text)
	}

	// data
	convData.markdownLines[0] = text

	// inline
	convData, _ = inlineConv(convData)

	return convData
}

// listClose ...close list
func listClose(convData convertedData) convertedData {
	var text = ""
	for i := len(convData.listNest) - 1; i >= 0; i-- {
		text = fmt.Sprintf("%s</li></%s>", text, convData.listNest[i])
	}
	convData.markdownLines[0] = fmt.Sprintf("%s%s", text, convData.markdownLines[0])
	convData.listNest = nil

	return convData
}
