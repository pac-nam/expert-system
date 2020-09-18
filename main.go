package main

import (
	p "expert-system/parse"
	s "expert-system/solve"
	"fmt"
)

func main() {
	ctx, err := p.Parse()
	if err != "" {
		fmt.Println("Error:", err)
	} else {
		err = s.Algo(ctx)
		if err != "" {
			fmt.Println(err)
		} else {
			fmt.Println(ctx)
		}
	}
}
