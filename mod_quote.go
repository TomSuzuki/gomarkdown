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
	for convData.nestQuote < nest {
		convData.markdownLines[0] = fmt.Sprintf("<blockquote><p>%s", convData.markdownLines[0])
		if convData.nestQuote != 0 {
			convData.markdownLines[0] = fmt.Sprintf("</p>%s", convData.markdownLines[0])
		}
		convData.nestQuote++
	}

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
