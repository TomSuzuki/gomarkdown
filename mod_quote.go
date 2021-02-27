package gomarkdown

import (
	"strings"
)

// isQuote
func (convData *convertedData) isQuote() bool {
	return (strings.Trim(convData.markdownLines[0], " ") + "  ")[:2] == "> " || (convData.lineType == typeQuote && !convData.isNone())
}

// convQuote
func (convData *convertedData) convQuote() {
	// >
	nest := strings.Count(convData.markdownLines[0], "> ")
	if nest == 0 {
		nest = convData.nestQuote
	} else {
		convData.markdownLines[0] = convData.markdownLines[0][2*nest:]
	}

	// open
	if convData.nestQuote < nest {
		var text []string
		var oldNest = convData.nestQuote
		for convData.nestQuote < nest {
			convData.nestQuote++
			if oldNest != 0 {
				text = append(text, "</p>")
			}
			text = append(text, "<blockquote>")

		}
		text = append(text, "<p>")
		text = append(text, convData.markdownLines[0])
		convData.markdownLines[0] = strings.Join(text, "")
	}

	// close
	convData.quoteTagClose(nest)

	// inline
	convData.inlineConv()
}

// closeQuote
func (convData *convertedData) closeQuote() {
	convData.shiftLine()
	convData.quoteTagClose(0)
}

// quoteTagClose
func (convData *convertedData) quoteTagClose(nest int) {
	if convData.nestQuote > nest {
		var text []string
		text = append(text, "</p>")
		text = append(text, convData.markdownLines[0])
		for convData.nestQuote > nest {
			text = append(text, "</blockquote>")
			convData.nestQuote--
		}
		convData.markdownLines[0] = strings.Join(text, "")
	}
}
