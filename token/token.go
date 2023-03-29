package token

//定义词法单元

//类型属性
type TokenType string

type Token struct {
  Type TokenType
  Literal string   //保存字面量
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  IDENT = "IDENT"
  INT = "INT"

  ASSIGN = "="
  PLUS = "+"

  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  FUNCTION = "FUNCTION"
  LET = "LET"
)
