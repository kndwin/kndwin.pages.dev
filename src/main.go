package main

import (
	"fmt"
	"kndwin.pages.dev/scripts"
)

var mds = `# header

Sample text.

[link](http://example.com)
`

func main() {

	md := []byte(mds)
	html := scripts.MdToHTML(md)

	scripts.ReadMdFromFlatDirectory("./blogs/2024/")

	fmt.Printf("--- Markdown:\n%s\n\n--- HTML:\n%s\n", md, html)
}
