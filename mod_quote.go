package gomarkdown

import (
	"fmt"
	"strings"
)

// quoteConv
func (convData *convertedData) quoteConv() {
	// >
	nest := 0
	for {
		n := strings.Index(convData.markdownLines[0], "> ")
		if (n > 0 && strings.Trim(convData.markdownLines[0][:n], " ") != "") || n == -1 {
			break
		} else {
			convData.markdownLines[0] = convData.markdownLines[0][n+2:]
			nest++
		}
	}
	if nest == 0 {
		nest = convData.nestQuote
	}

	// open
	var oldNest = convData.nestQuote
	var tags = ""
	for convData.nestQuote < nest {
		convData.nestQuote++
		if convData.nestQuote == nest {
			tags = fmt.Sprintf("%s<blockquote><p>", tags)
		} else {
			tags = fmt.Sprintf("%s<blockquote>", tags)
		}
		if oldNest != 0 {
			tags = fmt.Sprintf("</p>%s", tags)
		}
	}
	convData.markdownLines[0] = tags + convData.markdownLines[0]

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
	var oldNest = convData.nestQuote

	for convData.nestQuote > nest {
		convData.markdownLines[0] = fmt.Sprintf("%s</blockquote>", convData.markdownLines[0])
		convData.nestQuote--
	}

	if oldNest > nest {
		convData.markdownLines[0] = fmt.Sprintf("</p>%s", convData.markdownLines[0])
	}
}
