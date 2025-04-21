package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/SadS4ndWiCh/painkiller/internal/compiler"
)

func main() {
	input := flag.String("input", "input.pnk", "-input=/path/to/input.pnk")
	flag.Parse()

	content, err := os.ReadFile(*input)
	if err != nil {
		log.Panicf("failed to read '%s' file", *input)
	}

	lexer := compiler.NewLexer(content)
	parser := compiler.NewParser(lexer)
	blocks := parser.Parse()

	html := compiler.ParseToHTML(blocks)

	fmt.Println(html)
}
