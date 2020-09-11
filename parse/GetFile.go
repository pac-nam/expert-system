package parse

import (
	"bufio"
	"os"
	"strings"
	m "expert-system/messages"
	"fmt"
	s "expert-system/structures"
)

func epur(line string) string {
	line = strings.Split(line, "#")[0]
	toChange := []string{"\r", "\v", "\t", "\f", "\n", " "}
	for _, replace := range toChange {
			line = strings.Replace(line, replace, "", -1)
	}
	return line
}

func parseLine(ctx *s.Context, line string) string {
	// fmt.Println(line)
	if len(line) == 0 {
		return ""
	}
	if line[0] == '=' {
		if ctx.Initial[0] != '=' {
			return m.MultipleInitialStates
		}
		ctx.Initial = []byte(line[1:])
		ctx.Variables = make(map[byte]bool)
		for _, char := range ctx.Initial {
			ctx.Variables[char] = true
		}
	} else if line[0] == '?' {
		if ctx.Query[0] != '?' {
			return m.MultipleQuery
		}
		ctx.Query = []byte(line[1:])
	} else {
		if err := checkLine(line); err != "" {
			return err
		}
		ctx.Rules = append(ctx.Rules, []byte(line))
	}
	return ""
}

func parseFile(ctx *s.Context, filename string) string {
	ctx.Rules = make([][]byte, 0)
	file, err := os.Open(filename)
	if err != nil {
			return m.OpenError
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err2 := parseLine(ctx, epur(scanner.Text())); err2 != "" {
			return err2
		}
	}
	if err = scanner.Err(); err != nil {
			return m.ReadError
	}
	return ""
}

func InitVariables(ctx *s.Context) string {
	for _, char := range ctx.Initial {
		_, exist := ctx.Variables[char]
		if exist {
			return "double variable: '" + string(char) + "' in initial state"
		}
		ctx.Variables[char] = true
	}
	return ""
}

func Parse() (*s.Context, string) {
	ctx := s.Context{Initial: []byte("="), Query: []byte("?"), Variables: make(map[byte]bool)}
	if len(os.Args) != 2 || os.Args[1] == "-h" {
		fmt.Println(m.Help)
		os.Exit(0)
	}
	err := parseFile(&ctx, os.Args[1])
	if err != "" {
		return &ctx, err
	}
	err = InitVariables(&ctx)
	// fmt.Print(ctx)
	return &ctx, err
}