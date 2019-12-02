package sha

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/OhYee/goutils/functional"
	"strings"
)

type register64 struct {
	v []uint64
	l int
}

func newRegister64(iv ...uint64) *register64 {
	v := make([]uint64, len(iv))
	copy(v, iv)
	return &register64{
		v: v,
		l: len(iv),
	}
}

func (reg *register64) toString() string {
	s := make([]string, reg.l)
	for i, v := range reg.v {
		s[i] = fmt.Sprintf("%016x", v)
	}
	return strings.Join(s, " ")
}

func (reg *register64) op(reg2 *register64, f func(uint64, uint64) uint64) {
	for i := range reg.v {
		reg.v[i] = f(reg.v[i], reg2.v[i])
	}
}

func (reg register64) copy() *register64 {
	return newRegister64(reg.v...)
}

func (reg register64) toBytes() []byte {
	// transfer from []uint64 to []byte
	buf := bytes.NewBuffer([]byte{})
	temp := make([]byte, 8)
	fp.MapUint64(func(data uint64, idx int) uint64 {
		binary.BigEndian.PutUint64(temp, data)
		buf.Write(temp)
		return 0
	}, reg.v)
	return buf.Bytes()
}
