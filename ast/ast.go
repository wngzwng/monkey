package ast

import (
	"bytes"

	"monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string // 添加 String() 方法，方便后续调试时打印AST节点，也可以用来比较AST节点
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (ls *LetStatement) String() string {
	var out bytes.Buffer 

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
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

func (i *Identifier) String() string {
	return i.Value
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

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer 

	out.WriteString(rs.TokenLiteral() + " ")
	
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}


// ========== Expression ==========
// x + 10; 只包含表达式的单行语句
type ExpressionStatement struct {
	Token token.Token // 该表达式中的第一个词法单元
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}


// ========== 整数字面量 ==========
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}


// ========== 前缀表达式 ==========
type PrefixExpression struct {
	Token		token.Token	// 前缀词法单元，如 ！
	Operator	string
	Right		Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}



