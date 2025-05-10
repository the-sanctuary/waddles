package parser

import "github.com/rs/zerolog/log"

// CmdTree tree object for holding a parsed command string
type CmdTree struct {
	Data  string
	Trees []*CmdTree
}

// BuildCmdTree constructor-like factory function for the CmdTree object
func BuildCmdTree() *CmdTree {
	tree := &CmdTree{
		Data:  "",
		Trees: []*CmdTree{},
	}
	return tree
}

// BuildCmdTreeArg constructor-like factory function for the CmdTree object and passing in data
func BuildCmdTreeArg(data string) *CmdTree {
	tree := &CmdTree{
		Data:  data,
		Trees: []*CmdTree{},
	}
	return tree
}

// LRTraverse traverses a CmdTree object from bottom up, left to right
func (t *CmdTree) LRTraverse() {
	if len(t.Trees) == 0 {
		log.Trace().Msgf("LRTraverse> %s", t.Data)
		return
	}
	for _, tree := range t.Trees {
		tree.LRTraverse()
	}
	log.Trace().Msgf("LRTraverse> %s", t.Data)
}
