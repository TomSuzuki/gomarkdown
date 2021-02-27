package gomarkdown

import (
	"fmt"
	"strings"
)

// listStyle
var listStyle = map[string]string{
	"- ":  "ul",
	"* ":  "ul",
	"1. ": "ol",
	"+ ":  "ul",
}

// isList
func (convData *convertedData) isList() bool {
	var line = convData.markdownLines[0]
	for md := range listStyle {
		if strings.Index(strings.Trim(line, " "), md) == 0 {
			return true
		}
	}
	return false
}

// listConv ...list generation
func (convData *convertedData) listConv() {
	var line = convData.markdownLines[0]
	var openTags = ""
	var nest = 0
	var oldNest = len(convData.listNest)

	// list type and open list
	for md, tag := range listStyle {
		if strings.Index(strings.Trim(line, " "), md) == 0 {
			nest = 1 + strings.Index(line, md)/2
			line = line[strings.Index(line, md)+len(md):]

			// open <ul> or <ol>
			for nest > len(convData.listNest) {
				convData.listNest = append(convData.listNest, tag)
				openTags = fmt.Sprintf("<%s>", tag)
			}
		}
	}

	// open <li>
	convData.markdownLines[0] = fmt.Sprintf("%s<li>%s", openTags, line)

	// close
	convData.listTagClose(nest, oldNest)

	// inline
	convData.inlineConv()
}

// listClose ...close list
func (convData *convertedData) listClose() {
	convData.shiftLine()
	convData.listTagClose(0, len(convData.listNest))
}

// listTagClose
func (convData *convertedData) listTagClose(nest int, oldNest int) {
	var tags = ""
	var tagl = "" // Add </li> that could not be added to the list when the nesting is finished.

	// close
	for nest < len(convData.listNest) {
		tags = fmt.Sprintf("%s</li></%s>", tags, convData.listNest[len(convData.listNest)-1])
		convData.listNest = convData.listNest[:len(convData.listNest)-1]
	}

	// </li>
	if nest <= oldNest && nest != 0 {
		tagl = "</li>"
	}

	// append
	if oldNest > nest {
		convData.markdownLines[0] = fmt.Sprintf("%s%s%s", tags, tagl, convData.markdownLines[0])
	} else {
		convData.markdownLines[0] = fmt.Sprintf("%s%s%s", tagl, convData.markdownLines[0], tags)
	}
}
