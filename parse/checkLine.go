package parse

import (
	"strings"
)

const (
	// ALPHABET is alphabet
	ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//ALPHABETP is alphabet with exclamation mark
	ALPHABETP = ALPHABET + "!"
)

func createCheckMap() map[rune]string {
	res := make(map[rune]string)
	res['('] = ALPHABETP
	res[')'] = ALPHABETP
	res['<'] = "="
	res['='] = ">"
	res['>'] = ALPHABETP
	res['!'] = ALPHABET
	res['+'] = ALPHABETP
	res['|'] = ALPHABETP
	res['^'] = ALPHABETP
	return res
}

func checkLine(line string) string {
	equal := false
	lastVerif := len(line) - 1
	bracket := 0
	if !strings.ContainsRune(ALPHABET, rune(line[lastVerif])) {
		return "error with rule: " + line
	}
	for index, char := range line[:lastVerif] {
		if strings.ContainsRune(ALPHABET, char) {
			if !strings.ContainsRune("<+=^|)", rune(line[index+1])) {
				return "wrong character after '" + string(char) + "' in rule: " + line
			}
		} else if char == '=' {
			if equal {
				return "error multiple sign '=' in rule: " + line
			} else if line[index+1] != '>' {
				return "missing '>' after '=' in rule: " + line
			} else if bracket != 0 {
				return "parenthesis error in rule: " + line
			}
			equal = true
		} else if char == '<' {
			if line[index+1] != '=' {
				return "wrong character after '" + string(char) + "' in rule: " + line
			}
		} else if char == '>' {
			if !strings.ContainsRune(ALPHABETP, rune(line[index+1])) {
				return "wrong character after '" + string(char) + "' in rule: " + line
			}
		} else if char == '!' {
			if !strings.ContainsRune(ALPHABET+"(", rune(line[index+1])) {
				return "wrong character after '" + string(char) + "' in rule: " + line
			}
		} else if char == '+' {
			if !strings.ContainsRune(ALPHABET+"(!", rune(line[index+1])) {
				return "wrong character after '" + string(char) + "' in rule: " + line
			}
		} else if strings.ContainsRune("|^", char) {
			if !strings.ContainsRune(ALPHABET+"(!", rune(line[index+1])) {
				return "wrong character after '" + string(char) + "' in rule: " + line
			} else if equal {
				return "invalid sign '" + string(char) + "' in conclusion in rule: " + line
			}
		} else if char == '(' {
			if !strings.ContainsRune(ALPHABET+"!", rune(line[index+1])) {
				return "wrong character after '" + string(char) + "' in rule: " + line
			} else if equal {
				return "parenthesis in conclusion in rule: " + line
			}
			bracket++
		} else if char == ')' {
			if !strings.ContainsRune("+|^<=", rune(line[index+1])) {
				return "wrong character after '" + string(char) + "' in rule: " + line
			} else if bracket < 1 {
				return "parenthesis error in rule: " + line
			}
			bracket--
		}
	}
	return ""
}
