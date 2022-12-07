package number

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Gcd(a, b int) int {
	if a < 0 || b < 0 {
		return Gcd(Abs(a), Abs(b))
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	if a < 0 || b < 0 {
		return Lcm(Abs(a), Abs(b))
	}
	if a == 0 && b == 0 {
		return 0
	}
	return a / Gcd(a, b) * b
}
