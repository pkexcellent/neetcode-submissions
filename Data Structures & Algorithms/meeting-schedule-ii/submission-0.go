/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	// sweep line
	points := make(map[int]int)
	for _, interval := range intervals {
		points[interval.start]++
		points[interval.end]--
	}
	t := []int{}
	for k, _ := range points {
		t = append(t, k)
	}
	sort.Ints(t)
	maxRoom := 0
	curRoom := 0
	for _, ti := range t {
		curRoom += points[ti]
		if curRoom > maxRoom {
			maxRoom = curRoom
		}
	}
	return maxRoom
}
