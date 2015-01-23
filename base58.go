// Package base58 implements a human-friendly base58 encoding.
package base58

import (
	"math/big"
	"strconv"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var decodeMap [256]byte

func init() {
	for i := 0; i < len(decodeMap); i++ {
		decodeMap[i] = 0xFF
	}
	for i := 0; i < len(alphabet); i++ {
		decodeMap[alphabet[i]] = byte(i)
	}
}

type CorruptInputError int64

func (e CorruptInputError) Error() string {
	return "illegal base58 data at input byte " + strconv.FormatInt(int64(e), 10)
}

// Decode decodes src and return the decoded data.
// Also returns an error on corrupt input.
func Decode(src string) ([]byte, error) {
	buf := src[:]
	n := new(big.Int)
	radix := big.NewInt(58)
	for i := 0; i < len(buf); i++ {
		b := decodeMap[buf[i]]
		if b == 0xFF {
			return nil, CorruptInputError(i)
		}
		n.Mul(n, radix)
		n.Add(n, big.NewInt(int64(b)))
	}
	return n.Bytes(), nil
}

// Encode encodes src and returns the encoded data.
func Encode(src []byte) string {
	n := new(big.Int)
	n.SetBytes(src)
	radix := big.NewInt(58)
	zero := big.NewInt(0)

	dst := []byte{}
	for n.Cmp(zero) > 0 {
		mod := new(big.Int)
		n.DivMod(n, radix, mod)
		dst = append(dst, alphabet[mod.Int64()])
	}

	for i, j := 0, len(dst)-1; i < j; i, j = i+1, j-1 {
		dst[i], dst[j] = dst[j], dst[i]
	}
	return string(dst)
}
