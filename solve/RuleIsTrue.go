package solve

import (
	"strings"
	// "fmt"
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

func simpleVar(Conditions string, Variables map[rune]bool) int {
	length := len(Conditions)

	if length == 2 {
		switch Conditions[1] {
		case 't':
			return s.FALSE
		case 'f':
			return s.TRUE
		case 'u':
			return s.UNKNOW
		default:
			variable, exist := Variables[rune(Conditions[1])]
			return retBoolSprecial(!variable, exist)
		}
	} else if length == 1 {
		switch Conditions[0] {
		case 't':
			return s.TRUE
		case 'f':
			return s.FALSE
		case 'u':
			return s.UNKNOW
		default:
			variable, exist := Variables[rune(Conditions[0])]
			return retBoolSprecial(variable, exist)
		}
	}
	panic("error: read while reading rule: '" + Conditions + "'")
}

func andRule(Conditions string, Variables map[rune]bool) int {
	tab := strings.Split(Conditions, "+")
	res := s.TRUE
	for _, part := range tab {
		tmp := simpleVar(part, Variables)
		if tmp == s.FALSE {
			return s.FALSE
		} else if tmp == s.UNKNOW {
			res = s.UNKNOW
		}
	}
	return res
}

func orRule(Conditions string, Variables map[rune]bool) int {
	tab := strings.Split(Conditions, "|")
	res := s.FALSE
	for _, part := range tab {
		tmp := andRule(part, Variables)
		if tmp == s.TRUE {
			return s.TRUE
		} else if tmp == s.UNKNOW {
			res = s.UNKNOW
		}
	}
	return res
}

func xorRule(Conditions string, Variables map[rune]bool) (res int, residue string) {
	tab := strings.SplitN(Conditions, ")", 2)
	if len(tab) == 2 {
		residue = tab[1]
	}
	tab = strings.Split(tab[0], "^")
	res = andRule(tab[0], Variables)
	if res == s.UNKNOW {
		return
	}
	for i := 1; i < len(tab); i++ {
		tmp := andRule(tab[i], Variables)
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

func RuleIsTrue(Conditions string, Variables map[rune]bool) int {
	tab := strings.Split(Conditions, "(")
	for length := len(tab) - 1; length > 1; length--{
		boolRes, residue := xorRule(tab[length], Variables)
		if boolRes == s.TRUE {
			tab[length-1] += "t" + residue
		} else if boolRes == s.FALSE {
			tab[length-1] += "f" + residue
		} else {
			tab[length-1] += "u" + residue
		}
		tab = tab[:length]
	}
	finalResult, _ := xorRule(tab[0], Variables)
	return finalResult
}
