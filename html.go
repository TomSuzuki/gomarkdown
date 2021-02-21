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

type regList struct {
	reg  string
	html string
}

// MarkdownToHTML ...import markdown text, it will return HTML text
func MarkdownToHTML(markdown string) string {
	var html string
	markdown = convNewline(markdown, "\n")
	lines := strings.Split(markdown, "\n")
	var nestlist []string
	isCode := false

	regList := []regList{
		{`\*\*([^\*]*)\*\*`, "<b>$1</b>"},
		{`!\[(.*)\]\((.*)\)`, "<img alt='$1' src='$2'>"},
		{`\[(.*)\]\((.*)\)`, "<a href='$2'>$1</a>"},
		{`\*([^\*]*)\*|_([^_]*)_|__([^_]*)__`, "<em>$1</em>"},
		{`^#\s([^#]*?$)`, "<h1>$1</h1>"},
		{`^##\s([^#]*?$)`, "<h2>$1</h2>"},
		{`^###\s([^#]*?$)`, "<h3>$1</h3>"},
		{`^####\s([^#]*?$)`, "<h4>$1</h4>"},
		{`^#####\s([^#]*?$)`, "<h5>$1</h5>"},
		{`^######\s([^#]*?$)`, "<h6>$1</h6>"},
		{`^>\s(.*$)`, "<blockquote>$1</blockquote>"},
		{`\s\s$`, "<br>"},
		{`^(\* ){3,}$|^\*.$|^(- ){3,}|^-{3,}$|^(_ ){3,}$|^_{3,}$`, "<hr>"},
		{"~([^~]*)~", "<s>$1</s>"},
		{"`([^`]*)`", "<code>$1</code>"},
	}

	for i := range lines {
		line := lines[i]

		line, isCode = parserCode(line, isCode)
		if !isCode {
			line, nestlist = parserList(line, nestlist)

			for i := range regList {
				re := regexp.MustCompile(regList[i].reg)
				line = re.ReplaceAllString(line, regList[i].html)
			}
		}

		html += line
	}

	s, _ := parserList("", nestlist)
	html += s

	return html
}

// parserCode ...
func parserCode(line string, isCode bool) (string, bool) {
	if len(line) >= 3 && line[:3] == "```" {
		if isCode {
			line = "</code></pre>"
		} else {
			line = "<pre><code>"
		}
		isCode = isCode == false
	}
	return line, isCode
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
		line = line[strings.Index(line, "1. ")+3:]
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
