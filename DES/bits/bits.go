package bits

type Bits uint64

func (b Bits) bit() Bits {
	if b > 0 {
		return 1
	}
	return 0
}

func (b Bits) Get(pos int) Bits {
	return (b & (1 << pos)).bit()
}

func (b *Bits) Set(pos int) Bits {
	*b = *b | (1 << pos)
	return *b
}

func (b *Bits) Clr(pos int) Bits {
	*b = *b & ^(1 << pos)
	return *b
}

func (b *Bits) SetValue(pos int, bb Bits) Bits {
	if bb.bit() > 0 {
		return b.Set(pos)
	}
	return b.Clr(pos)
}

func (b Bits) Mask(len int) Bits {
	return b & ((1 << len) - 1)
}

func (b Bits) LeftLoop(pos int, len int) Bits {
	value := b & ((1 << len) - 1)
	for i := 0; i < pos; i++ {
		overflow := value.Get(len - 1)
		value <<= 1
		value.SetValue(0, overflow)
	}
	return value & ((1 << len) - 1)
}
