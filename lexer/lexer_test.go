package lexer

import "banana/token"

//词法分析器： 输入源代码，输出对应的词法单元

func TestNextTok(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, v := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, v.expectedType, tok.Type)
		}

		if tok.Type != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, v.expectedLiteral, tok.Literal)
		}

	}
}
