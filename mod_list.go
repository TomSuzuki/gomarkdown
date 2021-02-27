package gomarkdown

import (
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
	var line = strings.Trim(convData.markdownLines[0], " ")
	for md := range listStyle {
		if strings.Index(line, md) == 0 {
			return true
		}
	}
	return false
}

// convList ...list generation
func (convData *convertedData) convList() {
	var text []string
	var line = convData.markdownLines[0]
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
				text = append(text, "<")
				text = append(text, tag)
				text = append(text, ">")
			}
		}
	}

	// open <li>
	text = append(text, "<li>")
	text = append(text, line)
	convData.markdownLines[0] = strings.Join(text, "")

	// close
	convData.listTagClose(nest, oldNest)

	// inline
	convData.inlineConv()
}

// closeList ...close list
func (convData *convertedData) closeList() {
	convData.shiftLine()
	convData.listTagClose(0, len(convData.listNest))
}

// listTagClose
func (convData *convertedData) listTagClose(nest int, oldNest int) {
	var text = []string{""}

	// close
	for nest < len(convData.listNest) {
		text = append(text, "</li></")
		text = append(text, convData.listNest[len(convData.listNest)-1])
		text = append(text, ">")
		convData.listNest = convData.listNest[:len(convData.listNest)-1]
	}

	// </li>
	if nest <= oldNest && nest != 0 {
		if oldNest > nest {
			text = append(text, "</li>")
		} else {
			text = append(text[:1], text...)
			text[0] = "</li>"
		}
	}

	// append
	text = append(text, convData.markdownLines[0])

	// join
	convData.markdownLines[0] = strings.Join(text, "")
}
