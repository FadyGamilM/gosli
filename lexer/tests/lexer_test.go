package tests

import (
	"testing"

	"github.com/FadyGamilM/gosli/lexer"
	"github.com/stretchr/testify/assert"
)

// my function is expected to work correctly if we passed to it a right positioned cursor, in the lex function i will move the cursor but here i am sticking with a fixed cursor pointing to the right position and i need to check just the logic of it on how it will return the right token and move the cursor to the right position after the token
func TestLexOperands(t *testing.T) {
	testCases := []struct {
		source         []rune
		cursor         int
		expectedToken  string
		expectedCursor int
	}{
		{
			source:         []rune("+ 6 5"),
			cursor:         2,
			expectedToken:  "6",
			expectedCursor: 3,
		},
		{
			source:         []rune("add abc53 3"),
			cursor:         10,
			expectedToken:  "3",
			expectedCursor: 11,
		},
		{
			source:         []rune("add abc53 3"),
			cursor:         4,
			expectedToken:  "abc53",
			expectedCursor: 9,
		},
	}

	for _, testCase := range testCases {
		// t.Logf("{test-case no.%v}", idx)
		currCursor, lexedToken := lexer.LexOperands(testCase.source, testCase.cursor)
		assert.Equal(t, testCase.expectedCursor, currCursor)
		assert.Equal(t, testCase.expectedToken, lexedToken.Val)
	}
}

func TestLexOperator(t *testing.T) {
	testCases := []struct {
		source         []rune
		cursor         int
		expectedToken  string
		expectedCursor int
	}{
		// now define the expected values ..
		{
			source:         []rune("+ 5 6"),
			cursor:         0,
			expectedToken:  "+",
			expectedCursor: 1,
		},
		{
			source:         []rune("(+ 5 (*6 7))"),
			cursor:         6,
			expectedToken:  "*",
			expectedCursor: 7,
		},
	}
	for _, testCase := range testCases {
		actualCursor, actualTokenVal := lexer.LexOperator(testCase.source, testCase.cursor)
		assert.Equal(t, testCase.expectedToken, actualTokenVal.Val)
		assert.Equal(t, lexer.OperatorToken, actualTokenVal.Kind)
		assert.Equal(t, testCase.expectedCursor, actualCursor)
	}
}
