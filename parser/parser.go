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
ast = [] value : {
	value : {
		kind : literal,
		literal : "+"
	}
	value : {
		kind : literal,
		literal : "13"
	}
	value : {
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
}

*/


type NodeKind string 
const (
	LiteralNode = "LITERAL_NODE"
	ListNode = "LIST_NODE"
)

type AST []Node

type Node struct {
	Kind NodeKind
	Literal *lexer.Token
	List *AST
}

func Parse (tokens []lexer.Token) *AST{
	
}