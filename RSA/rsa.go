package rsa

import (
	"crypto/rand"
	gb "github.com/OhYee/goutils/bytes"
	"math"
	"math/big"
)

var (
	bn1 = big.NewInt(-1)
	b0  = big.NewInt(0)
	b1  = big.NewInt(1)
	b2  = big.NewInt(2)
)

func nb() *big.Int {
	return big.NewInt(0)
}

func ncopy(b *big.Int) *big.Int {
	t, _ := nb().SetString(b.String(), 10)
	return t
}

func pow(a *big.Int, n *big.Int, m *big.Int) (res *big.Int) {
	if n.Cmp(b0) == 0 {
		res = b1
	} else if n.Cmp(b1) == 0 {
		res = a
	} else {
		res = pow(a, nb().Div(n, b2), m)
		res.Mul(res, res)
		res.Mod(res, m)

		if nb().Mod(n, b2).Cmp(b1) == 0 {
			res.Mul(res, a).Mod(res, m)
		}
	}
	res = ncopy(res)
	return
}

func randomBig() *big.Int {
	num, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	return num
}

func isPrime(n *big.Int) bool {
	if n.Cmp(b2) == 0 {
		return true
	}
	if n.Cmp(b2) < 0 || nb().Mod(n, b2).Cmp(b2) == 0 {
		return false
	}
	for i := 0; i < 10; i++ {
		a := randomBig()
		a.Mod(a, nb().Add(n, bn1))
		a.Add(a, b1)
		if pow(a, nb().Add(n, bn1), n).Cmp(b1) != 0 {
			return false
		}
	}
	return true
}

func randomPrime() (n *big.Int) {
	for {
		n = randomBig()
		if isPrime(n) {
			return
		}
	}
}

// func randomUint16() (n uint16) {
// 	buf := bytes.NewBuffer(gb.FromUint32(random.Uint32()))
// 	num, err := rand.Int(buf, big.NewInt(math.MaxUint16))
// 	if err != nil {
// 		panic(err)
// 	}
// 	n = uint16(num.Uint64())
// 	return
// }

func gcd(a *big.Int, b *big.Int) *big.Int {
	if b.Cmp(b0) == 0 {
		return a
	}
	return gcd(b, nb().Mod(a, b))
}

func exgcd(a *big.Int, b *big.Int) (r *big.Int, x *big.Int, y *big.Int) {
	if b.Cmp(b0) == 0 {
		return a, ncopy(b1), ncopy(b0)
	}

	c := nb().Div(a, b)
	d := nb().Mod(a, b)
	r, x, y = exgcd(b, d)

	t := x
	x = y
	y = nb().Add(t, nb().Neg(nb().Mul(c, y)))
	return r, x, y
}

// Generate a RSA key-pair
func Generate() (privateKey []byte, publicKey []byte, err error) {
	p := nb()
	q := nb()
	n := nb()
	phi := nb()
	e := nb()
	d := nb()

	for {
		p = randomPrime()
		q = randomPrime()
		n.Mul(p, q)
		phi.Mul(nb().Add(p, bn1), nb().Add(q, bn1))
		if n.Cmp(big.NewInt(256)) < 0 {
			continue
		}
		e = nb().Mod(randomBig(), phi)
		for ; e.Cmp(phi) < 0; e.Add(e, b1) {
			if gcd(e, phi).Cmp(b1) == 0 {
				break
			}
		}
		if gcd(e, phi).Cmp(b1) != 0 {
			continue
		}
		break
	}
	_, d, _ = exgcd(e, phi)
	for d.Cmp(b0) < 0 {
		d.Add(d, phi)
	}
	for d.Cmp(phi) > 0 {
		d.Add(d, nb().Neg(phi))
	}

	publicKey = append(gb.FromUint64(e.Uint64()), gb.FromUint64(n.Uint64())...)
	privateKey = append(gb.FromUint64(d.Uint64()), gb.FromUint64(n.Uint64())...)
	return
}

var sizeList = []int64{0, 256, 65536, 16777216, 4294967296, 1099511627776, 281474976710656, 72057594037927936}

func getSize(n *big.Int) (size int) {
	nn := int64(n.Uint64())
	for size = 0; size < len(sizeList); size++ {
		if nn < sizeList[size] {
			break
		}
	}
	return
}

func Encrypto(plainText []byte, privateKey []byte) (cryptoText []byte) {
	k := nb().SetBytes(privateKey[0:8])
	m := nb().SetBytes(privateKey[8:16])

	extend := getSize(m)
	cryptoText = make([]byte, 0)

	for _, b := range plainText {
		num := pow(nb().SetBytes([]byte{b}), k, m).Uint64()
		bb := gb.FromUint64(num)[8-extend : 8]
		cryptoText = append(cryptoText, bb...)
	}
	return
}

func Decrypto(cryptoText []byte, publicKey []byte) (plainText []byte) {
	k := nb().SetBytes(publicKey[0:8])
	m := nb().SetBytes(publicKey[8:16])

	extend := getSize(m)
	plainText = make([]byte, 0)

	for i := 0; i < len(cryptoText)/extend; i++ {
		data := make([]byte, 8-extend)
		data = append(data, cryptoText[i*extend:(i+1)*extend]...)
		a := nb().SetBytes(data)
		plainText = append(plainText, pow(a, k, m).Bytes()[0])
	}
	return
}
