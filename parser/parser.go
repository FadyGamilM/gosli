package parser

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

*/