package main

import (
	"fmt"
	p "expert-system/parse"
	s "expert-system/solve"
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