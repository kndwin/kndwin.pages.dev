package scripts

import (
	"fmt"
	"log"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func MdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func ReadMdFromFlatDirectory(dir string) {
	entries, err := os.ReadDir(dir)
	failIf(err, "Coudln't read directory")

	for entry = range entries {
		fmt.Print(entry)
	}

}

func debug(message string) {
	log.Print(message)
}

func failIf(err error, message string) {
	if err != nil {
		log.Panic(message, err)
	}
}
