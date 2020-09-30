package command

type Router struct {
	Commands []*Command
	Prefix   string
}
