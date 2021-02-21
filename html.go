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
	"regexp"
	"strconv"
	"strings"
)

// MarkdownToHTML ...import markdown text, it will return HTML text
func MarkdownToHTML(markdown string) string {
	var html string
	markdown = convNewline(markdown, "\n")
	lines := strings.Split(markdown, "\n")
	var nestlist []string
	isCode := false

	for i := range lines {
		line := lines[i]

		line, isCode = parserCode(line, isCode)
		if !isCode {
			line, nestlist = parserList(line, nestlist)
			line = parserHeader(line)
			line = parser(line, "**", "b")
			line = parser(line, "__", "b")
			line = parser(line, "*", "em")
			line = parser(line, "_", "em")
			line = parser(line, "~", "s")
			line = parser(line, "`", "code")
			line = parser(line, "  ", "br")
			line = parserHr(line)
		}

		html += line
	}

	s, _ := parserList("", nestlist)
	html += s

	return html
}

// parserCode ...
func parserCode(line string, isCode bool) (string, bool) {
	if line == "```" {
		if isCode {
			line = "</pre></code>"
		} else {
			line = "<code><pre>"
		}
		isCode = isCode == false
	}
	return line, isCode
}

// CleanHTML ...
func CleanHTML(html string) string {
	re := regexp.MustCompile(`</([^>]+?)>`)
	return re.ReplaceAllString(html, "</$1>\n")
}

// parserList ...
func parserList(line string, nestlist []string) (string, []string) {
	text := ""
	apd := ""
	// nest
	nest := len(nestlist)
	if strings.Index(strings.Trim(line, " "), "- ") == 0 {
		apd = "ul"
		nest = 1 + strings.Index(line, "- ")/2
		line = line[strings.Index(line, "- ")+2:]
	}
	if strings.Index(strings.Trim(line, " "), "1. ") == 0 {
		apd = "ol"
		nest = 1 + strings.Index(line, "1. ")/2
		line = line[strings.Index(line, "1. ")+2:]
	}
	// open
	for nest > len(nestlist) {
		nestlist = append(nestlist, apd)
		text += "<" + apd + ">"
	}
	// li
	if apd != "" {
		text += "<li>" + line + "</li>"
	} else {
		text += line
	}
	// close
	if strings.Trim(line, " ") == "" {
		nest = 0
	}
	for nest < len(nestlist) {
		text += ("</" + nestlist[len(nestlist)-1] + ">")
		nestlist = nestlist[:len(nestlist)-1]
	}

	return text, nestlist
}

// convert newline
func convNewline(str string, nlcode string) string {
	return strings.NewReplacer("\r\n", nlcode, "\r", nlcode, "\n", nlcode).Replace(str)
}

// parserHeader ...h1-h6
func parserHeader(line string) string {
	text := line
	h := strings.Split(line, " ")[0]
	if strings.NewReplacer("#", "").Replace(h) == "" {
		n := strings.Count(h, "#")
		if 1 <= n && n <= 6 {
			s := strconv.Itoa(n)
			text = "<h" + s + ">" + text[n+1:] + "</h" + s + ">"
		}
	}
	return text
}

// parser ...?***? -> <!>***</!>
func parser(line string, md string, tag string) string {
	s := strings.Split(line, md)
	var text string
	if len(s) >= 2 {
		for i := range s {
			text += s[i]
			if len(s)-1 != i {
				if i%2 == 0 {
					text += ("<" + tag + ">")
				} else {
					text += ("</" + tag + ">")
				}
			}
		}
	} else {
		text = line
	}
	return text
}

// parserHr ...--- -> <hr>
func parserHr(line string) string {
	if line != "" && strings.NewReplacer(" ", "").Replace(line) != "" {
		temp := strings.NewReplacer(" ", "", "*", "-").Replace(line)
		if len(temp) >= 3 && strings.NewReplacer("-", "").Replace(temp) == "" {
			return "<hr>"
		}
	}
	return line
}
