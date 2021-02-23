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
)

// all data
type convertedData struct {
	markdownLines []string
	html          string
	lineType      linetype
	tableAlign    []string
	listNest      []string
	typeChenged   bool
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
		convData.lineType = getLineType(convData)

		// if type changed
		convData.typeChenged = convData.lineType != oldType
		if convData.typeChenged {
			switch oldType {
			case typeTable:
				convData = tableClose(convData)
			case typeCode:
				convData = codeClose(convData)
			case typeList:
				convData = listClose(convData)
			case typeParagraph:
				convData = paragraphClose(convData)
			}
		}

		// markdown -> html
		switch convData.lineType {
		case typeTable:
			convData = tableConv(convData)
		case typeCodeMarker:
			convData = codeMarkerConv(convData)
		case typeList:
			convData = listConv(convData)
		case typeParagraph:
			convData = paragraphConv(convData)
		}

		// add html
		convData.html += convData.markdownLines[0]
		convData.markdownLines = convData.markdownLines[1:]
	}

	return convData.html
}

// getLineType ...determine the type of line
func getLineType(convData convertedData) linetype {
	var line = convData.markdownLines[0]

	// check
	if (line + "   ")[:3] == "```" {
		return typeCodeMarker
	} else if convData.lineType == typeCodeMarker || convData.lineType == typeCode {
		return typeCode
	} else if (strings.Trim(line, " ") + "  ")[:2] == "- " || (strings.Trim(line, " ") + "   ")[:3] == "1. " {
		return typeList
	} else if (strings.Trim(line, " ") + " ")[:1] == "|" && strings.Count(line, "|") > 1 {
		return typeTable
	} else if strings.Trim(line, " ") == "" {
		return typeNone
	}
	return typeParagraph
}
