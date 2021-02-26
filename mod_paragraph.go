package gomarkdown

import (
	"fmt"
	"strings"
)

// isNone
func (convData *convertedData) isNone() bool {
	return strings.Trim(convData.markdownLines[0], " ") == ""
}

// paragraphConv ...
func (convData *convertedData) paragraphConv() {
	// inline
	convData.inlineConv()

	// open <p>
	if convData.typeChenged {
		convData.markdownLines[0] = fmt.Sprintf("<p>%s", convData.markdownLines[0])
	}
}

// paragraphClose ...
func (convData *convertedData) paragraphClose() {
	convData.shiftLine()
	convData.markdownLines[0] = "</p>"
}
