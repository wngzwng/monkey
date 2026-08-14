package token

type TokenType string 

// Token 词法单元结构体
type Token struct {
	Type TokenType 			// 区分不同类型的词法单元
	Literal string
}


// TokenType 常量定义
const (
	ILLEGAL = "ILLEGAL" // 非法字符
	EOF	    = "EOF"     // 文件结束符

	// 标识符 + 字面量
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// 运算符
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	BANG   = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ = "=="
	NOT_EQ = "!="

	// 分隔符
	COMMA 	  = ","
	SEMICOLON = ";"
	
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"


)


var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

// LookupIdent 判断标识符是否为关键字
func LookupIdent(ident string) TokenType {
	if tokeType, ok := keywords[ident]; ok {
		return tokeType
	}
	return IDENT
}
