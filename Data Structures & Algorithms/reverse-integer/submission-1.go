func reverse(x int) int {
    MIN := -2147483648 // -2^31
    MAX := 2147483647  // 2^31 - 1

    res := 0
    for x != 0 {
        digit := int(math.Mod(float64(x), 10))
        x = int(float64(x) / 10)

        if res > MAX/10 || (res == MAX/10 && digit > MAX%10) {
            return 0
        }
        if res < MIN/10 || (res == MIN/10 && digit < MIN%10) {
            return 0
        }
        res = (res * 10) + digit
		//fmt.Println(digit, x, res, x%10, x/10)
    }

    return res
}
