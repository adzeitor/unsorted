package shunting_yard

import (
	"reflect"
	"testing"
)

func Test_toPostfix(t *testing.T) {
	tests := []struct {
		name string
		expr []string
		want []string
	}{
		{
			name: "42 - 99",
			expr: []string{
				"42", "-", "99",
			},
			want: []string{
				"42", "99", "-",
			},
		},
		{
			name: "1 + ( 2 * 3 )",
			expr: []string{
				"1", "+", "(", "2", "*", "3", ")",
			},
			want: []string{
				"1", "2", "3", "*", "+",
			},
		},
		{
			name: "2 * (3 + 1)",
			expr: []string{
				"2", "*", "(", "3", "+", "1", ")",
			},
			want: []string{
				"2", "3", "1", "+", "*",
			},
		},
		// FIXME
		//{
		//	name: "sin(5)",
		//	expr: []string{
		//		"sin", "(", "5", ")",
		//	},
		//	want: []string{"5", "sin"},
		//},
		{
			name: "1 + ( 2 * 3 ) * ( 4 * 5 )",
			expr: []string{
				"1", "+", "(", "2", "*", "3", ")", "*", "(", "4", "*", "5", ")",
			},
			want: []string{
				"1", "2", "3", "*", "4", "5", "*", "*", "+",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToPostfix(tt.expr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toPostfix() = %v, want %v", got, tt.want)
			}
		})
	}
}
