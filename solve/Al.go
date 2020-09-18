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
		if exist && variable != asked {
			return "Error"
		}
		ctx.Variables[letter] = asked
	}
	return ""
}

func ComplexeCase(ctx *s.Context, retByter byte) (string, bool){
	for i, rule := range(ctx.Rules) {
		isTrue := RuleIsTrue(string(rule.Premice), ctx.Variables)
		if !rule.Used && isTrue == s.UNKNOW {
			for _, byter := range(rule.UsedVar) {
				if byter != retByter {
					ctx.Variables[byter] = false
					ComplexeCase(ctx, byter)
				}
			}
		} else if !rule.Used && isTrue == s.TRUE {
			if VerifConclusion(ctx, string(rule.Conclusion)) != "" {
				ctx.Rules[i].Used = false
				delete(ctx.Variables, retByter)
				ComplexeCase(ctx, retByter)
			} else {
				ctx.Rules[i].Used = true
				return "", true
			}
		} else if !rule.Used && isTrue == s.FALSE {
			ctx.Rules[i].Used = true
		}
	}
	return "", false
}

func VerifRules(ctx *s.Context) (string, bool) {
	for i, rule := range(ctx.Rules) {
		if !rule.Used && RuleIsTrue(string(rule.Premice), ctx.Variables) == s.TRUE {
			ctx.Rules[i].Used = true
			if VerifConclusion(ctx, string(rule.Conclusion)) != "" {
				return "Logical error with rule : "+fmt.Sprint(rule), false
			}
			return "", true
		} else if !rule.Used && RuleIsTrue(string(rule.Premice), ctx.Variables) == s.FALSE {
			ctx.Rules[i].Used = true
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
	ComplexeCase(ctx, '0')
	return ""
}