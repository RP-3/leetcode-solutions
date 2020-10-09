// https://leetcode.com/problems/serialize-and-deserialize-bst/


// Approach 1: Straight-forward DFS with null-terminators
var serialize = function(root) {
    const result = [];

    const dfs = (node) => {
        if(!node) return result.push('x'); // use x to signify null
        dfs(node.left);
        result.push(node.val.toString(10));
        dfs(node.right);
    };

    dfs(root);

    return result.join(':');
};

var deserialize = function(data) {
    data = data.split(':');
    let i = 0;

    const dfs = () => {
        const current = data[i++];
        if(current === 'x' || current === undefined) return null; // again, use x or end of array to signify null
        const newNode = new TreeNode(parseInt(current, 10));
        newNode.left = dfs();
        newNode.right = dfs();
        return newNode;
    };

    return dfs();
};


// Approach 2: Pre-order traversal sans null terminators
var serialize = function(root) {
    const result = [];

    const pot = (node) => {
        if(!node) return;
        result.push(node.val.toString(10)); // For a full tree, this saves about 50% space
        pot(node.left);
        pot(node.right);
    };

    pot(root);
    return result.join(':');
};

var deserialize = function(data) {
    data = data.split(':');
    let i = 0;

    const pot = (min, max) => {
        if(i >= data.length || !data[i]) return null; // As usual, if we run out of data, stop

        const curr = parseInt(data[i], 10); // parse out the next value
        if(curr > max || curr < min) return null; // it has to be within the bounds specified by the parent

        i++;
        const newNode = new TreeNode(curr);
        newNode.left = pot(min, curr);
        newNode.right = pot(curr, max);
        return newNode;
    };

    return pot(-Infinity, Infinity); // the first value is the root, so anything goes
};