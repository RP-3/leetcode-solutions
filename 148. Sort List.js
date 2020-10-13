// https://leetcode.com/problems/sort-list/

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var sortList = function(head) {
    if(!head || !head.next) return head; // Unit case. List of size 1 or 0

    const [leftHead, rightHead] = splitList(head); // List is size two or more, so divide
    const [leftSorted, rightSorted] = [sortList(leftHead), sortList(rightHead)]; // and recurse

    // Now we know that we have lists that adhere to the properties of sortList.
    // I.e., sortList will *always* return a sorted list, or a single node, or null.
    // Merge the (sorted lists | single nodes | null nodes) together to get a single sorted list
    return zipLists(leftSorted, rightSorted);
};

const splitList = (head) => {
    if(!head) return [null, null];
    let [slow, fast] = [head, head.next];

    while(fast && fast.next){
        slow = slow.next;
        fast = fast.next.next;
    }

    const rightHead = slow.next;
    slow.next = null;
    return [head, rightHead];
};

const zipLists = (l, r) => {
    const dummy = {};
    let tail = dummy;

    while(l || r){
        if(!l){
            tail.next = r;
            r = r.next;
        }
        else if(!r){
            tail.next = l;
            l = l.next;
        }
        else{
            if(l.val < r.val){
                tail.next = l;
                l = l.next;
            }else{
                tail.next = r;
                r = r.next;
            }
        }
        tail = tail.next;
    }
    tail.next = null;

    return dummy.next;
};