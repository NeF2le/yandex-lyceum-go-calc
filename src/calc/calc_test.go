package calc

import (
	"testing"
)

func TestCalc(t *testing.T) {
	cases := []struct {
		name       string
		expression string
		want       float64
		wantErr    bool
	}{
		{
			name:       "default expression",
			expression: "3 + 4",
			want:       7,
			wantErr:    false,
		},
		{
			name:       "with brackets",
			expression: "6 + 7 - (8 + 2)",
			want:       3,
			wantErr:    false,
		},
		{
			name:       "prioritization",
			expression: "3 + 7 * 9",
			want:       66,
			wantErr:    false,
		},
		{
			name:       "division by zero",
			expression: "3 / 0",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "complex expression with nested brackets",
			expression: "5 + (6 - 2) * (7 + 3) - 8 / 2",
			want:       41,
			wantErr:    false,
		},
		{
			name:       "negative numbers",
			expression: "-3 + 4 * (-2)",
			want:       -11,
			wantErr:    false,
		},
		{
			name:       "all operations",
			expression: "4 + 18 / 3 - 2 * 3",
			want:       4,
			wantErr:    false,
		},
		{
			name:       "empty expression",
			expression: "",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "only whitespace expression",
			expression: "   ",
			want:       0,
			wantErr:    true,
		},
		{
			name: "without whitespace expression",
			expression: "3+8*9",
			want: 75,
			wantErr: false,
		},
		{
			name: "invalid operators",
			expression: "5 + 9 = 14",
			want: 0,
			wantErr: true,
		},
		{
			name:       "simple addition with decimals",
			expression: "3.5 + 4.2",
			want:       7.7,
			wantErr:    false,
		},
		{
			name:       "incorrect expression with missing operand",
			expression: "3 + * 5",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "unbalanced opening bracket",
			expression: "(3 + 2",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "unbalanced closing bracket",
			expression: "3 + 2)",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "extra closing bracket",
			expression: "(3 + 2)) - 4",
			want:       0,
			wantErr:    true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Calc(tc.expression)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("Calc(%q) = %v; want error", tc.expression, got)
				}
			} else {
				if err != nil {
					t.Fatalf("Calc(%q) returned error: %v, want %v", tc.expression, err, tc.want)
				}
				if got != tc.want {
					t.Fatalf("Calc(%q) = %.2f, want %.2f", tc.expression, got, tc.want)
				}
			}
		})
	}
}