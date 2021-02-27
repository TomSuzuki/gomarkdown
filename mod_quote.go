package gomarkdown

import (
	"fmt"
	"strings"
)

// isQuote
func (convData *convertedData) isQuote() bool {
	return (strings.Trim(convData.markdownLines[0], " ") + "  ")[:2] == "> " || (convData.lineType == typeQuote && !convData.isNone())
}

// quoteConv
func (convData *convertedData) quoteConv() {
	// >
	nest := strings.Count(convData.markdownLines[0], "> ")
	if nest == 0 {
		nest = convData.nestQuote
	} else {
		convData.markdownLines[0] = convData.markdownLines[0][2*nest:]
	}

	// open
	if convData.nestQuote < nest {
		var oldNest = convData.nestQuote
		var tags = ""
		for convData.nestQuote < nest {
			convData.nestQuote++
			tags = fmt.Sprintf("%s<blockquote>", tags)
			if oldNest != 0 {
				tags = fmt.Sprintf("</p>%s", tags)
			}
		}

		convData.markdownLines[0] = fmt.Sprintf("%s<p>%s", tags, convData.markdownLines[0])
	}

	// close
	convData.quoteTagClose(nest)

	// inline
	convData.inlineConv()
}

// quoteClose
func (convData *convertedData) quoteClose() {
	convData.shiftLine()
	convData.quoteTagClose(0)
}

// quoteTagClose
func (convData *convertedData) quoteTagClose(nest int) {
	if convData.nestQuote > nest {
		for convData.nestQuote > nest {
			convData.markdownLines[0] = fmt.Sprintf("%s</blockquote>", convData.markdownLines[0])
			convData.nestQuote--
		}
		convData.markdownLines[0] = fmt.Sprintf("</p>%s", convData.markdownLines[0])
	}
}
