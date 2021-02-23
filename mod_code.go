package gomarkdown

// codeMarkerConv ...start code lines
func (convData *convertedData) codeMarkerConv() {
	if convData.typeChenged {
		convData.markdownLines[0] = "<pre><code>"
	}
}

// codeClose ...close code lines
func (convData *convertedData) codeClose() {
	if convData.lineType == typeCodeMarker {
		convData.markdownLines[0] = "</code></pre>"
		convData.lineType = typeNone
	}
}
