package scheme

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func scan(scanner *bufio.Scanner) bool {
	return scanner.Scan()
}

func isReadyToEval(s string) bool {
	bracesBalance := 0
	isString := false
	hasSexpr := false
	for _, c := range []rune(s) {
		if !strings.ContainsRune(whitespace, c) {
			hasSexpr = true
		}
		if c == '"' {
			isString = !isString
		}
		if c == '(' && !isString {
			bracesBalance++
		}
		if c == ')' && !isString {
			bracesBalance--
		}
	}
	return bracesBalance <= 0 && hasSexpr && !isString
}

func RunRepl(input io.Reader, output io.Writer) {
	// FIXME: it is not fully suitable here because newline may break sexpr which
	// results in parser error.
	buf := bufio.NewScanner(input)

	var result ExprT
	env := DefaultEnvironment()
	line := ""
	fmt.Fprint(output, "> ")
	for scan(buf) {
		line += buf.Text()
		if !isReadyToEval(line) {
			continue
		}
		result, env = EvalInEnvironment(line, env)
		fmt.Fprintln(output, result)
		fmt.Fprintln(output)
		line = ""
		fmt.Fprint(output, "> ")
	}
}
