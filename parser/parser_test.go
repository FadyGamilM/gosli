package parser

import (
	"testing"

	"github.com/FadyGamilM/gosli/lexer"
	"github.com/stretchr/testify/assert"
)

func TestNoRecursionParser(t *testing.T) {
	tests := []struct {
		tokens []lexer.Token
		output AST
	}{
		// {
		// 	tokens: []lexer.Token{
		// 		{Kind: lexer.OpeningParenthesis, Val: "("},
		// 		{Kind: lexer.OperatorToken, Val: "+"},
		// 		{Kind: lexer.OperandToken, Val: "13"},
		// 		{Kind: lexer.OpeningParenthesis, Val: "("},
		// 		{Kind: lexer.OperatorToken, Val: "-"},
		// 		{Kind: lexer.OperandToken, Val: "12"},
		// 		{Kind: lexer.OperandToken, Val: "1"},
		// 		{Kind: lexer.ClosedParenthesis, Val: ")"},
		// 		{Kind: lexer.ClosedParenthesis, Val: ")"},
		// 	},
		// },
		{
			tokens: []lexer.Token{
				{Kind: lexer.OpeningParenthesis, Val: "("},
				{Kind: lexer.OperatorToken, Val: "+"},
				{Kind: lexer.OperandToken, Val: "3"},
				{Kind: lexer.OperandToken, Val: "2"},
				{Kind: lexer.ClosedParenthesis, Val: ")"},
			},
			output: AST{
				{
					Kind:    LiteralNode,
					Literal: &lexer.Token{Kind: lexer.OperatorToken, Val: "+"},
				},
				{
					Kind:    LiteralNode,
					Literal: &lexer.Token{Kind: lexer.OperandToken, Val: "3"},
				},
				{
					Kind:    LiteralNode,
					Literal: &lexer.Token{Kind: lexer.OperandToken, Val: "2"},
				},
			},
		},
	}

	for _, testCase := range tests {
		ast, _ := Parse(testCase.tokens, 0)
		t.Logf(" ➜ AST : %v", ast)
		t.Logf(" ➜ AST : %v", testCase.output)
		assert.Equal(t, testCase.output, ast)
	}
}

/*
[
	{OPENING_PARENTHESIS_TOKEN (}
	{OPERATOR_TOKEN +}
	{OPERAND_TOKEN 13}
	{OPENING_PARENTHESIS_TOKEN (}
	{OPERATOR_TOKEN -}
	{OPERAND_TOKEN 12}
	{OPERAND_TOKEN 1}
	{CLOSED_PARENTHESIS_TOKEN )}
	{CLOSED_PARENTHESIS_TOKEN )}]

*/
