package math

type Ints interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr
}

type Floats interface {
	~float32 | ~float64
}

// Abs は与引数の絶対値を返す
// Time: O(1)
func Abs[T Ints](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

// Gcd は与引数 (a,b) の最大公約数( >= 0)を返す
// (a,b) = (0,0) の場合は 0 を返す
// Time: O(log N)
func Gcd[T Ints](a, b T) T {
	if a < 0 || b < 0 {
		return Gcd(Abs(a), Abs(b))
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Lcm は与引数 (a,b) の最小公倍数( >= 0)を返す
// (a,b) = (0,0) の場合は 0 を返す
// Time: O(log N)
func Lcm[T Ints](a, b T) T {
	if a < 0 || b < 0 {
		return Lcm(Abs(a), Abs(b))
	}
	if a == 0 && b == 0 {
		return 0
	}
	return a / Gcd(a, b) * b
}
