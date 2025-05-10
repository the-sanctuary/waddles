package commands

import "github.com/the-sanctuary/waddles/pkg/cmd"

/* === top level command === */
var lex *cmd.Command = &cmd.Command{
	Name:        "lex",
	Aliases:     []string{""},
	Description: "Temporary command to test the lexer and parser (based on topics).",
	Usage:       "lex (add|list|remove|manage)",
	SubCommands: []*cmd.Command{lexAdd, lexList, lexRemove, lexManage},
}

/* === primary subcommands === */
var lexAdd *cmd.Command = &cmd.Command{
	Name:        "add",
	Aliases:     []string{""},
	Description: "Temporary command to test the lexer and parser (based on topics).",
	Usage:       "lex add <slug>",
	SubCommands: []*cmd.Command{},
}

var lexList *cmd.Command = &cmd.Command{
	Name:        "list",
	Aliases:     []string{""},
	Description: "Temporary command to test the lexer and parser (based on topics).",
	Usage:       "lex list",
	SubCommands: []*cmd.Command{},
}

var lexRemove *cmd.Command = &cmd.Command{
	Name:        "remove",
	Aliases:     []string{""},
	Description: "Temporary command to test the lexer and parser (based on topics).",
	Usage:       "lex remove <slug>",
	SubCommands: []*cmd.Command{},
}

var lexManage *cmd.Command = &cmd.Command{
	Name:        "manage",
	Aliases:     []string{""},
	Description: "Temporary command to test the lexer and parser (based on topics).",
	Usage:       "lex manage (add|remove|edit|tags)",
	SubCommands: []*cmd.Command{lexManageAdd, lexManageRemove},
}

/* === secondary subommands */
var lexManageAdd *cmd.Command = &cmd.Command{
	Name:        "add",
	Aliases:     []string{""},
	Description: "Temporary command to test the lexer and parser (based on topics).",
	Usage:       "lex manage add <slug> <name> [<tag>] <description>",
	SubCommands: []*cmd.Command{},
}

var lexManageRemove *cmd.Command = &cmd.Command{
	Name:        "remove",
	Aliases:     []string{""},
	Description: "Temporary command to test the lexer and parser (based on topics).",
	Usage:       "lex manage <slug>",
	SubCommands: []*cmd.Command{},
}

// var lexManageEdit *cmd.Command = &cmd.Command{
// 	Name:        "edit",
// 	Aliases:     []string{""},
// 	Description: "Temporary command to test the lexer and parser (based on topics).",
// 	Usage:       "lex manage edit",
// 	SubCommands: []*cmd.Command{},
// }

// var lexManageTags *cmd.Command = &cmd.Command{
// 	Name:        "tags",
// 	Aliases:     []string{""},
// 	Description: "Temporary command to test the lexer and parser (based on topics).",
// 	Usage:       "lex manage tags",
// 	SubCommands: []*cmd.Command{},
// }
