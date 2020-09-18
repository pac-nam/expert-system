package solve

import (
	"strings"
	"fmt"
	// "os"
	s "expert-system/structures"
)

const (
	ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func retBoolSprecial(variable bool, exist bool) int {
	if !exist {
		return s.UNKNOW
	} else if variable {
		return s.TRUE
	}
	return s.FALSE
}

func noAnd(Conditions string, Variables map[byte]bool) int {
	length := len(Conditions)

	if length == 2 && Conditions[0] == '!' {
		if strings.ContainsRune(ALPHABET, rune(Conditions[1])) {
			variable, exist := Variables[byte(Conditions[1])]
			return retBoolSprecial(!variable, exist)
		} else if Conditions[1] == 't' {
			return s.FALSE
		} else if Conditions[1] == 'f' {
			return s.TRUE
		} else if Conditions[1] == 'u' {
			return s.UNKNOW
		}
	} else if length == 1 {
		if strings.ContainsRune(ALPHABET, rune(Conditions[0])) {
			variable, exist := Variables[byte(Conditions[0])]
			return retBoolSprecial(variable, exist)
		} else if Conditions[0] == 't' {
			return s.TRUE
		} else if Conditions[0] == 'f' {
			return s.FALSE
		} else if Conditions[0] == 'u' {
			return s.UNKNOW
		}
	}
	panic("error: read while reading rule: '" + Conditions + "'")
}

func noOr(Conditions string, Variables map[byte]bool) int {
	tab := strings.Split(Conditions, "+")
	res := s.TRUE
	for _, part := range tab {
		tmp := noAnd(part, Variables)
		// fmt.Println(tmp)
		if tmp == s.FALSE {
			return s.FALSE
		} else if tmp == s.UNKNOW {
			res = s.UNKNOW
		}
	}
	return res
}

func noXor(Conditions string, Variables map[byte]bool) int {
	tab := strings.Split(Conditions, "|")
	res := s.FALSE
	for _, part := range tab {
		tmp := noOr(part, Variables)
		if tmp == s.TRUE {
			return s.TRUE
		} else if tmp == s.UNKNOW {
			res = s.UNKNOW
		}
	}
	return res
}

func noParenthesisRule(Conditions string, Variables map[byte]bool) (res int, residue string) {
	tab := strings.SplitN(Conditions, ")", 2)
	// fmt.Println(tab[0])
	// os.Exit(0)
	if len(tab) == 2 {
		residue = tab[1]
	}
	tab = strings.Split(tab[0], "^")
	res = noXor(tab[0], Variables)
	if res == s.UNKNOW {
		return
	}
	for i := 1; i < len(tab); i++ {
		tmp := noXor(tab[i], Variables)
		if tmp == s.UNKNOW {
			return s.UNKNOW, residue
		} else if tmp != res {
			res = s.TRUE
		} else {
			res = s.FALSE
		}
	}
	return
}

func RuleIsTrue(Conditions string, Variables map[byte]bool) int {
	tab := strings.Split(Conditions, "(")
	for length := len(tab); length > 1; {
		length--
		// for _, str := range tab {
		// 	fmt.Print(" '" + str + "'")
		// }
		// fmt.Println()
		boolRes, residue := noParenthesisRule(tab[length], Variables)
		if boolRes == s.TRUE {
			tab[length-1] += "t" + residue
		} else if boolRes == s.FALSE {
			tab[length-1] += "f" + residue
		} else {
			tab[length-1] += "u" + residue
		}
		tab = tab[:length]
	}
	finalResult, _ := noParenthesisRule(tab[0], Variables)
	return finalResult
}
