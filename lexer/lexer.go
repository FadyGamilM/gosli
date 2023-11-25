package lexer

import (
	"log"
	"unicode"
)

func Lex(inp []rune) ([]Token, error) {
	tokens := make([]Token, 0)

	i := 0
	cursor := 0

	for cursor < len(inp) {
		log.Printf("➜ Iteration No.{%v}, the value of cursor is {%v}\n", i, cursor)
		i++

		// drop all white spaces untill we hit something, so we start discover it
		cursor = dropWhiteSpaces(inp, cursor)
		// if the cursor = length of the soruce, so we have a white space line of code and our cursor is at the end
		if cursor == len(inp) {
			log.Println("➜ there is nothing to lex, we have only white space/s")
			break
		}

		// check if its a open/closed paraenthesis > if its, continue
		var lexedToken *Token
		cursor, lexedToken = LexOpenParenthesis(inp, cursor)
		if lexedToken != nil {
			tokens = append(tokens, *lexedToken)
			// if we have a space after this token, we should go back and ignore it first by increasing our cursor
			continue
		}

		log.Println("cursor is =====> ", cursor)
		// check if its an operator (+, -, *, /)
		cursor, lexedToken = LexOperator(inp, cursor)
		if lexedToken != nil {
			tokens = append(tokens, *lexedToken)
			// if we have a space after this token, we should go back and ignore it first by increasing our cursor
			continue
		}

		cursor, lexedToken = LexOperands(inp, cursor)
		if lexedToken != nil {
			tokens = append(tokens, *lexedToken)
			// if we have a space after this token, we should go back and ignore it first by increasing our cursor
			continue
		}

		_, lexedToken = LexReservedKeywords(inp, cursor)
		if lexedToken != nil {
			tokens = append(tokens, *lexedToken)
			// if we have a space after this token, we should go back and ignore it first by increasing our cursor
			continue
		}
	}

	return tokens, nil
}

func LexOperator(inp []rune, cursor int) (int, *Token) {
	if IsOperator(inp[cursor]) {
		cursor++
		return cursor, &Token{
			Kind: OperatorToken,
			Val:  string(inp[cursor-1]),
		}
	}
	return cursor, nil
}

func IsOperator(val rune) bool {
	if val == '+' || val == '-' || val == '*' || val == '/' {
		return true
	}
	return false
}

func LexOpenParenthesis(inp []rune, cursor int) (int, *Token) {
	if inp[cursor] == '(' || inp[cursor] == ')' {
		var kind TokenKind
		if inp[cursor] == '(' {
			kind = OpeningParenthesis
		}else {
			kind = ClosedParenthesis
		}
		// log.Printf("[Lex-Open-Parenthesis] ➜\tThe value at cursor = {%v} is {%v}", cursor, inp[cursor])
		token := &Token{
			Kind: kind,
			Val:  string(inp[cursor]),
		}
		cursor++
		return cursor, token
	}
	return cursor, nil
}

func isWhiteSpace(r rune) bool {
	return unicode.IsSpace(r)
}

// dropWhiteSpaces pass the white spaces untill it hits something, it returns the cursor pointing to this something so we can classify its token kind
/*
	- inp : which is the entire source, so we can skip the white spaces for e.g if we have white spaces at the start of the line
	- cursor : current cursor pointing on the current rune that we need to know if its a white space or not
*/
func dropWhiteSpaces(inp []rune, cursor int) int {
	for cursor < len(inp) {
		if isWhiteSpace(inp[cursor]) {
			// log.Printf("➜ found a white space when cursor = {%v}\n", cursor)
			cursor++
			continue
		}

		// log.Printf("➜ going to break, found no white space when cursor = {%v}\n", cursor)
		// we hit something so lets go back to lex() to classify it
		break
	}

	// return where we find this non-white-space thing
	return cursor
}

/*
inp = "add 123 452", cursor = 4
-> so the cursor is on the "1"
so we expect the returned token to be "123"
*/
func LexOperands(inp []rune, cursor int) (int, *Token) {
	tokenVal := make([]rune, 0)
	for cursor < len(inp) {
		if IsIntNumberOperand(inp[cursor]) || IsCharOperand(inp[cursor]) {
			// append the value because we need to grap the entire value "123" if we have a cursor on "1"
			tokenVal = append(tokenVal, inp[cursor])
			// move the cursor one rune
			cursor++
			continue
		}

		break
	}

	if len(tokenVal) == 0 {
		return cursor, nil
	}

	return cursor, &Token{
		Kind: OperandToken,
		Val:  string(tokenVal),
	}
}

// func LexIdentifierToken (inp )

// reserved keywords : [lambda, if, let, define, cond, else, begin]
func LexReservedKeywords(inp []rune, cursor int) (int, *Token) {

	keyword := make([]rune, 0)
	for cursor < len(inp) {
		if IsCharOperand(inp[cursor]) {
			keyword = append(keyword, inp[cursor])
			cursor++
			continue
		}
		break
	}

	if len(keyword) == 0 {
		return cursor, nil
	}

	return cursor, &Token{
		Kind: ReservedKeyworkdToken,
		Val:  string(keyword),
	}
}

// this takes the entire word not a char of it
// so we pass a rune by rune to the IsCharOperand and once we have a complete word, we pass it to the reserved keyword checker to check if its one of them or its just a normal operand
func IsReservedKeyword(word string) bool {
	if word == "lambda" ||
		word == "let" ||
		word == "if" ||
		word == "define" ||
		word == "cond" ||
		word == "else" ||
		word == "begin" {
		return true
	}
	return false
}
func IsIntNumberOperand(val rune) bool {
	return unicode.IsNumber(val)
}

func IsCharOperand(val rune) bool {
	return unicode.IsLetter(val)
}
