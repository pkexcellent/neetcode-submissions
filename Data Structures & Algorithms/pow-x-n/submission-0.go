func myPow(x float64, n int) float64 {
	if floatIsZero(x) {
		return x
	}
	if n == 0 {
		return 1
	}

    base := 1.0
	if n < 0 {
		x = float64(1)/x
		n = n * (-1)
	}
	for n > 0 {
		base = base*x
		n--
	}
	return base
}

func floatIsZero(a float64) bool {
	epsilon := 1e-9
	return a < epsilon && a*(-1) < epsilon
}
