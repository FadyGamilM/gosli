package main

import (
	"log"

	"github.com/FadyGamilM/gosli/lexer"
)

func main() {
	tokens, err := lexer.Lex([]rune("(+ 13 (- 12 1))"))
	if err != nil {
		log.Println("error ➜ ", err)
	}
	log.Println("length of tokens ➜ ", len(tokens))
	log.Println("tokens ➜ ", tokens)

}
