package gomarkdown

import (
	"strings"
)

// isNone
func (convData *convertedData) isNone() bool {
	return strings.Trim(convData.markdownLines[0], " ") == ""
}

// convParagraph ...
func (convData *convertedData) convParagraph() {
	// inline
	convData.inlineConv()

	// open <p>
	if convData.typeChenged {
		convData.markdownLines[0] = ("<p>" + convData.markdownLines[0])
	}
}

// closeParagraph ...
func (convData *convertedData) closeParagraph() {
	convData.shiftLine()
	convData.markdownLines[0] = "</p>"
}
