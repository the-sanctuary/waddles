package parser

// Parser object to hold data of the command parser
type Parser struct {
	StringDelim string
}

// BuildParser constructor-like factory function for the Parser object
func BuildParser(stringDelim string) Parser {
	p := Parser{
		StringDelim: stringDelim,
	}
	return p
}

// Parse parses a command string into a command tree
func (p *Parser) Parse(rawCmd string, tree *CmdTree) {
	c1 := BuildCmdTreeArg("hello")
	c2 := BuildCmdTreeArg("world")
	tree.Data = rawCmd
	tree.Trees = []*CmdTree{c1, c2}
}
