package main

import (
	p "expert-system/parse"
	s "expert-system/solve"
	"fmt"
)

func main() {
	// tmp := make(map[byte]bool)
	// tmp['A'] = true
	// fmt.Println(s.RuleIsTrue("A", tmp))
	// return
	ctx, err := p.Parse()
	if err != "" {
		fmt.Println("Error:", err)
	} else {
		err = s.Algo(ctx)
		if err != "" {
			fmt.Println(err)
		} else {
			if ctx.Flag_v {
				fmt.Print(ctx.Verbose)
			}
			for _, char := range ctx.Query {
				fmt.Println(string(char) + ":", ctx.Variables[char])
			}
		}
	}
}
