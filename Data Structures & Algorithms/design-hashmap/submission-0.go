type MyHashMap struct {
	ele []int
}

func Constructor() MyHashMap {
    mh := MyHashMap {
		ele: make([]int, 1000001),
	}
	for i, _ := range mh.ele {
		mh.ele[i] = -1
	}
	return mh
}

func (this *MyHashMap) Put(key int, value int) {
    this.ele[key] = value
}

func (this *MyHashMap) Get(key int) int {
    return this.ele[key]
}

func (this *MyHashMap) Remove(key int) {
    this.ele[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */