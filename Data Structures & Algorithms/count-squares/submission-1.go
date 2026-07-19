type CountSquares struct {
    pointXMap map[int][][]int
	pointYMap map[int][][]int
	pointsCnt    map[[2]int]int
}

func Constructor() CountSquares {
    return CountSquares {
		pointXMap: make(map[int][][]int),
		pointYMap: make(map[int][][]int),
		pointsCnt: make(map[[2]int]int),
	}
}

func (this *CountSquares) Add(point []int)  {
    this.pointXMap[point[0]] = append(this.pointXMap[point[0]], point)
	this.pointYMap[point[1]] = append(this.pointYMap[point[1]], point)
	this.pointsCnt[[2]int{point[0], point[1]}]++
}

func (this *CountSquares) Count(point []int) int {
	rs := 0
    x, y := point[0], point[1]
	pointsInX := this.pointXMap[x]
	pointsInY := this.pointYMap[y]
	for _, xpoint := range pointsInX {
		xpy := xpoint[1]
		lenx := abs(xpy - y)
		if lenx == 0 {continue} // not count itself
		for _, ypoint := range pointsInY {
			ypx := ypoint[0]
			leny := abs(ypx - x)
			if leny == lenx {
				// find the last point in cornor
				targetPointX := ypoint[0]
				targetPointY := xpoint[1]
				if cnt, exist := this.pointsCnt[[2]int{targetPointX, targetPointY}]; exist {
					fmt.Println(targetPointX, targetPointY)
					rs += cnt
				}
			}
		} 
	}
	return rs
}

func abs(x int) int {
	if x < 0 {
		return x * (-1)
	}
	return x
}
