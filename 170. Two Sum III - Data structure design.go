// https://leetcode.com/problems/two-sum-iii-data-structure-design/

package leetcode

type TwoSum struct {
	data map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{data: make(map[int]int)}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.data[number]++
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for n, count := range this.data {
		remainder := value - n
		if remainder == n && count > 1 {
			return true
		}
		if _, exists := this.data[remainder]; exists {
			if (remainder != n) || (remainder == n && count > 1) {
				return true
			}
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
