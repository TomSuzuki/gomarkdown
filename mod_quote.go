package gomarkdown

import (
	"fmt"
	"strings"
)

// quoteConv
func (convData *convertedData) quoteConv() {
	// del >
	n := strings.Index(convData.markdownLines[0], "> ")
	convData.markdownLines[0] = convData.markdownLines[0][n+2:]

	// open
	if convData.typeChenged {
		convData.markdownLines[0] = fmt.Sprintf("<blockquote><p>%s", convData.markdownLines[0])
	}

	// inline
	convData.inlineConv()
}

// quoteClose
func (convData *convertedData) quoteClose() {
	convData.markdownLines[0] = "</p></blockquote>"
}
