package parser

// Parser object to hold data of the command parser
type Parser struct {
	StringDelims []string
	LRDelims     []string
}

// BuildParser constructor-like factory function for the Parser object
func BuildParser(stringDelims []string) Parser {
	p := Parser{}
	for _, sd := range stringDelims {
		if len(sd) == 1 {
			p.StringDelims = append(p.StringDelims, sd)
		} else if len(sd) == 2 {
			p.LRDelims = append(p.LRDelims, string(sd[0]))
			p.LRDelims = append(p.LRDelims, string(sd[1]))
		}
	}
	return p
}

// Parse parses a command string into a command tree
func (p *Parser) Parse(raw string, tree *CmdTree) {
	c1 := BuildCmdTreeArg("hello")
	c2 := BuildCmdTreeArg("world")
	tree.Data = raw
	tree.Trees = []*CmdTree{c1, c2}
}
