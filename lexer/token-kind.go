package lexer

// Enum represents different types of tokens
type TokenKind uint

const (

	// represents the variable names
	// for example : (define (add x) (+ x x))
	// {add} here is an idetifier
	// {+} here is an idetifier
	IdentifierToken TokenKind = iota

	// (
	OpeningParenthesis

	// )
	ClosedParenthesis

	// represents the reserved keywords in the language
	// for example : (define (add x) (+ x x))  {define} here is a reserved keyword
	ReservedKeyworkdToken

	// (+ 5 6) 5 or 6 are operands
	OperandToken

	// represents the data types such as int, float, ...
	DataTypeToken

	// represents operators
	// for example : +, -, *, /, &&, &, ||, |, ==, =, !=
	OperatorToken

	// represents the values assigned to variables
	// for example : (display "Enter two numbers: ") {display} is a literal string
	LiteralToken

	// represnet a comment
	// for example : //, /*, */
	CommentToken
)
