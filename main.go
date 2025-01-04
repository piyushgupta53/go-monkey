package main

import (
	"fmt"

	"github.com/piyushgupta53/go-monkey/lexer"
)

func main() {
	input := `let x = 5;`
	l := lexer.New(input)

	// Read all tokens
	for {
		tok := l.NextToken()
		fmt.Printf("%+v\n", tok)

		if tok.Type == "EOF" {
			break
		}
	}
}
