/*
 * gomarkdown
 * available at https://github.com/TomSuzuki/gomarkdown/
 *
 * Copyright 2021 TomSuzuki
 * LICENSE: MIT
 *
 */

package gomarkdown

import (
	"strings"
)

// type of line
type linetype int

const (
	typeNone linetype = iota
	typeParagraph
	typeList
	typeCode
	typeCodeMarker
	typeTable
	typeQuote
)

// all data
type convertedData struct {
	markdownLines []string
	html          string
	lineType      linetype
	tableAlign    []string
	listNest      []string
	typeChenged   bool
	nestQuote     int
}

// MarkdownToHTML ...import markdown text, it will return HTML text
func MarkdownToHTML(markdown string) string {
	// init
	var convData convertedData
	convData.lineType = typeNone
	convData.markdownLines = append(strings.Split(strings.NewReplacer("\r\n", "\n", "\r", "\n", "\n", "\n").Replace(markdown), "\n"), "")

	// lines
	for len(convData.markdownLines) > 0 {
		// get line type
		oldType := convData.lineType
		convData.lineType = convData.getLineType()

		// if type changed
		convData.typeChenged = convData.lineType != oldType
		if convData.typeChenged {
			switch oldType {
			case typeTable:
				convData.tableClose()
			case typeCode:
				convData.codeClose()
			case typeList:
				convData.listClose()
			case typeParagraph:
				convData.paragraphClose()
			case typeQuote:
				convData.quoteClose()
			}
		}

		// markdown -> html
		switch convData.lineType {
		case typeTable:
			convData.tableConv()
		case typeCodeMarker:
			convData.codeMarkerConv()
		case typeList:
			convData.listConv()
		case typeParagraph:
			convData.paragraphConv()
		case typeQuote:
			convData.quoteConv()
		}

		// add html
		convData.html += convData.markdownLines[0]
		convData.markdownLines = convData.markdownLines[1:]
	}

	return convData.html
}

// getLineType ...determine the type of line
func (convData *convertedData) getLineType() linetype {
	var line = convData.markdownLines[0]

	// check
	if (line + "   ")[:3] == "```" {
		return typeCodeMarker
	} else if convData.lineType == typeCodeMarker || convData.lineType == typeCode {
		return typeCode
	} else if (strings.Trim(line, " ") + " ")[:1] == ">" || (convData.lineType == typeQuote && strings.Trim(line, " ") != "") {
		return typeQuote
	} else if (strings.Trim(line, " ") + "  ")[:2] == "- " || (strings.Trim(line, " ") + "   ")[:3] == "1. " {
		return typeList
	} else if (strings.Trim(line, " ") + " ")[:1] == "|" && strings.Count(line, "|") > 1 {
		return typeTable
	} else if strings.Trim(line, " ") == "" {
		return typeNone
	}
	return typeParagraph
}
