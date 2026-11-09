/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l, r := 1, n
	for l <= r {
		m1, m2 := l + (r-l)/3, l + (r-l)*2/3
		//fmt.Println(m1, m2)
		rs1, rs2 := guess(m1), guess(m2)
		if rs1 == 0 {
			return m1
		} else if rs2 == 0 {
			return m2
		} else if rs1 == -1 {
			r = m1 - 1
		} else if rs2 == 1 {
			l = m2 + 1
		} else {
			l = m1 + 1
			r = m2 - 1
		}
	}
	return 0
}
