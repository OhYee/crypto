package bits

import (
	// "fmt"
	"testing"
)

func Test_main(t *testing.T) {

	var b Bits = 0x01

	judge := func(bb Bits) {
		if b != bb {
			t.Errorf("Excepted 0x%x, got 0x%x", bb, b)
		}
	}

	b.Clr(0)
	judge(0x00)

	b.Set(0)
	judge(0x01)

	b.Set(63)
	judge(0x8000000000000001)

	b.SetValue(7, 1)
	judge(0x8000000000000081)

	b.SetValue(7, 0)
	judge(0x8000000000000001)

	if c := b.Get(63); c != 1 {
		t.Errorf("Excepted %d, got %d", 1, c)
	}
	if c := b.Get(1); c != 0 {
		t.Errorf("Excepted %d, got %d", 0, c)
	}

	b = 0xfffffffffffffffe
	b = b.LeftLoop(1, 4)
	judge(0xD)
	b = b.LeftLoop(2, 4)
	judge(0x7)

	b = 0xabcdef
	b = b.Mask(3)
	judge(0x7)
}
