package solve

import (
	"fmt"
	s "expert-system/structures"
	"strings"
)

func VerifConclusion(ctx *s.Context, Conclusion string) string{
	tab := strings.Split(Conclusion, "+")
	var variable, exist, asked bool
	var letter byte
	for _, part := range tab {
		if len(part) == 1 {
			letter = part[0]
			asked = true
		} else {
			letter = part[1]
			asked = false
		}
		variable, exist = ctx.Variables[letter]
		if exist == true && variable != asked {
			return "Error"
		}
		ctx.Variables[letter] = asked
	}
	return ""
}

func VerifRules(ctx *s.Context) (string, bool) {
	for i, rule := range(ctx.Rules) {
		if !rule.Used && RuleIsTrue(string(rule.Premice), ctx.Variables) {
			ctx.Rules[i].Used = true
			if VerifConclusion(ctx, string(rule.Conclusion)) != "" {
				return "Logical error with rule : "+fmt.Sprint(rule), false
			}
			return "", true
		}
	}
	return "", false
}

func Algo(ctx *s.Context) string {
	activate := true
	var err string
	for activate {
		err, activate = VerifRules(ctx)
		if err != "" {
			return err
		}
	}
	return ""
}