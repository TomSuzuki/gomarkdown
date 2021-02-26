package gomarkdown

import "regexp"

const horizonMD = `^(\* ){3,}$|^\*.$|^(- ){3,}|^-{3,}$|^(_ ){3,}$|^_{3,}$`
const horizonHTML = "<hr>"

// isHorizon
func (convData *convertedData) isHorizon() bool {
	var line = convData.markdownLines[0]
	return line != regexp.MustCompile(horizonMD).ReplaceAllString(line, horizonHTML)
}

// horizonConv
func (convData *convertedData) horizonConv() {
	var line = convData.markdownLines[0]
	convData.markdownLines[0] = regexp.MustCompile(horizonMD).ReplaceAllString(line, horizonHTML)
}
