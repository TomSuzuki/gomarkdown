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

// define type
const (
	typeNone linetype = iota
	typeParagraph
	typeList
	typeCode
	typeCodeMarker
	typeTable
	typeQuote
	typeHeader
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
		case typeHeader:
			convData.headerConv()
		}

		// add html
		convData.html += convData.markdownLines[0]
		convData.markdownLines = convData.markdownLines[1:]
	}

	return convData.html
}

// getLineType ...determine the type of line
func (convData *convertedData) getLineType() linetype {
	// the higher it is, the higher the priority
	switch true {
	case convData.isCodeMarker():
		return typeCodeMarker
	case convData.isCode():
		return typeCode
	case convData.isHeader():
		return typeHeader
	case convData.isQuote():
		return typeQuote
	case convData.isList():
		return typeList
	case convData.isTable():
		return typeTable
	case convData.isNone():
		return typeNone
	default:
		return typeParagraph
	}
}

// shiftLine ...at the end of the block, it may break the next element if it is not a replacement
func (convData *convertedData) shiftLine() {
	convData.markdownLines = append([]string{""}, convData.markdownLines...)
}
