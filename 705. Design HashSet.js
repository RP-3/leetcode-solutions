// https://leetcode.com/problems/design-hashset/

/**
 * Initialize your data structure here.
 */
var MyHashSet = function() {
    this.storage = new Array(10_000).fill(0).map(() => []);
};

/**
 * @param {number} key
 * @return {void}
 */
MyHashSet.prototype.add = function(key) {
    const index = key % 10_000;
    for(let i=0; i<this.storage[index].length; i++){
        if(this.storage[index][i] === key) return;
    }
    this.storage[index].push(key);
};

/**
 * @param {number} key
 * @return {void}
 */
MyHashSet.prototype.remove = function(key) {
    const index = key % 10_000;
    const sublist = this.storage[index];
    const len = sublist.length;
    for(let i=0; i<len; i++){
        if(sublist[i] === key){
            if(i !== len-1){
                [sublist[i], sublist[len-1]] = [sublist[len-1], sublist[i]];
            }
            return sublist.pop();
        }
    }
};

/**
 * Returns true if this set contains the specified element
 * @param {number} key
 * @return {boolean}
 */
MyHashSet.prototype.contains = function(key) {
    const index = key % 10_000;
    for(let i=0; i<this.storage[index].length; i++){
        if(this.storage[index][i] === key) return true;
    }
    return false;
};

/**
 * Your MyHashSet object will be instantiated and called as such:
 * var obj = new MyHashSet()
 * obj.add(key)
 * obj.remove(key)
 * var param_3 = obj.contains(key)
 */