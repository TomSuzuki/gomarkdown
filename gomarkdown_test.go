/*
 * gomarkdown
 * available at https://github.com/TomSuzuki/gomarkdown/
 *
 * Copyright 2021 TomSuzuki
 * LICENSE: MIT
 *
 */

//package gomarkdown
package gomarkdown_test

import (
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/TomSuzuki/gomarkdown"
)

type testFile struct {
	markdown string
	html     string
}

func Test(t *testing.T) {
	// test case (markdown, html)
	dir := "./testcase/"
	var testfile []testFile

	// get list
	files, _ := ioutil.ReadDir("./testcase/")
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".md" {
			testfile = append(testfile, testFile{
				markdown: dir + file.Name(),
				html:     dir + file.Name() + ".html",
			})
		}
	}

	// test
	for i := range testfile {
		test(testfile[i], t)
	}
}

// speed test
func BenchmarkSpeedTest(b *testing.B) {
	file := "./testcase/00.md"
	md, _ := ioutil.ReadFile(file)

	b.ResetTimer()
	for ct := 0; ct < 1500; ct++ {
		gomarkdown.MarkdownToHTML(string(md))
	}
}

func test(test testFile, t *testing.T) {
	// load
	b, _ := ioutil.ReadFile(test.html)
	sample := string(b)
	b, _ = ioutil.ReadFile(test.markdown)
	answer := string(b)

	// html -> markdown
	answer = gomarkdown.MarkdownToHTML(answer)

	// trim
	sample = strings.NewReplacer("\r\n", "", "\r", "", "\n", "", " ", "", "'", "\"").Replace(sample)
	answer = strings.NewReplacer("\r\n", "", "\r", "", "\n", "", " ", "", "'", "\"").Replace(answer)

	// html
	sampleHTML := template.HTML(sample)
	answerHTML := template.HTML(answer)

	// check
	if sampleHTML != answerHTML {
		t.Logf("☒  failed test: \t%s", test.markdown)
		t.Logf(" - sample: %s", sampleHTML)
		t.Logf(" - answer: %s", answerHTML)
		t.Logf("")
		t.Fail()
	} else {
		t.Logf("☑  success test: \t%s", test.markdown)
	}
}
