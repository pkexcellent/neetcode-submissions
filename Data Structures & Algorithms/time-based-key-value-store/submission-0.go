type pair struct {
	timestamp int
	value string
}
type TimeMap struct {
	// need a map with key as key
	data map[string][]pair
}

func Constructor() TimeMap {
	return TimeMap {
		data: make(map[string][]pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if ps, exist := this.data[key]; !exist {
		return ""
	} else {
		idx := sort.Search(len(ps), func(i int) bool {return ps[i].timestamp > timestamp})
		if idx == 0 {
			return ""
		} else {
			return ps[idx-1].value
		}
	}
	return ""
}
