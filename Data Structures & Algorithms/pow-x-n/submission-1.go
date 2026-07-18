func myPow(x float64, n int) float64 {
    if floatIsZero(x) {
		return x
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1/x
		n = n * (-1)
	}
    base := 1.0
	for n > 0 {
		if n & 1 != 0 {
			base = base * x
		}
		x = x*x
		n = n >> 1
	}
	return base

}

func floatIsZero(a float64) bool {
	const epsilon = 1e-9
	return a < epsilon && a*(-1) < epsilon
}
