package main

import (
	"log"

	"github.com/FadyGamilM/gosli/lexer"
)

func main() {
	tokens, err := lexer.Lex([]rune("(+ 5 6)"))
	if err != nil {
		log.Println("error ➜ ", err)
	}
	log.Println("length of tokens ➜ ", len(tokens))
	log.Println("tokens ➜ ", tokens)
}
