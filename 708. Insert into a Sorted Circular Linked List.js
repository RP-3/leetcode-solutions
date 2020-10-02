// https://leetcode.com/problems/insert-into-a-sorted-circular-linked-list/

/**
 * // Definition for a Node.
 * function Node(val, next) {
 *     this.val = val;
 *     this.next = next;
 * };
 */

/**
 * @param {Node} head
 * @param {number} insertVal
 * @return {Node}
 */
var insert = function(head, insertVal) {
    const result = new Node(insertVal);
    if(!head){
        result.next = result;
        return result;
    }

    let [curr, prev] = [head.next, head];

    while(true){
    	if(prev.val <= insertVal && insertVal <= curr.val){
    		result.next = curr; prev.next = result;	return head; // insert node here
    	}
    	if(curr === head) break; // loop complete
		prev = curr;
		curr = curr.next;
    }

    // If we made it here:
    // 1. InsertVal is outside the range of the loop, or
    // 2. The numbers in the loop are uniform

    // If its case 1, find the inflection point and insert it
    // in there.
    [curr, prev] = [head.next, head];
  	while(true){
	    if(prev.val > curr.val){
	    	result.next = curr; prev.next = result; return head;
	    }
	    if(curr === head) break; // loop complete
		prev = curr;
		curr = curr.next;
  	}

  	// The loop is uniform. Insert anywhere
  	result.next = curr; prev.next = result;

    return head;
};