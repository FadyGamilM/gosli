package lexer

import "unicode"

func Lex(inp []rune) ([]Token, error) {
	tokens := make([]Token, 0)

	cursor := 0

	for cursor < len(inp) {
		// drop all white spaces
		cursor = dropWhiteSpaces(inp, cursor)

		// check if its a open/closed paraenthesis > if its, continue
	}

	return tokens, nil
}

func isWhiteSpace(r rune) bool {
	return unicode.IsSpace(r)
}

func dropWhiteSpaces(inp []rune, cursor int) int {
	return 0
}

/*
	inp = "add 123 452", cursor = 4
	-> so the cursor is on the "1"
	so we expect the returned token to be "123"
*/
func LexOperands(inp []rune, cursor int) (int, *Token) {
	tokenVal := make([]rune, 0)
	for cursor < len(inp) {
		if IsIntNumberOperand(inp[cursor]) {
			// append the value because we need to grap the entire value "123" if we have a cursor on "1"
			tokenVal = append(tokenVal, inp[cursor])
			// move the cursor one rune
			cursor++
			// continue to the next iteration
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

func IsIntNumberOperand(val rune) bool {
	return unicode.IsNumber(val)
}
