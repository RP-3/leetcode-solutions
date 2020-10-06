// https://leetcode.com/problems/insert-into-a-binary-search-tree/

/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} val
 * @return {TreeNode}
 */
var insertIntoBST = function(root, val) {
    if(!root) return new TreeNode(val);

    const inner = (node) => {
        if(!node) return new TreeNode(val);
        if(val < node.val) node.left = inner(node.left);
        else node.right = inner(node.right);
        return node;
    };

    return inner(root);
};