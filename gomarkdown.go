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
	"strings"
)

// Regular expression information ...Regular expression information
type regList struct {
	reg  string
	html string
}

// MarkdownToHTML ...import markdown text, it will return HTML text
func MarkdownToHTML(markdown string) string {
	// init
	var html string
	var nestlist []string
	isCode := false
	regs := listRegInfo()

	// markdown -> line
	markdown = convNewline(markdown, "\n")
	lines := strings.Split(markdown, "\n")

	// lines -> line
	for i := range lines {
		line := lines[i]
		// code line?
		line, isCode = parserCode(line, isCode)
		if !isCode {
			// list
			line, nestlist = parserList(line, nestlist)

			// reg replace
			for i := range regs {
				line = regexp.MustCompile(regs[i].reg).ReplaceAllString(line, regs[i].html)
			}
		}

		// add
		html += line
	}

	// </ul> or </ol>
	s, _ := parserList("", nestlist)
	html += s

	return html
}

// listRegInfo ...replacement data
func listRegInfo() []regList {
	return []regList{
		{`\*\*([^\*]*)\*\*`, "<strong>$1</strong>"},
		{`!\[(.*)\]\((.*)\)`, "<img alt='$1' src='$2'>"},
		{`\[(.*)\]\((.*)\)`, "<a href='$2'>$1</a>"},
		{`\*([^\*]*)\*|_([^_]*)_|__([^_]*)__`, "<em>$1</em>"},
		{`^#\s([^#]*?$)`, "<h1>$1</h1>"},
		{`^##\s([^#]*?$)`, "<h2>$1</h2>"},
		{`^###\s([^#]*?$)`, "<h3>$1</h3>"},
		{`^####\s([^#]*?$)`, "<h4>$1</h4>"},
		{`^#####\s([^#]*?$)`, "<h5>$1</h5>"},
		{`^######\s([^#]*?$)`, "<h6>$1</h6>"},
		{`^>\s(.*$)`, "<blockquote><p>$1</p></blockquote>"},
		{`\s\s$`, "<br>"},
		{`^(\* ){3,}$|^\*.$|^(- ){3,}|^-{3,}$|^(_ ){3,}$|^_{3,}$`, "<hr>"},
		{"~([^~]*)~", "<s>$1</s>"},
		{"`([^`]*)`", "<code>$1</code>"},
	}
}

// parserCode ...
func parserCode(line string, isCode bool) (string, bool) {
	if (line + "   ")[:3] == "```" {
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
	// init
	var text string
	var apd string
	var isOpen = false
	var isClose = false
	var info = []struct {
		html     string
		markdown string
	}{
		{"ul", "- "},
		{"ol", "1. "},
	}

	// this line
	nest := len(nestlist)
	for i := range info {
		f := info[i]
		if strings.Index(strings.Trim(line, " "), f.markdown) == 0 {
			apd = f.html
			nest = 1 + strings.Index(line, f.markdown)/2
			line = line[strings.Index(line, f.markdown)+len(f.markdown):]
		}
	}

	// open list
	for nest > len(nestlist) {
		nestlist = append(nestlist, apd)
		text += "<" + apd + ">"
		isOpen = true
	}

	// li
	if apd != "" {
		text += "<li>" + line
	} else {
		text += line
	}

	// no text
	if strings.Trim(line, " ") == "" {
		nest = 0
	}

	// close list
	var close string
	for nest < len(nestlist) {
		isClose = true
		close += ("</li></" + nestlist[len(nestlist)-1] + ">")
		nestlist = nestlist[:len(nestlist)-1]
	}
	if len(nestlist) > 0 && isClose {
		text = "</li>" + text
	}
	text = close + text

	// close li
	if len(nestlist) > 0 && !isOpen && !isClose {
		text = "</li>" + text
	}

	return text, nestlist
}

// convert newline
func convNewline(str string, nlcode string) string {
	return strings.NewReplacer("\r\n", nlcode, "\r", nlcode, "\n", nlcode).Replace(str)
}
