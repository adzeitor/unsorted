package scheme

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunRepl(t *testing.T) {
	t.Run("simplest case", func(t *testing.T) {
		// arrange
		input := bytes.NewBufferString("(+ 22 20)\n")
		output := bytes.NewBufferString("")

		// act
		RunRepl(input, output)

		// assert
		assert.Contains(t, output.String(), "42")
	})

	t.Run("multi-line expressions", func(t *testing.T) {
		// arrange
		input := bytes.NewBufferString("(+\n 10\n 20)\n")
		output := bytes.NewBufferString("")

		// act
		RunRepl(input, output)

		// assert
		assert.Contains(t, output.String(), "30")
	})

	t.Run("multi-line expressions with strings", func(t *testing.T) {
		// arrange
		input := bytes.NewBufferString(
			`( + 5 
					(car 
						(cdr 
							(quote 
								(
									"some string with braces))))))))" 
									42 
									3)))))
				`)
		output := bytes.NewBufferString("")

		// act
		RunRepl(input, output)

		// assert
		assert.Contains(t, output.String(), "47")
	})

	t.Run("multi-line define and one-line var", func(t *testing.T) {
		// arrange
		input := bytes.NewBufferString(
			`
					(define foo 42)
					foo
				`)
		output := bytes.NewBufferString("")

		// act
		RunRepl(input, output)

		// assert
		assert.Contains(t, output.String(), "42")
	})
}
