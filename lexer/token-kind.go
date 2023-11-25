package lexer

// Enum represents different types of tokens
type TokenKind string

const (

	// represents the variable names
	// for example : (define (add x) (+ x x))
	// {add} here is an idetifier
	// {+} here is an idetifier
	IdentifierToken TokenKind = "IDENTIFIER_TOKEN"

	// (
	OpeningParenthesis TokenKind = "OPENING_PARENTHESIS_TOKEN"

	// )
	ClosedParenthesis TokenKind = "CLOSED_PARENTHESIS_TOKEN"

	// represents the reserved keywords in the language
	// for example : (define (add x) (+ x x))  {define} here is a reserved keyword
	ReservedKeyworkdToken TokenKind = "REVERSED_KEYWORD_TOKEN"

	// (+ 5 6) 5 and 6 are operands
	// (+ x y) x and y are opearnds too if they are prev defined
	OperandToken TokenKind = "OPERAND_TOKEN"

	// represents the data types such as {list}
	DataTypeToken TokenKind = "DATA_TYPE_TOKEN"

	// represents operators
	// for example : +, -, *, /, &&, &, ||, |, ==, =, !=
	OperatorToken TokenKind = "OPERATOR_TOKEN"

	// represents the values assigned to variables
	// for example : (display "Enter two numbers: ") {display} is a literal string
	LiteralToken TokenKind = "LITERAL_TOKEN"

	// represnet a comment
	// for example : //, /*, */
	CommentToken TokenKind = "COMMENT_TOKEN"
)
