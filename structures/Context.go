package structures

import (
	"fmt"
)

type Context struct {
	Rules		[]Rule
	Initial		[]byte
	Query		[]byte
	Variables	map[byte]bool
}

func (ctx Context) String() string {
	res := "------------------------------------Context------------------------------------\n"
	res += "Initial state:\n"
	res += fmt.Sprintln(string(ctx.Initial))
	res += "\nQuery:\n"
	res += fmt.Sprintln(string(ctx.Query))
	res += "\nRules:\n"
	for _, rule := range ctx.Rules {
		res += fmt.Sprintln(rule)
	}
	res += "Variables:\n"
	for key, Value := range ctx.Variables {
		res += fmt.Sprintln(string(key)+":", Value)
	}
	return res
}
