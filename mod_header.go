package gomarkdown

import (
	"strings"
)

// isHeader ...markdownLines[0] is header?
func (convData *convertedData) isHeader() bool {
	line := strings.Split(strings.Trim(convData.markdownLines[0], " "), " ")[0]
	return line != "" && strings.Trim(line, "#") == ""
}

const headText = "123456"

// convHeader ...
func (convData *convertedData) convHeader() {
	// <h1> - <h6>
	var text []string
	convData.markdownLines[0] = strings.Trim(convData.markdownLines[0], " ")
	h := strings.Count(strings.Split(convData.markdownLines[0], " ")[0], "#")
	if h <= 6 && h >= 1 {
		text = append(text, "<h")
		text = append(text, headText[h-1:h])
		text = append(text, ">")
		text = append(text, convData.markdownLines[0][h+1:])
		text = append(text, "</h")
		text = append(text, headText[h-1:h])
		text = append(text, ">")
		convData.markdownLines[0] = strings.Join(text, "")
	}

	// inline
	convData.inlineConv()
}
