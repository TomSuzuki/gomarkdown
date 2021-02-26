package gomarkdown

import (
	"fmt"
	"strings"
)

// isHeader ...markdownLines[0] is header?
func (convData *convertedData) isHeader() bool {
	line := strings.Split(strings.Trim(convData.markdownLines[0], " "), " ")[0]
	return line != "" && strings.Trim(line, "#") == ""
}

// headerConv ...
func (convData *convertedData) headerConv() {
	// <h1> - <h6>
	convData.markdownLines[0] = strings.Trim(convData.markdownLines[0], " ")
	head := strings.Split(convData.markdownLines[0], " ")[0]
	h := strings.Count(head, "#")
	text := convData.markdownLines[0][h+1:]
	convData.markdownLines[0] = fmt.Sprintf("<h%d>%s</h%d>", h, text, h)

	// inline
	convData.inlineConv()
}
