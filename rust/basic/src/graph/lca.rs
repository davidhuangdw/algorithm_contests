use std::iter::repeat_with;

// O(n+e) find all quried LCA(least common ancestor)
pub fn tarjan_lca(root: usize, adj: &[Vec<usize>], queries: &[(usize, usize)]) -> Vec<usize> {
    struct State<'a> {
        chd: &'a [Vec<usize>],
        queries: &'a [(usize, usize)],
        fa: Vec<usize>,
        vis: Vec<bool>,
        q_idx: Vec<Vec<usize>>,
        ans: Vec<usize>,
    }

    // union_find
    fn find(i: usize, fa: &mut Vec<usize>) -> usize {
        if fa[i] != i {
            fa[i] = find(fa[i], fa);
        }
        fa[i]
    }

    fn dfs(i: usize, p: usize, st: &mut State) {
        if st.vis[i] {
            panic!("not a tree: has cycle or multiple parent");
        }
        st.vis[i] = true;
        for &qi in &st.q_idx[i] {
            let (a, b) = st.queries[qi];
            let j = a ^ b ^ i;
            if st.vis[j] {
                st.ans[qi] = find(j, &mut st.fa);
            }
        }

        for &j in &st.chd[i] {
            if j != p {
                dfs(j, i, st);
                st.fa[j] = i;
            }
        }
    }

    let n = adj.len();
    let mut st = State {
        chd: adj,
        queries: queries,
        fa: (0..n).collect(),
        vis: vec![false; n],
        q_idx: repeat_with(|| vec![]).take(n).collect(),
        ans: vec![0; queries.len()],
    };
    for (i, (a, b)) in queries.iter().enumerate() {
        st.q_idx[*a].push(i);
        st.q_idx[*b].push(i);
    }
    dfs(root, root, &mut st);
    for i in 0..n {
        if !st.vis[i] {
            panic!("{} is unreachable", i);
        }
    }
    st.ans
}

#[cfg(test)]
mod tests {
    use super::*;

    /// Brute force: root the tree, then climb by depth.
    fn brute(root: usize, adj: &[Vec<usize>], qs: &[(usize, usize)]) -> Vec<usize> {
        let n = adj.len();
        let (mut par, mut dep, mut seen) = (vec![root; n], vec![0; n], vec![false; n]);
        let mut stack = vec![root];
        seen[root] = true;
        while let Some(u) = stack.pop() {
            for &v in &adj[u] {
                if !seen[v] {
                    seen[v] = true;
                    par[v] = u;
                    dep[v] = dep[u] + 1;
                    stack.push(v);
                }
            }
        }
        qs.iter()
            .map(|&(mut a, mut b)| {
                while dep[a] > dep[b] {
                    a = par[a];
                }
                while dep[b] > dep[a] {
                    b = par[b];
                }
                while a != b {
                    a = par[a];
                    b = par[b];
                }
                a
            })
            .collect()
    }

    #[test]
    fn matches_brute_force_on_random_trees() {
        let mut seed = 1u64;
        for _ in 0..50 {
            seed = seed.wrapping_mul(6364136223846793005).wrapping_add(1);
            let n = (seed % 12) as usize + 1;
            let mut adj = vec![vec![]; n];
            for v in 1..n {
                seed = seed.wrapping_mul(6364136223846793005).wrapping_add(1);
                let p = (seed >> 33) as usize % v;
                adj[p].push(v);
                adj[v].push(p); // undirected tree from root 0
            }
            let qs: Vec<_> = (0..n).flat_map(|a| (0..n).map(move |b| (a, b))).collect();
            assert_eq!(brute(0, &adj, &qs), tarjan_lca(0, &adj, &qs), "n = {n}");
        }
    }

    #[test]
    fn edge_cases() {
        assert_eq!(vec![0], tarjan_lca(0, &[vec![]], &[(0, 0)]));

        let chain = [vec![1], vec![0, 2], vec![1]];
        let qs = [(1, 2), (0, 2), (2, 2), (0, 1)];
        assert_eq!(vec![1, 0, 2, 0], tarjan_lca(0, &chain, &qs));

        let star = [vec![1, 2], vec![0], vec![0]];
        assert_eq!(vec![0], tarjan_lca(0, &star, &[(1, 2)]));

        // Children-list (directed) form works as long as each node has one parent.
        let chd = [vec![1, 2], vec![3, 4], vec![5], vec![], vec![], vec![]];
        let qs = [(3, 4), (3, 5), (4, 5), (1, 5)];
        assert_eq!(vec![1, 0, 0, 0], tarjan_lca(0, &chd, &qs));
    }

    #[test]
    #[should_panic(expected = "is unreachable")]
    fn rejects_unreachable_node() {
        let forest = [vec![1], vec![0], vec![]]; // node 2 is not reachable from 0
        tarjan_lca(0, &forest, &[(2, 2)]);
    }

    #[test]
    #[should_panic(expected = "has cycle")]
    fn rejects_cycle() {
        let triangle = [vec![1, 2], vec![0, 2], vec![1, 0]];
        tarjan_lca(0, &triangle, &[(0, 1)]);
    }
}
