// https://leetcode.com/problems/inorder-successor-in-bst-ii/

/**
 * // Definition for a Node.
 * function Node(val) {
 *    this.val = val;
 *    this.left = null;
 *    this.right = null;
 *    this.parent = null;
 * };
 */

/**
 * @param {Node} node
 * @return {Node}
 */
var inorderSuccessor = function(node) {

    let [parent, child] = [node, node.right];

    if(child){
        while(child){
            parent = child;
            child = child.left;
        }
        return parent;
    };

    [child, parent] = [node, node.parent];
    while(parent){
        if(child === parent.left) return parent;
        child = parent;
        parent = parent.parent;
    }

    return null;
};