package gomarkdown

import (
	"fmt"
	"strings"
)

// listConv ...list generation
func (convData *convertedData) listConv() {
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
	convData.markdownLines[0] = fmt.Sprintf("%s<li>%s", text, line)

	// close
	convData.listTagClose(nest, oldNest, true)

	// inline
	convData.inlineConv()
}

// listClose ...close list
func (convData *convertedData) listClose() {
	convData.listTagClose(0, len(convData.listNest), false)
}

// listTagClose
func (convData *convertedData) listTagClose(nest int, oldNest int, inlist bool) {
	var tags = ""

	// close
	for nest < len(convData.listNest) {
		tags = fmt.Sprintf("%s</li></%s>", tags, convData.listNest[len(convData.listNest)-1])
		convData.listNest = convData.listNest[:len(convData.listNest)-1]
	}

	//
	if nest < oldNest && nest != 0 {
		tags = tags + "</li>"
	}

	// append
	if !inlist || oldNest > nest {
		convData.markdownLines[0] = fmt.Sprintf("%s%s", tags, convData.markdownLines[0])
	} else {
		convData.markdownLines[0] = fmt.Sprintf("%s%s", convData.markdownLines[0], tags)
		if oldNest >= nest && len(convData.listNest) > 0 {
			convData.markdownLines[0] = fmt.Sprintf("</li>%s", convData.markdownLines[0])
		}
	}

}
