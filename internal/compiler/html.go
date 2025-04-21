package compiler

import (
	"fmt"
	"log"
	"strings"
)

func ParseToHTML(parser *Parser) string {
	html := strings.Builder{}

	for block := parser.NextBlock(); block.Type != BLK_END; block = parser.NextBlock() {
		switch block.Type {
		case BLK_HEADING:
			props, valid := block.Props.(BlockHeadingProps)
			if !valid {
				log.Fatalf("invalid props of heading")
			}

			heading := fmt.Sprintf("<h%d>%s</h%d>\n", props.Level, block.Content, props.Level)
			html.WriteString(heading)
		case BLK_PARAGRAPH:
			paragraph := fmt.Sprintf("<p>%s</p>\n", block.Content)
			html.WriteString(paragraph)
		}
	}

	return html.String()
}
