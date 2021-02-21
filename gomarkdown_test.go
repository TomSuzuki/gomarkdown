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
	"testing"

	"github.com/TomSuzuki/gomarkdown"
)

func Test(t *testing.T) {
	test := `
# Header 1
### Header 3
--
<br>  
**Bold****Bold**

- item
- item
- item

1. aaa
1. bbb
1. ccc
`

	t.Log(gomarkdown.MarkdownToHTML(test))
}
