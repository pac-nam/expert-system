package solve

import (
	s "expert-system/structures"
	"strings"
)

func RemoveIndex(tab []rune, i int) []rune {
	if i == len(tab) - 1 {
		if i == 0 {
			return []rune{}
		}
		return tab[:i]
	}
	return append(tab[:i], tab[i+1:]...)
}

func CleanCanChange(tab []rune, variables map[rune]bool) []rune {
	for i := 0; i < len(tab); i++ {
		if _, exist := variables[tab[i]]; exist {
			tab = RemoveIndex(tab, i)
			i--
		}
	}
	return tab
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