pub fn topo_order(adj: &[Vec<usize>]) -> Vec<usize> {
    let n = adj.len();
    fn dfs(i: usize, adj: &[Vec<usize>], vis: &mut Vec<bool>, ans: &mut Vec<usize>) {
        if vis[i] {
            return;
        }
        vis[i] = true;
        for &j in &adj[i] {
            dfs(j, adj, vis, ans);
        }
        ans.push(i);
    }
    let mut ans = vec![];
    let mut vis = vec![false; n];
    for i in 0..n {
        dfs(i, adj, &mut vis, &mut ans);
    }
    ans.reverse();
    ans
}

#[cfg(test)]
mod tests {
    use super::*;

    /// An order is valid iff it is a permutation of `0..n` and every edge
    /// points forward in it.
    fn check(adj: &[Vec<usize>]) {
        let ord = topo_order(adj);
        let n = adj.len();
        let mut pos = vec![usize::MAX; n];
        assert_eq!(n, ord.len(), "adj = {adj:?}");
        for (p, &v) in ord.iter().enumerate() {
            assert_eq!(usize::MAX, pos[v], "duplicate {v} in {ord:?}");
            pos[v] = p;
        }
        for (u, to) in adj.iter().enumerate() {
            for &v in to {
                assert!(pos[u] < pos[v], "{u} -> {v} reversed in {ord:?}");
            }
        }
    }

    #[test]
    fn valid_on_edge_shapes() {
        check(&[]);
        check(&[vec![]]);
        check(&[vec![1, 2], vec![3], vec![3], vec![]]); // diamond
        check(&[vec![], vec![0], vec![], vec![2]]); // isolated nodes
        check(&[vec![3], vec![0], vec![1, 3], vec![]]); // mixed directions
    }

    #[test]
    fn valid_on_random_dags() {
        let mut seed = 1u64;
        for _ in 0..30 {
            seed = seed.wrapping_mul(6364136223846793005).wrapping_add(1);
            let n = (seed % 10) as usize;
            let mut adj = vec![vec![]; n];
            for u in 0..n {
                for v in u + 1..n {
                    seed = seed.wrapping_mul(6364136223846793005).wrapping_add(1);
                    if seed % 3 == 0 {
                        adj[u].push(v);
                    }
                }
            }
            check(&adj);
        }
    }
}
