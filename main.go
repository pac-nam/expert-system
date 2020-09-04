package main

import (
	"fmt"
	p "expert-system/parse"
)

func main() {
	if err := p.Parse(); err != "" {
		fmt.Println("Error:", err)
	}
}