package scheme

import (
	"strconv"
	"strings"
	"unicode"
)

const whitespace = " \n\t"

func parseInt(s string) (value ExprT, remains string, ok bool) {
	var accum string
	for _, c := range []rune(s) {
		if !unicode.IsDigit(c) {
			break
		}
		ok = true
		accum = accum + string(c)
	}
	n, _ := strconv.Atoi(accum)
	return Int(n), s[len(accum):], ok
}

func parseSymbol(s string) (value ExprT, remains string, ok bool) {
	const allowedSymbolChars = "+-*=?/abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var accum string
	for _, c := range []rune(s) {
		if !strings.ContainsRune(allowedSymbolChars, c) {
			break
		}
		ok = true
		accum = accum + string(c)
	}
	return SymbolT(accum), s[len(accum):], ok
}

func parseString(s string) (value ExprT, remains string, ok bool) {
	var accum string
	s, ok = skipRune(s, `"`)
	if !ok {
		return
	}
	for _, c := range []rune(s) {
		if c == '"' {
			break
		}
		ok = true
		accum = accum + string(c)
	}
	remains = s[1+len(accum):]
	return StringT(accum), remains, ok
}

func skipRune(s string, runesToSkip string) (remains string, ok bool) {
	if len(s) == 0 {
		return s, false
	}
	firstRune := []rune(s)[0]
	if strings.ContainsRune(runesToSkip, firstRune) {
		return string([]rune(s)[1:]), true
	}
	return s, false
}

func skipManyRune(s string, toSkip string) (remains string, ok bool) {
	ok = true
	for ok {
		s, ok = skipRune(s, toSkip)
	}
	return s, true
}

// use sepBy
func parseList(s string) (value ExprT, remains string, ok bool) {
	remains, ok = skipRune(s, "(")
	if !ok {
		return
	}

	var list []ExprT
	for {
		s, _ = skipManyRune(s, whitespace)
		value, remains, ok = Parse(remains)
		if !ok {
			break
		}
		list = append(list, value)

		// optional space between elements (not for last element)
		// FIXME: many spaces
		remains, ok = skipManyRune(remains, whitespace)
		if !ok {
			break
		}
	}

	remains, ok = skipRune(remains, ")")
	if !ok {
		return
	}
	return List(list...), remains, ok
}

// oneOf?
func Parse(s string) (value ExprT, remains string, ok bool) {
	s, _ = skipManyRune(s, whitespace)
	value, remains, ok = parseInt(s)
	if ok {
		return value, remains, ok
	}

	value, remains, ok = parseSymbol(s)
	if ok {
		return value, remains, ok
	}

	value, remains, ok = parseString(s)
	if ok {
		return value, remains, ok
	}

	value, remains, ok = parseList(s)
	if ok {
		return value, remains, ok
	}
	return value, s, false
}
