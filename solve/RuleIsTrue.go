package solve

import (
	"strings"
	// "fmt"
	// "os"
)
const (
	ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func noAnd(Conditions string, Variables map[byte]bool) bool {
	length := len(Conditions)
	if length == 2 && Conditions[0] == '!'{
		if strings.ContainsRune(ALPHABET, rune(Conditions[1])) {
			return !Variables[Conditions[1]]
		} else if Conditions[1] == 't' {
			return false
		} else if Conditions[1] == 'f' {
			return true
		}
	} else if length == 1 {
		if strings.ContainsRune(ALPHABET, rune(Conditions[0])) {
			return Variables[Conditions[0]]
		} else if Conditions[0] == 't' {
			return true
		} else if Conditions[0] == 'f' {
			return false
		}
	}
	panic("error: read while reading rule: '" + Conditions + "'")
}

func noOr(Conditions string, Variables map[byte]bool) bool {
	tab := strings.Split(Conditions, "+")
	for _, part := range tab {
		if !noAnd(part, Variables) {
			return false
		}
	}
	return true
}

func noXor(Conditions string, Variables map[byte]bool) bool {
	tab := strings.Split(Conditions, "|")
	for _, part := range tab {
		if noOr(part, Variables) {
			return true
		}
	}
	return false
}

func noParenthesisRule(Conditions string, Variables map[byte]bool) (res bool, residue string) {
	tab := strings.SplitN(Conditions, ")", 2)
	// fmt.Println(tab[0])
	// os.Exit(0)
	if len(tab) == 2 {
		residue = tab[1]
	}
	tab = strings.Split(tab[0], "^")
	res = noXor(tab[0], Variables)
	for i := 1; i < len(tab); i++ {
		if res != noXor(tab[i], Variables) {
			res = true
		} else {
			res = false
		}
	}
	return
}

func RuleIsTrue(Conditions string, Variables map[byte]bool) bool {
	tab := strings.Split(Conditions, "(")
	for length := len(tab); length > 1; {
		length--
		// for _, str := range tab {
		// 	fmt.Print(" '" + str + "'")
		// }
		// fmt.Println()
		boolRes, residue := noParenthesisRule(tab[length], Variables)
		if boolRes {
			tab[length - 1] += "t" + residue
		} else {
			tab[length - 1] += "f" + residue
		}
		tab = tab[:length]
	}
	finalResult, _ := noParenthesisRule(tab[0], Variables)
	return finalResult
}