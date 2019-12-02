package sha

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/OhYee/goutils/functional"
	"strings"
)

type register32 struct {
	v []uint32
	l int
}

func newRegister32(iv ...uint32) *register32 {
	v := make([]uint32, len(iv))
	copy(v, iv)
	return &register32{
		v: v,
		l: len(iv),
	}
}

func (reg *register32) toString() string {
	s := make([]string, reg.l)
	for i, v := range reg.v {
		s[i] = fmt.Sprintf("%08x", v)
	}
	return strings.Join(s, " ")
}

func (reg *register32) op(reg2 *register32, f func(uint32, uint32) uint32) {
	for i := range reg.v {
		reg.v[i] = f(reg.v[i], reg2.v[i])
	}
}

func (reg register32) copy() *register32 {
	return newRegister32(reg.v...)
}

func (reg register32) toBytes() []byte {
	// transfer from []uint32 to []byte
	buf := bytes.NewBuffer([]byte{})
	temp := make([]byte, 4)
	fp.MapUint32(func(data uint32, idx int) uint32 {
		binary.BigEndian.PutUint32(temp, data)
		buf.Write(temp)
		return 0
	}, reg.v)
	return buf.Bytes()
}
