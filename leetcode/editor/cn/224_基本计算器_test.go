package leet

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "基本计算器",
			args: args{s: "1 + 1"},
			want: 2,
		}, {
			name: "基本计算器",
			args: args{s: "(1+(4+5+2)-3)+(6+8)"},
			want: 23,
		}, {
			name: "基本计算器",
			args: args{s: "2-(5-6)"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
