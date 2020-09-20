package solve

import (
	s "expert-system/structures"
	m "expert-system/messages"
	"strings"
)

func VerifConclusion(ctx *s.Context, Conclusion string) (success bool) {
	tab := strings.Split(Conclusion, "+")
	var variable, exist, asked bool
	var letter rune
	for _, part := range tab {
		if len(part) == 1 {
			letter = rune(part[0])
			asked = true
		} else {
			letter = rune(part[1])
			asked = false
		}
		variable, exist = ctx.Variables[letter]
		if exist && variable != asked {
			return false
		}
		ctx.Variables[letter] = asked
	}
	return true
}

func ComplexeCase(ctx *s.Context, CanChange []rune) (finalMap map[rune]bool, end bool) {
	if !VerifRules(ctx) {
		return ctx.Variables, false
	}
	if len(ctx.Rules) == 0 {
		return ctx.Variables, true
	}
	CanChange = CleanCanChange(CanChange, ctx.Variables)
	for i, char := range CanChange {
		ctx.Variables[char] = false
		finalMap, end := ComplexeCase(ctx.Copy(), RemoveIndex(CanChange, i))
		if end {
			return finalMap, true
		}
		delete(ctx.Variables, char)
	}
	return ctx.Variables, false
}

func VerifRules(ctx *s.Context) (success bool) {
	VerifRulesStart:
	for i:= 0; i < len(ctx.Rules); {
		switch RuleIsTrue(ctx.Rules[i].Premice, ctx.Variables) {
		case s.TRUE:
			Conclusion := ctx.Rules[i].Conclusion
			ctx.RemoveRule(i)
			if !VerifConclusion(ctx, Conclusion) {
				return false
			}
			goto VerifRulesStart
		case s.FALSE:
			ctx.RemoveRule(i)
		default:
			i++
		}
	}
	return true
}

func Algo(ctx *s.Context) string {
	finalMap, end := ComplexeCase(ctx, ChangableVariables(ctx))
	ctx.Variables = finalMap
	if !end {
		return m.Impossible
	}
	return ""
}