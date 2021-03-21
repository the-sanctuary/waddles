package parser

type Lexer struct {
	Lexemes map[string][]string
}

// BuildLexer returns a Lexer with the proper keyowrds and tokens
func BuildLexer() Lexer {
	lex := Lexer{}

	// Define the lexemes for our lexer
	lex.Lexemes = map[string][]string{
		"<cmd>": {""},
	}
	return lex
}

func addAllCommands()
