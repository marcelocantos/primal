package primal

import (
	"math/bits"
)

// given 0 <= a,b,n < 2^64, computes (a*b)%n without overflow
func safe_mul(a, b, n uint64) uint64 {
	hi, lo := bits.Mul64(a, b)
	_, rem := bits.Div64(hi, lo, n)
	return rem
}

// given 0 <= a,b,n < 2^64, computes (a^b)%n without overflow
func safe_exp(a, b, n uint64) uint64 {
	var res, pw uint64 = 1, 1
	for i := 0; i < 64 && pw <= b; i++ {
		if b&pw != 0 {
			res = safe_mul(res, a, n)
		}
		a = safe_mul(a, a, n)
		pw <<= 1
	}
	return res
}

// given 2 <= n,a < 2^64, a prime, check whether n is a-SPRP
func is_SPRP(n, a uint64) bool {
	if n == a {
		return true
	}
	if n%a == 0 {
		return false
	}
	d := n - 1
	s := 0
	for d%2 == 0 {
		s++
		d /= 2
	}
	cur := safe_exp(a, d, n)
	if cur == 1 {
		return true
	}
	for r := 0; r < s; r++ {
		if cur == n-1 {
			return true
		}
		cur = safe_mul(cur, cur, n)
	}
	return false
}

func IsPrime(x uint64) bool {
	if x == 2 || x == 3 || x == 5 || x == 7 {
		return true
	}
	if x%2 == 0 || x%3 == 0 || x%5 == 0 || x%7 == 0 {
		return false
	}
	if x < 121 {
		return x > 1
	}
	if !is_SPRP(x, 2) {
		return false
	}
	h := x
	h = ((h >> 32) ^ h) * 0x45d9f3b3335b369
	h = ((h >> 32) ^ h) * 0x3335b36945d9f3b
	h = ((h >> 32) ^ h)
	b := uint64(bases[h&16383])
	return is_SPRP(x, b&4095) && is_SPRP(x, b>>12)
}
