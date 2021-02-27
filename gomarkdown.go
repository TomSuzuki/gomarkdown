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
	typeTableHead
	typeTableBody
	typeQuote
	typeHeader
	typeHorizon
)

// all data
type convertedData struct {
	markdownLines []string
	html          []string
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
	convData.markdownLines = append(strings.Split(strings.NewReplacer("\r\n", "\n", "\r", "\n").Replace(markdown), "\n"), "")

	// closeBlockFunc
	var closeBlockFunc = map[linetype]func(){
		typeTableBody: convData.closeTableBody,
		typeTableHead: convData.closeTableHead,
		typeCode:      convData.closeCode,
		typeList:      convData.closeList,
		typeParagraph: convData.closeParagraph,
		typeQuote:     convData.closeQuote,
	}

	// convBlockFunc
	var convBlockFunc = map[linetype]func(){
		typeTableBody:  convData.convTableBody,
		typeTableHead:  convData.convTableHead,
		typeCodeMarker: convData.convCodeMarker,
		typeList:       convData.convList,
		typeParagraph:  convData.convParagraph,
		typeQuote:      convData.convQuote,
		typeHeader:     convData.convHeader,
		typeHorizon:    convData.convHorizon,
	}

	// lines
	for len(convData.markdownLines) > 0 {
		// if type changed
		oldType := convData.lineType
		convData.lineType = convData.getLineType()
		convData.typeChenged = convData.lineType != oldType

		// close tag
		if f := closeBlockFunc[oldType]; convData.typeChenged && f != nil {
			f()
		}

		// markdown -> html
		if f := convBlockFunc[convData.lineType]; f != nil {
			f()
		}

		// add html
		convData.html = append(convData.html, convData.markdownLines[0])
		convData.markdownLines = convData.markdownLines[1:]
	}

	return strings.Join(convData.html, "\n")
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
	case convData.isList():
		return typeList
	case convData.isQuote():
		return typeQuote
	case convData.isHorizon():
		return typeHorizon
	case convData.isTableBody():
		return typeTableBody
	case convData.isTableHead():
		return typeTableHead
	case convData.isNone():
		return typeNone
	default:
		return typeParagraph
	}
}

// shiftLine ...at the end of the block, it may break the next element if it is not a replacement
func (convData *convertedData) shiftLine() {
	convData.markdownLines = append([]string{""}, convData.markdownLines...)
	convData.lineType = typeNone
}
