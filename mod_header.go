package gomarkdown

import (
	"strconv"
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
	var text []string
	convData.markdownLines[0] = strings.Trim(convData.markdownLines[0], " ")
	h := strings.Count(strings.Split(convData.markdownLines[0], " ")[0], "#")
	text = append(text, "<h")
	text = append(text, strconv.Itoa(h))
	text = append(text, ">")
	text = append(text, convData.markdownLines[0][h+1:])
	text = append(text, "</h")
	text = append(text, strconv.Itoa(h))
	text = append(text, ">")
	convData.markdownLines[0] = strings.Join(text, "")

	// inline
	convData.inlineConv()
}
