package gomarkdown

import (
	"fmt"
	"strings"
)

// quoteConv
func (convData *convertedData) quoteConv() {
	// >
	nest := 0
	for flg := true; flg; {
		n := strings.Index(convData.markdownLines[0], "> ")
		if n > 0 && strings.Trim(convData.markdownLines[0][:n], " ") != "" {
			n = -1
		}
		if n != -1 {
			convData.markdownLines[0] = convData.markdownLines[0][n+2:]
			nest++
		} else {
			flg = false
		}
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
	convData.quoteTagClose(nest, true)

	// inline
	convData.inlineConv()
}

// quoteClose
func (convData *convertedData) quoteClose() {
	convData.quoteTagClose(0, false)
}

// quoteTagClose
func (convData *convertedData) quoteTagClose(nest int, inquote bool) {
	var oldNest = convData.nestQuote

	for convData.nestQuote > nest {
		if inquote {
			convData.markdownLines[0] = fmt.Sprintf("%s</blockquote>", convData.markdownLines[0])
		} else {
			convData.markdownLines[0] = fmt.Sprintf("</blockquote>%s", convData.markdownLines[0])
		}
		convData.nestQuote--
	}

	if oldNest > nest {
		convData.markdownLines[0] = fmt.Sprintf("</p>%s", convData.markdownLines[0])
	}
}
