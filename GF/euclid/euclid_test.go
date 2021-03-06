package euclid

import (
	"testing"
)

func TestGCD(t *testing.T) {
	type args struct {
		a         int
		b         int
		plus      Operator
		multiplus Operator
		divide    Operator2
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "GCD(2,5)",
			args: args{
				a:         2,
				b:         5,
				plus:      Plus,
				multiplus: Multiplus,
				divide:    Divide,
			},
			want: 1,
		},
		{
			name: "GCD(12,24)",
			args: args{
				a:         12,
				b:         24,
				plus:      Plus,
				multiplus: Multiplus,
				divide:    Divide,
			},
			want: 12,
		},
		{
			name: "GCD(26,91)",
			args: args{
				a:         26,
				b:         91,
				plus:      Plus,
				multiplus: Multiplus,
				divide:    Divide,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCD(tt.args.a, tt.args.b, tt.args.plus, tt.args.multiplus, tt.args.divide); got != tt.want {
				t.Errorf("GCD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExGCD(t *testing.T) {
	type args struct {
		a         int
		b         int
		plus      Operator
		multiplus Operator
		divide    Operator2
	}
	tests := []struct {
		name  string
		args  args
		wantR int
		wantX int
		wantY int
	}{
		{
			name: "ExGCD(12,24)",
			args: args{
				a:         12,
				b:         24,
				plus:      Plus,
				multiplus: Multiplus,
				divide:    Divide,
			},
			wantR: 12,
			wantX: 1,
			wantY: 0,
		},
		{
			name: "ExGCD(26,91)",
			args: args{
				a:         26,
				b:         91,
				plus:      Plus,
				multiplus: Multiplus,
				divide:    Divide,
			},
			wantR: 13,
			wantX: -3,
			wantY: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotX, gotY := ExGCD(tt.args.a, tt.args.b, tt.args.plus, tt.args.multiplus, tt.args.divide)
			if gotR != tt.wantR {
				t.Errorf("ExGCD() gotR = %v, want %v", gotR, tt.wantR)
			}
			if gotX != tt.wantX {
				t.Errorf("ExGCD() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("ExGCD() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}
