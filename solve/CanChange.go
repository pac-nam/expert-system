package solve

import (
	s "expert-system/structures"
	"strings"
	"fmt"
)

func RemoveIndex(tab []rune, i int) []rune {
	if i == len(tab) - 1 {
		if i == 0 {
			return []rune{}
		}
		return tab[:i]
	}
	fmt.Println(len(tab), i)
	return append(tab[:i], tab[i+1:]...)
}

func CleanCanChange(ctx *s.Context) {
	for i := 0; i < len(ctx.CanChange); i++ {
		if _, exist := ctx.Variables[ctx.CanChange[i]]; exist {
			// fmt.Println(string(ctx.CanChange))
			ctx.CanChange = RemoveIndex(ctx.CanChange, i)
			i--
		}
	}
}

func UsedVar(rule string) (res string) {
	for _, char := range rule {
		if strings.ContainsRune(ALPHABET, char) {
			if !strings.ContainsRune(res, char) {
				res += string(char)
			}
		}
	}
	return
}

func ChangableVariables(ctx *s.Context) []rune {
	res := ""
	for _, rule := range ctx.Rules {
		for _, char := range UsedVar(rule.Premice) {
			if _, exist := ctx.Variables[char]; !exist && !strings.ContainsRune(res, char) {
				res += string(char)
			}
		}
	}
	return []rune(res)
}