package structures

import (
	"fmt"
)

type Context struct {
	Flag_v		bool
	Flag_a		bool
	Verbose		string
	DeepLevel	int
	Rules		[]Rule
	Initial		string
	Query		string
	CanChange	[]rune
	Variables	map[rune]bool
}

func (ctx Context) String() string {
	res := "------------------------------------Context------------------------------------\n"
	if ctx.Initial != "" {
		res += "Initial state:\n"
		res += fmt.Sprintln(ctx.Initial)
	}
	if ctx.Query != "" {
		res += "\nQuery:\n"
		res += fmt.Sprintln(ctx.Query)
	}
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

func (ctx *Context) Copy() *Context {
	newCtx := Context{DeepLevel: ctx.DeepLevel+1}
	newCtx.Rules = make([]Rule, len(ctx.Rules))
	copy(newCtx.Rules, ctx.Rules)
	newCtx.CanChange = make([]rune, len(ctx.CanChange))
	copy(newCtx.CanChange, ctx.CanChange)
	newCtx.Variables = make(map[rune]bool)
	for key, value := range ctx.Variables {
		newCtx.Variables[key] = value
	}
	return &newCtx
}

func (ctx *Context) RemoveRule(i int) {
	if i == len(ctx.Rules) - 1 {
		ctx.Rules = ctx.Rules[:i]
	} else {
		ctx.Rules = append(ctx.Rules[:i], ctx.Rules[i+1:]...)
	}
}