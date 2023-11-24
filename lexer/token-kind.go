package lexer

// Enum represents different types of tokens
type TokenKind uint

const (
	// represents the variable names
	// for example : let {count} = 5, the count is an identifier token
	IdentifierToken TokenKind = iota

	// represents the reserved keywords in the language
	// for example : for, if, var, const, return, break, continue ....
	ReservedKeyworkdToken

	// represents the data types such as int, float, ...
	DataTypeToken

	// represents operators
	// for example : +, -, *, /, &&, &, ||, |, ==, =, !=
	OperatorToken

	// represents the values assigned to variables
	// for example : var username = "fady gamil", the "fady gamil" is specified as a Literal token
	LiteralToken

	// represnet a comment
	// for example : //, /*, */
	CommentToken
)
