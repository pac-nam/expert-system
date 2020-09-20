package solve

import (
	s "expert-system/structures"
	m "expert-system/messages"
	// "strings"
	// "fmt"
)

func VerifConclusion(ctx *s.Context, Conclusion string, CanChange []rune) bool {
	if len(CanChange) == 0 {
		// fmt.Println("Conclusion:", Conclusion)
		if RuleIsTrue(Conclusion, ctx.Variables) == s.TRUE {
			newctx := ctx.Copy()
			finalmap, end := ComplexeCase(newctx)
			if end {
				ctx.Rules = newctx.Rules
				ctx.Variables = finalmap
				return true
			}
		}
	} else if _, exist := ctx.Variables[CanChange[0]]; !exist {
		ctx.Variables[CanChange[0]] = true
		if VerifConclusion(ctx, Conclusion, CanChange[1:]) {
			return true
		}
		ctx.Variables[CanChange[0]] = false
		if VerifConclusion(ctx, Conclusion, CanChange[1:]) {
			return true
		}
		delete(ctx.Variables, CanChange[0])
	} else {
		if VerifConclusion(ctx, Conclusion, CanChange[1:]) {
			return true
		}
	}
	return false
}

func ComplexeCase(ctx *s.Context) (finalMap map[rune]bool, end bool) {
	if !VerifRules(ctx) {
		return ctx.Variables, false
	}
	if len(ctx.Rules) == 0 {
		return ctx.Variables, true
	}
	CleanCanChange(ctx)
	for i, char := range ctx.CanChange {
		ctx.Variables[char] = false
		newctx := ctx.Copy()
		newctx.CanChange = RemoveIndex(ctx.CanChange, i)
		finalMap, end := ComplexeCase(newctx)
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
			if !VerifConclusion(ctx, Conclusion, []rune(UsedVar(Conclusion))) {
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
	ctx.CanChange = ChangableVariables(ctx)
	// fmt.Println(string(ctx.CanChange))
	finalMap, end := ComplexeCase(ctx)
	ctx.Variables = finalMap
	if !end {
		return m.Impossible
	}
	return ""
}