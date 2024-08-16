package token

type Type int

const (
	LEFT_PAREN  Type = iota
	RIGHT_PAREN Type = iota
	LEFT_BRACE  Type = iota
	RIGHT_BRACE Type = iota

	COMMA     Type = iota
	DOT       Type = iota
	MINUS     Type = iota
	PLUS      Type = iota
	SEMICOLON Type = iota
	SLASH     Type = iota
	STAR      Type = iota

	BANG          Type = iota
	BANG_EQUAL    Type = iota
	EQUAL         Type = iota
	EQUAL_EQUAL   Type = iota
	GREATER       Type = iota
	GREATER_EQUAL Type = iota
	LESS          Type = iota
	LESS_EQUAL    Type = iota

	IDENTIFIER Type = iota
	STRING     Type = iota
	NUMBER     Type = iota

	AND    Type = iota
	CLASS  Type = iota
	ELSE   Type = iota
	FALSE  Type = iota
	FUN    Type = iota
	FOR    Type = iota
	IF     Type = iota
	NIL    Type = iota
	OR     Type = iota
	PRINT  Type = iota
	RETURN Type = iota
	SUPER  Type = iota
	THIS   Type = iota
	TRUE   Type = iota
	VAR    Type = iota
	WHILE  Type = iota
	EOF    Type = iota
)

func (t Type) String() string {
	switch t {
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case COMMA:
		return "COMMA"
	case DOT:
		return "DOT"
	case MINUS:
		return "MINUS"
	case PLUS:
		return "PLUS"
	case SEMICOLON:
		return "SEMICOLON"
	case SLASH:
		return "SLASH"
	case STAR:
		return "STAR"
	case BANG:
		return "BANG"
	case BANG_EQUAL:
		return "BANG_EQUAL"
	case EQUAL:
		return "EQUAL"
	case EQUAL_EQUAL:
		return "EQUAL_EQUAL"
	case GREATER:
		return "GREATER"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case LESS:
		return "LESS"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case IDENTIFIER:
		return "IDENTIFIER"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case AND:
		return "AND"
	case CLASS:
		return "CLASS"
	case ELSE:
		return "ELSE"
	case FALSE:
		return "FALSE"
	case FUN:
		return "FUN"
	case FOR:
		return "FOR"
	case IF:
		return "IF"
	case NIL:
		return "NIL"
	case OR:
		return "OR"
	case PRINT:
		return "PRINT"
	case RETURN:
		return "RETURN"
	case SUPER:
		return "SUPER"
	case THIS:
		return "THIS"
	case TRUE:
		return "TRUE"
	case VAR:
		return "VAR"
	case WHILE:
		return "WHILE"
	case EOF:
		return "EOF"
	default:
		return ""
	}
}

/*var TYPE_TO_LEXEME = map[Type]string{
	LEFT_PAREN:    "(",
	RIGHT_PAREN:   ")",
	LEFT_BRACE:    "{",
	RIGHT_BRACE:   "}",
	COMMA:         ",",
	DOT:           ".",
	MINUS:         "-",
	PLUS:          "+",
	SEMICOLON:     ";",
	SLASH:         "/",
	STAR:          "*",
	BANG:          "!",
	BANG_EQUAL:    "!=",
	EQUAL:         "=",
	EQUAL_EQUAL:   "==",
	GREATER:       ">",
	GREATER_EQUAL: ">=",
	LESS:          "<",
	LESS_EQUAL:    "<=",
	//IDENTIFIER:    "IDENTIFIER",
	//STRING:        "STRING",
	//NUMBER:        "NUMBER",
	//AND:           "&&",
	CLASS: "class",
	ELSE:  "else",
	FALSE: "false",
	//FUN:           "func",
	FOR: "for",
	IF:  "if",
	NIL: "nil",
	//OR:            "||",
	PRINT:  "print",
	RETURN: "return",
	SUPER:  "super",
	THIS:   "this",
	TRUE:   "true",
	VAR:    "var",
	WHILE:  "while"}

func GetLexemeByType(tokenType Type) string {
	return TYPE_TO_LEXEME[tokenType]
}

func GetTypeByLexeme(lexeme string) Type {
	var result Type = -1
	for key, value := range TYPE_TO_LEXEME {
		if value == lexeme {
			result = key
			break
		}
	}
	return result
}
*/
