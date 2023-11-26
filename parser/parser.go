package parser

import "github.com/FadyGamilM/gosli/lexer"

/*
! How i represent the AST
(+ 13 (- 12 1))
=========== LEXER ============
[
	{kind : operator, val : "+"}
	{kind : operand, val : "13"}
	{kind : operator, val : "-"}
	{kind : operand, val : "12"}
	{kind : operand, val : "1"}
]

	+
  /  \
13    -
    /  \
   12   1

ast = [] value
ast = [] value :
[
	value :
	{
		kind : literal,
		literal : "+"
	}
	value :
	{
		kind : literal,
		literal : "13"
	}
	value :
	{
		kind : list,
		list : [] value {
			value : {
				kind : literal,
				literal : "-"
			}
			value : {
				kind : literal,
				literal : "12"
			}
			value : {
				kind : literal,
				literal : "1"
			}
		}
	}
]

*/

type NodeKind string

const (
	LiteralNode = "LITERAL_NODE"
	ListNode    = "LIST_NODE"
)

type AST []Node

func NewAST() AST {
	return []Node{}
}

type Node struct {
	Kind    NodeKind
	Literal *lexer.Token
	List    *AST
}

func Parse(tokens []lexer.Token, cursor int) (AST, int) {
	// if we are not gonna face a recursion case, this will be the one and only AST that we will return
	// if we will face a recursion, so this is the nested ast, we have to create a new one here so we will append all the subsequent nodes into it and then return it to the first (parent) AST we have created
	ast := NewAST()
	// we faced a "(" so we have to pass this opening parenthesis and apply our logic on the next token, thats why i increased the cursor
	cursor++
	for cursor < len(tokens) {
		token := tokens[cursor]
		// if its "(" => we have to create a new AST and append the tokens into it and then get back and append this AST into the original parent AST
		if token.Kind == lexer.OpeningParenthesis {
			childAST, nextCursor := Parse(tokens, cursor)
			// now append the childAST into the parent AST
			ast = append(ast, Node{Kind: ListNode, List: &childAST})
			// update the cursor to go to the next index after the curr cursor which was ")"
			cursor = nextCursor
			// now continue to go to the next iteration and skip checking the other conditions
			continue
		}

		if token.Kind == lexer.ClosedParenthesis {
			// now we are closing the current child AST we were working on, so we have to return the child ast and the cursor pointing to the next token after the ")" because we are here due to hitting a ")"
			return ast, cursor + 1
		}

		if token.Kind == lexer.OperandToken || token.Kind == lexer.OperatorToken {
			ast = append(ast, Node{Kind: LiteralNode, Literal: &token})
			//  update the cursor to go the next iteration with the next token
			cursor++
		}
	}
	return ast, -1
}
