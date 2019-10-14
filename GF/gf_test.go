package gf

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Plus(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "1^2",
			a:    0b1101,
			b:    0b1001,
			want: 0b0100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Plus(tt.a, tt.b); got != tt.want {
				t.Errorf("Plus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Multiplus(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "2*0x87",
			a:    2,
			b:    0x87,
			want: 0b100001110,
		},
		{
			name: "3*0x6e",
			a:    3,
			b:    0x6e,
			want: 0b10110010,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiplus(tt.a, tt.b); got != tt.want {
				t.Errorf("Multiplus() = %x, want %x", got, tt.want)
			}
		})
	}
}

func TestMultiPlusTable(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		m         int
		wantTable [][]int
	}{
		{
			name: "GF(2^3)",
			n:    3,
			m:    0b1011,
			wantTable: [][]int{
				{0b000, 0b000, 0b000, 0b000, 0b000, 0b000, 0b000, 0b000},
				{0b000, 0b001, 0b010, 0b011, 0b100, 0b101, 0b110, 0b111},
				{0b000, 0b010, 0b100, 0b110, 0b011, 0b001, 0b111, 0b101},
				{0b000, 0b011, 0b110, 0b101, 0b111, 0b100, 0b001, 0b010},
				{0b000, 0b100, 0b011, 0b111, 0b110, 0b010, 0b101, 0b001},
				{0b000, 0b101, 0b001, 0b100, 0b010, 0b111, 0b011, 0b110},
				{0b000, 0b110, 0b111, 0b001, 0b101, 0b011, 0b010, 0b100},
				{0b000, 0b111, 0b101, 0b010, 0b001, 0b110, 0b100, 0b011},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTable := MultiPlusTable(tt.n, tt.m); !reflect.DeepEqual(gotTable, tt.wantTable) {
				l := len(gotTable)
				s := "\n"
				for i := 0; i < l; i++ {
					for j := 0; j < l; j++ {
						s += fmt.Sprintf(fmt.Sprintf("%%0%db ", tt.n), gotTable[i][j])
					}
					s += "\n"
				}
				t.Errorf("MultiPlusTable() = %v, want %v", s, tt.wantTable)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		b     int
		wantC int
		wantR int
	}{
		{
			name:  "11100 / 1011 = 11 ... 1",
			a:     0b11100,
			b:     0b1011,
			wantC: 0b11,
			wantR: 0b1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, gotR := Divide(tt.a, tt.b)
			if gotC != tt.wantC {
				t.Errorf("Divide() gotC = %v, want %v", gotC, tt.wantC)
			}
			if gotR != tt.wantR {
				t.Errorf("Divide() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestInverse(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		m     int
		wantC int
	}{
		{
			name:  "7^-1 mod 8",
			a:     0b111,
			m:     0b1011,
			wantC: 0b100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := Inverse(tt.a, tt.m); gotC != tt.wantC {
				t.Errorf("Inverse() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
