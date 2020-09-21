package solve

import (
	"fmt"
	s "expert-system/structures"
)

func VerboseAddLine(ctx *s.Context, line string) {
	for i := 0; i < ctx.DeepLevel; i++ {
		ctx.Verbose += "  |"
	}
	ctx.Verbose += line
}

func VerboseVariables(ctx *s.Context) {
	return
	if ctx.Flag_v {
		for key, value := range ctx.Variables {
			VerboseAddLine(ctx, fmt.Sprintln(string(key) + ":", value))
		}
	}
}