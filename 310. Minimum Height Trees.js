// https://leetcode.com/problems/minimum-height-trees/

/**
 * @param {number} n
 * @param {number[][]} edges
 * @return {number[]}
 */
var findMinHeightTrees = function(n, edges) {
    if(n === 1) return [0];

    const graph = new Map();
    edges.forEach(([from, to]) => {
        if(!graph.has(from)) graph.set(from, new Set());
        if(!graph.has(to)) graph.set(to, new Set());
        graph.get(from).add(to);
        graph.get(to).add(from);
    });

    let q = [];
    for(let [node, siblings] of graph) if(siblings.size === 1) q.push(node);

    while(q.length && graph.size > 2){
        const sisters = [];
        for(let i=0; i<q.length; i++){
            const node = q[i];
            for(let sibling of graph.get(node) || new Set()){
                (graph.get(sibling) || new Set()).delete(node);
                if((graph.get(sibling) || new Set()).size === 1) sisters.push(sibling);
            }
            delete graph.delete(node);
        }
        q = sisters;
    }

    return Array.from(graph.keys());
};