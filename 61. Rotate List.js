// https://leetcode.com/problems/rotate-list/

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */
var rotateRight = function(head, k) {
    if(!head) return null
    k %= length(head);
    if(k === 0) return head;

    let [slow, fast] = [head, head];
    while(k--) fast = fast.next;

    while(fast.next){
        slow = slow.next;
        fast = fast.next;
    }

    const newHead = slow.next;
    slow.next = null;
    fast.next = head;
    return newHead;

};

const length = (head) => {
    let result = 0;
    while(head){
        result++;
        head = head.next;
    }
    return result;
}