package scheme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// FIXME: table test?
func TestParse(t *testing.T) {
	t.Run("numbers", func(t *testing.T) {
		input := `5`
		want := Int(5)

		got, _, _ := Parse(input)
		assert.Equal(t, want, got)
	})

	t.Run("empty list", func(t *testing.T) {
		input := `()`
		want := List()

		got, _, _ := Parse(input)
		assert.Equal(t, want, got)
	})

	t.Run("symbol", func(t *testing.T) {
		input := `+`
		want := SymbolT("+")

		got, _, _ := Parse(input)
		assert.Equal(t, want, got)
	})

	t.Run("list of strings", func(t *testing.T) {
		input := `("foo" "" "bar")`
		want := List(StringT("foo"), StringT(""), StringT("bar"))

		got, _, _ := Parse(input)
		assert.Equal(t, want, got)
	})

	t.Run("(+ 1 2)", func(t *testing.T) {
		input := `(+ 1 2)`
		want := List(SymbolT("+"), Int(1), Int(2))

		got, _, _ := Parse(input)
		assert.Equal(t, want, got)
	})

	t.Run("(foo (+ 42 24))", func(t *testing.T) {
		input := `(foo (+ 42 24))`
		want := List(
			SymbolT("foo"),
			List(SymbolT("+"), Int(42), Int(24)),
		)

		got, _, _ := Parse(input)
		assert.Equal(t, want, got)
	})

	t.Run("(* (+ 1 1) (+ 20 1))", func(t *testing.T) {
		input := `(* (+ 1 1) (+ 20 1))`
		want := List(
			SymbolT("*"),
			List(SymbolT("+"), Int(1), Int(1)),
			List(SymbolT("+"), Int(20), Int(1)),
		)

		got, _, _ := Parse(input)
		assert.Equal(t, want, got)
	})

	t.Run("no spaces between elements", func(t *testing.T) {
		input := "(1\n2\n3)"
		want := List(
			Int(1), Int(2), Int(3),
		)

		got, _, _ := Parse(input)
		assert.Equal(t, want, got)
	})

	t.Run("no whitespaces", func(t *testing.T) {
		input := "((())(())()())"
		want := List(
			List(List()), List(List()), List(), List(),
		)

		got, _, _ := Parse(input)
		assert.Equal(t, want, got)
	})

	t.Run("with many white spaces", func(t *testing.T) {
		input := "\t (  1    2 \t3  \t )\t "
		want := List(
			Int(1), Int(2), Int(3),
		)

		got, _, _ := Parse(input)
		assert.Equal(t, want, got)
	})

	t.Run("with many white spaces but broken", func(t *testing.T) {
		input := "\t (  1    2\t \t3  \t )\t "
		want := List(
			Int(1), Int(2), Int(3),
		)

		got, _, _ := Parse(input)
		assert.Equal(t, want, got)
	})
}
