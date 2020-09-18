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

func ReturnNewByte(ctx *s.Context, retByter byte ) byte{
	for _, rule := range(ctx.Rules) {
		isTrue := RuleIsTrue(string(rule.Premice), ctx.Variables)
		if !rule.Used && isTrue == s.UNKNOW {
			for _, byter := range(rule.UsedVar) {
				fmt.Println("byte", string(byter))
				if byter != retByter {
					ctx.Variables[byter] = false
					fmt.Println(string(byter), "devient faux")
					return (byter)
				}
			}
		}
	}
	return('1')
}

func ComplexeCase(ctx *s.Context, retByter byte) (string, bool){
	newByte := ReturnNewByte(ctx, retByter)
	if newByte == '1' {
		fmt.Println("C'est la fin")
		return "", false
	}
	unknow_rule := false
	for i, rule := range(ctx.Rules) {
		for key, value := range ctx.Variables {
			fmt.Println(string(key), value)
		}
		isTrue := RuleIsTrue(string(rule.Premice), ctx.Variables)
		if !rule.Used && isTrue == s.TRUE {
			if VerifConclusion(ctx, string(rule.Conclusion)) != "" {
				fmt.Println("Fail conclusion", string(newByte))
				ctx.Rules[i].Used = false
				delete(ctx.Variables, newByte)
				ComplexeCase(ctx, newByte)
			} else {
				fmt.Println("Ok")
				fmt.Println(rule)
				ctx.Rules[i].Used = true
			}
		} else if !rule.Used && isTrue == s.FALSE {
			fmt.Println("False")
			fmt.Println(rule)
			ctx.Rules[i].Used = true
		} else if !rule.Used && isTrue == s.UNKNOW {
			unknow_rule = true
			fmt.Println("Unknow", isTrue, rule)
		}
		fmt.Println()
	}
	if unknow_rule == true {
		fmt.Println("Et on recommence")
		ComplexeCase(ctx, newByte)
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