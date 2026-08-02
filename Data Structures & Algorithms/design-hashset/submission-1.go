type MyHashSet struct {
	ele []bool
}

func Constructor() MyHashSet {
    return MyHashSet {
		ele: make([]bool, 10000000),
	}
}

func (this *MyHashSet) Add(key int) {
    this.ele[key] = true
}

func (this *MyHashSet) Remove(key int) {
    this.ele[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
    return this.ele[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 