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
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// MarkdownToHTML ...import markdown text, it will return HTML text
func MarkdownToHTML(markdown string) string {
	var convData convertedData
	convData.lineType = typeNone
	convData.html = ""
	convData.markdownLines = append(strings.Split(strings.NewReplacer("\r\n", "\n", "\r", "\n", "\n", "\n").Replace(markdown), "\n"), "")

	for convData.markdownLines != nil {
		// type change
		line := convData.markdownLines[0]
		oldType := convData.lineType
		if (line + "   ")[:3] == "```" {
			convData.lineType = typeCodeMarker
		} else if convData.lineType == typeCodeMarker || convData.lineType == typeCode {
			convData.lineType = typeCode
		} else if (strings.Trim(line, " ") + "  ")[:2] == "- " || (strings.Trim(line, " ") + "   ")[:3] == "1. " {
			convData.lineType = typeList
		} else if (strings.Trim(line, " ") + " ")[:1] == "|" && strings.Count(line, "|") > 1 {
			convData.lineType = typeTable
		} else if strings.Trim(line, " ") == "" {
			convData.lineType = typeNone
		} else {
			convData.lineType = typeParagraph
		}

		// if type changed
		typeChenged := convData.lineType != oldType
		if typeChenged {
			switch oldType {
			case typeTable:
				convData.markdownLines[0] = fmt.Sprintf("</table>%s", convData.markdownLines[0])
			case typeCode:
				if convData.lineType == typeCodeMarker {
					convData.markdownLines[0] = "</code></pre>"
					convData.lineType = typeNone
				}
			case typeList:
				var text = ""
				for i := len(convData.listNest) - 1; i >= 0; i-- {
					text = fmt.Sprintf("%s</li></%s>", text, convData.listNest[i])
				}
				convData.markdownLines[0] = fmt.Sprintf("%s%s", text, convData.markdownLines[0])
				convData.listNest = nil
			case typeParagraph:
				convData.markdownLines[0] = fmt.Sprintf("</p>%s", convData.markdownLines[0])
			}
		}

		// markdown -> html
		switch convData.lineType {
		case typeTable:
			convData = tableConv(convData)
			convData, _ = inlineConv(convData)
		case typeCodeMarker:
			if oldType != typeCode {
				convData.markdownLines[0] = "<pre><code>"
			}
		case typeList:
			convData = listConv(convData)
			convData, _ = inlineConv(convData)
		case typeParagraph:
			var inline bool
			convData, inline = inlineConv(convData)
			if !inline && oldType == typeParagraph {
				convData.lineType = typeNone
				convData.markdownLines[0] = fmt.Sprintf("</p>%s", convData.markdownLines[0])
			} else if !inline {
				convData.lineType = typeNone
			} else if inline && typeChenged {
				convData.markdownLines[0] = fmt.Sprintf("<p>%s", convData.markdownLines[0])
			}
		}

		// add
		convData.html += convData.markdownLines[0]
		if len(convData.markdownLines) > 1 {
			convData.markdownLines = convData.markdownLines[1:]
		} else {
			convData.markdownLines = nil
		}

	}

	return convData.html
}

// tableConv ...
func tableConv(convData convertedData) convertedData {
	var tag = "td"
	var text = ""

	// align
	if convData.tableAlign == nil {
		tag = "th"
		text = "<table>"
		if len(convData.markdownLines) == 1 {
			return convData
		}
		alignLine := strings.Split(convData.markdownLines[1], "|")
		if len(alignLine) < 2 {
			return convData
		}
		if len(convData.markdownLines) > 2 {
			convData.markdownLines = append([]string{convData.markdownLines[0]}, convData.markdownLines[2:]...)
		} else {
			convData.markdownLines = []string{convData.markdownLines[0]}
		}
		alignLine = alignLine[1 : len(alignLine)-1]
		for i := range alignLine {
			colonR := string(alignLine[i][len(alignLine[i])-1]) == ":"
			colonL := string(alignLine[i][0]) == ":"
			if colonR && !colonL {
				convData.tableAlign = append(convData.tableAlign, "right")
			} else if !colonR && colonL {
				convData.tableAlign = append(convData.tableAlign, "left")
			} else {
				convData.tableAlign = append(convData.tableAlign, "center")
			}
		}
	}

	// <tr>
	var reg = `\|`
	var htm = ""
	for i := range convData.tableAlign {
		reg += `([^|]*)\|`
		htm += fmt.Sprintf("<%s align='%s'>$%s</%s>", tag, convData.tableAlign[i], strconv.Itoa(i+1), tag)
	}
	reg += `$`
	text += "<tr>" + regexp.MustCompile(reg).ReplaceAllString(convData.markdownLines[0], htm) + "</tr>"
	convData.markdownLines[0] = text

	return convData
}

// inlineConv ...
func inlineConv(convData convertedData) (convertedData, bool) {
	var inline = true
	var regexpInfo = listRegInfo()

	for i := range regexpInfo {
		line := convData.markdownLines[0]
		line = regexp.MustCompile(regexpInfo[i].regexp).ReplaceAllString(line, regexpInfo[i].html)
		if convData.markdownLines[0] != line && !regexpInfo[i].isInline {
			inline = false
		}
		convData.markdownLines[0] = line
	}

	return convData, inline
}

// listConv ...
func listConv(convData convertedData) convertedData {
	var line = convData.markdownLines[0]
	var text = ""
	var tag = ""
	var nest = 0
	var oldNest = len(convData.listNest)
	var info = []struct {
		html     string
		markdown string
	}{
		{"ul", "- "},
		{"ol", "1. "},
	}

	// list type
	for i := range info {
		if strings.Index(strings.Trim(line, " "), info[i].markdown) == 0 {
			tag = info[i].html
			nest = 1 + strings.Index(line, info[i].markdown)/2
			line = line[strings.Index(line, info[i].markdown)+len(info[i].markdown):]
		}
	}

	// open <ul> or <ol>
	for nest > len(convData.listNest) {
		convData.listNest = append(convData.listNest, tag)
		text = fmt.Sprintf("<%s>", tag)
	}

	// open <li>
	text = fmt.Sprintf("%s<li>%s", text, line)

	// close </ul> or </ol>
	var closeTags = ""
	for nest < len(convData.listNest) {
		closeTags = fmt.Sprintf("</%s>%s", convData.listNest[len(convData.listNest)-1], closeTags)
		if len(convData.listNest) > 0 {
			closeTags = fmt.Sprintf("%s</li>", closeTags)
		}
		convData.listNest = convData.listNest[:len(convData.listNest)-1]
	}
	text = fmt.Sprintf("%s%s", closeTags, text)

	// close </li>
	if oldNest >= nest {
		text = fmt.Sprintf("</li>%s", text)
	}

	// data
	convData.markdownLines[0] = text

	return convData
}

const (
	typeNone       = iota
	typeParagraph  = iota
	typeList       = iota
	typeCode       = iota
	typeCodeMarker = iota
	typeTable      = iota
)

type convertedData struct {
	markdownLines []string
	html          string
	lineType      int
	tableAlign    []string
	listNest      []string
}

type regList struct {
	regexp   string
	html     string
	isInline bool
}

func listRegInfo() []regList {
	return []regList{
		{`\*\*([^\*]*)\*\*`, "<strong>$1</strong>", true},
		{`!\[(.*?)\]\((.*?)\)`, "<img alt='$1' src='$2'>", true},
		{`\[(.*)\]\((.*)\)`, "<a href='$2'>$1</a>", true},
		{`\*([^\*]*)\*|_([^_]*)_|__([^_]*)__`, "<em>$1</em>", true},
		{`^#\s([^#]*?$)`, "<h1>$1</h1>", false},
		{`^##\s([^#]*?$)`, "<h2>$1</h2>", false},
		{`^###\s([^#]*?$)`, "<h3>$1</h3>", false},
		{`^####\s([^#]*?$)`, "<h4>$1</h4>", false},
		{`^#####\s([^#]*?$)`, "<h5>$1</h5>", false},
		{`^######\s([^#]*?$)`, "<h6>$1</h6>", false},
		{`^>\s(.*$)`, "<blockquote><p>$1</p></blockquote>", false},
		{`\s\s$`, "<br>", true},
		{`^(\* ){3,}$|^\*.$|^(- ){3,}|^-{3,}$|^(_ ){3,}$|^_{3,}$`, "<hr>", false},
		{"~([^~]*)~", "<s>$1</s>", true},
		{"`([^`]*)`", "<code>$1</code>", true},
	}
}
