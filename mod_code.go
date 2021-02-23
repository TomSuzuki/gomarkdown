package gomarkdown

// codeMarkerConv ...start code lines
func codeMarkerConv(convData convertedData) convertedData {
	if convData.typeChenged {
		convData.markdownLines[0] = "<pre><code>"
	}
	return convData
}

// codeClose ...close code lines
func codeClose(convData convertedData) convertedData {
	if convData.lineType == typeCodeMarker {
		convData.markdownLines[0] = "</code></pre>"
		convData.lineType = typeNone
	}
	return convData
}
