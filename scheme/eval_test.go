package scheme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	t.Run("eval int value", func(t *testing.T) {
		got := Eval(`42`)
		assert.Equal(t, Int(42), got)
	})

	t.Run("eval int bultin function", func(t *testing.T) {
		assert.Equal(t, Int(42), Eval(`(+ 20 22)`))
		assert.Equal(t, Int(30), Eval(`(* 5 6)`))
		assert.Equal(t, Int(10), Eval(`(- 20 10)`))

		t.Run("multiple arguments", func(t *testing.T) {
			t.Skip("currently multiple arguments of + is not supported")
			assert.Equal(t, Int(52), Eval(`(+ 20 22 10)`))
		})

		t.Run("division", func(t *testing.T) {
			t.Skip("unsupported")
			assert.Equal(t, Int(5), Eval(`(/ 10 2)`))
		})
	})

	t.Run("complex plus", func(t *testing.T) {
		assert.Equal(t, Int(42), Eval(`(+ (* 10 2) (+ 2 20))`))
	})

	t.Run("boolean", func(t *testing.T) {
		assert.Equal(t, true, Eval(`(= 3 3)`))
		assert.Equal(t, false, Eval(`(= 3 4)`))
	})

	t.Run("if", func(t *testing.T) {
		assert.Equal(t, Int(5), Eval(`(if (= 3 3) 5 0)`))
		assert.Equal(t, Int(0), Eval(`(if (= 3 4) 5 0)`))
	})

	t.Run("null?", func(t *testing.T) {
		assert.Equal(t, true, Eval(`(null? ())`))
		assert.Equal(t, false, Eval(`(null? (quote (1)))`))
		assert.Equal(t, false, Eval(`(null? 5)`))
	})

	t.Run("car cdr cons", func(t *testing.T) {
		assert.Equal(t, Int(4), Eval(`(car (quote (4 5 6)))`))
		assert.Equal(t,
			List(Int(2), Int(3)),
			Eval(`(cdr (quote (1 2 3)))`),
		)
		assert.Equal(t, true, Eval(`(= 1 (car (cons 1 (quote (2 3)))))`))
		assert.Equal(t, true, Eval(`(= 2 (car (cdr (cons 1 (quote (2 3))))))`))
	})

	t.Run("quotes", func(t *testing.T) {
		assert.Equal(
			t,
			List(Int(1), Int(2), Int(3)),
			Eval(`(quote (1 2 3))`),
		)
		assert.Equal(
			t,
			Int(5),
			Eval(`(quote 5)`),
		)
	})

	t.Run("with variables", func(t *testing.T) {
		// arrange
		Eval(`(define forty-two (* 21 2))`)

		// act
		result := Eval(`(+ forty-two 1)`)

		// assert
		assert.Equal(t, Int(43), result)
	})

	t.Run("simple lambda", func(t *testing.T) {
		assert.Equal(t, Int(25), Eval(`((lambda (x) (* x x)) 5)`))
	})

	t.Run("define lambda", func(t *testing.T) {
		// arrange
		Eval(`(define square (lambda (x) (* x x)))`)

		// act
		result := Eval(`(square 6)`)

		// assert
		assert.Equal(t, Int(36), result)
	})

	t.Run("should evaluate arguments of lambda", func(t *testing.T) {
		// arrange
		Eval(`(define square (lambda (x) (* x x)))`)

		// act
		result := Eval(`(square (+ 3 3))`)

		// assert
		assert.Equal(t, Int(36), result)
	})

	t.Run("pass lambda in lambda lambda", func(t *testing.T) {
		// arrange
		Eval(`(define double (lambda (x) (+ x x)))`)
		Eval(`(define twice  (lambda (fn) (lambda (x) (fn (fn x)))))`)

		// act
		result := Eval(`((twice double) 10)`)

		// assert
		assert.Equal(t, Int(40), result)
	})

	t.Run("recursion", func(t *testing.T) {
		// arrange
		Eval(`
			(define fact
				(lambda (n)
					(if (= n 0)
						1
						(* n (fact (- n 1))))))
		`)

		// assert
		assert.Equal(t, Int(120), Eval(`(fact 5)`))
	})
}
