package solve

import (
	s "expert-system/structures"
	m "expert-system/messages"
	"fmt"
)

// VerifConclusion take four parameters
// ctx:			contain all informations about the program
// Conclusion:	is the right part of the rule, it need to be set to true to continue.
//				elseway, the rule is a logical error and backtrack is needed
// CanChange:	this string contain the variables of the Conclusion it permit a faster
//				acces to pass to true or false these variables
// changed:		this string contain the variables that have already been changed for
//				this conclusion. His only purpose is for the verbose.
// VerifConclusion return a boolean. true if the model is correct and false alseway.
// it can try different conclusion recursively. If after all it does not work, it can
// backtrack.
func VerifConclusion(ctx *s.Context, Conclusion string, CanChange []rune, changed string) bool {
	if len(CanChange) == 0 {
		if RuleIsTrue(Conclusion, ctx.Variables) == s.TRUE {
			ctx.Verbose += "conclusion:" + changed + "\n"
			newctx := ctx.Copy()
			finalctx, end := ComplexeCase(newctx)
			ctx.Verbose += finalctx.Verbose
			if end {
				ctx.Rules = newctx.Rules
				ctx.Variables = finalctx.Variables
				return true
			}
		}
	} else if _, exist := ctx.Variables[CanChange[0]]; !exist {
		ctx.Variables[CanChange[0]] = true
		if VerifConclusion(ctx, Conclusion, CanChange[1:], changed + " " + string(CanChange[0])) {
			return true
		}
		ctx.Variables[CanChange[0]] = false
		if VerifConclusion(ctx, Conclusion, CanChange[1:], changed + " !" + string(CanChange[0])) {
			return true
		}
		delete(ctx.Variables, CanChange[0])
	} else {
		if VerifConclusion(ctx, Conclusion, CanChange[1:], changed) {
			return true
		}
	}
	return false
}

// VerifRules take one parameter
// ctx:			contain all informations about the program
// VerifRules return a boolean. true if the rules do not create logical error.
// false elseway. It call VerifConclusion to set variables to true or false.
func VerifRules(ctx *s.Context) (success bool) {
	VerifRulesStart:
	for i:= 0; i < len(ctx.Rules); {
		switch RuleIsTrue(ctx.Rules[i].Premice, ctx.Variables) {
		case s.TRUE:
			rule := ctx.Rules[i]
			VerboseAddLine(ctx, fmt.Sprintln(rule))
			VerboseAddLine(ctx, "")
			ctx.RemoveRule(i)
			if !VerifConclusion(ctx, rule.Conclusion, []rune(UsedVar(rule.Conclusion)), "") {
				return false
			}
			goto VerifRulesStart
		case s.FALSE:
			VerboseAddLine(ctx, fmt.Sprintln(ctx.Rules[i]))
			VerboseAddLine(ctx, "impossible\n")
			ctx.RemoveRule(i)
		default:
			i++
		}
	}
	return true
}

// ComplexeCase take one parameter
// ctx:			contain all informations about the program
// ComplexeCase return two parameters
// finalctx:	this context will replace the previous one in the parent function
// end:			this boolean is a stop signal for the recursivity
// ComplexeCase is a recursive function wich call VerifRules to check
// if a rule need to be actived then if there is no more rules, a solution
// has been found and it can return the finalctx. if there is still rules,
// ComplexeCase will admit that a variable is false and call himself recursively
func ComplexeCase(ctx *s.Context) (finalctx *s.Context, end bool) {
	if !VerifRules(ctx) {
		return ctx, false
	}
	if len(ctx.Rules) == 0 {
		return ctx, true
	}
	CleanCanChange(ctx)
	for i, char := range ctx.CanChange {
		VerboseAddLine(ctx, "admit: " + string(char) + " is false\n")
		ctx.Variables[char] = false
		newctx := ctx.Copy()
		newctx.CanChange = RemoveIndex(ctx.CanChange, i)
		finalctx, end := ComplexeCase(newctx)
		ctx.Verbose += finalctx.Verbose
		if end {
			finalctx.Verbose = ctx.Verbose
			return finalctx, true
		}
		delete(ctx.Variables, char)
	}
	return ctx, false
}

func Algo(ctx *s.Context) string {
	ctx.CanChange = ChangableVariables(ctx)
	finalctx, end := ComplexeCase(ctx)
	ctx.Variables, ctx.Verbose = finalctx.Variables, finalctx.Verbose
	if !end {
		return m.Impossible
	}
	return ""
}