type RandomizedSet struct {
    itemsMap map[int]int
    itemsSlice []int
}

func Constructor() RandomizedSet {
    return RandomizedSet{
        itemsMap : make(map[int]int),
        itemsSlice: make([]int, 0),
    }

}

// returns true if the item was not present
func (this *RandomizedSet) Insert(val int) bool {
    _, exists := this.itemsMap[val]
    if exists {
        return false
    }
    idx := len(this.itemsSlice)
    this.itemsSlice = append(this.itemsSlice, val)
    this.itemsMap[val] = idx
    return true
}

// returns true if the was present
func (this *RandomizedSet) Remove(val int) bool {
    idx, exists := this.itemsMap[val]
    if !exists {
        return false
    }
    delete(this.itemsMap, val)
    lastIdx := len(this.itemsSlice) - 1
    if idx != len(this.itemsSlice) - 1 {
        lastVal := this.itemsSlice[lastIdx]
        this.itemsSlice[idx] = lastVal
        this.itemsMap[lastVal] = idx
    }
    this.itemsSlice = this.itemsSlice[:lastIdx]
    return true
}

func (this *RandomizedSet) GetRandom() int {
    return this.itemsSlice[rand.Intn(len(this.itemsSlice))]
}