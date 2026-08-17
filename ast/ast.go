package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

// Statement 语句接口, 不会返回值
type Statement interface {
	Node
	statementNode()
}

// Expression 表达式接口, 会返回值
type Expression interface {
	Node
	expressionNode()
}

// 主程序，由多个语句组成
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// ========== Let ==========
// let a = 5;
type LetStatement struct {
	Token token.Token // token.LET let
	Name  *Identifier // 标识符  a
	Value Expression  // let 语言产生值的表达式
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// let a = 5;  a 是 Indentifier 结构类型
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// ========== Return ==========
// return 4;  return add(x, y)
type ReturnStatement struct {
	Token       token.Token // token.RETURN return
	ReturnValue Expression  // 返回的值的表达式
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
