#[derive(Debug)]
pub struct SegTree {
    nodes: Vec<Node>,
    n: usize,
}

#[derive(Debug, Default, Clone, Copy)]
struct Node {
    sum: i64,
    lazy: Option<i64>,
}

impl Node {
    fn add_lazy(&mut self, add: i64, l: usize, r: usize) {
        self.lazy = Some(self.lazy.unwrap_or(0) + add);
        self.sum += add * (r + 1 - l) as i64;
    }
    fn merge(a: &Node, b: &Node) -> Node {
        let sum = a.sum + b.sum;
        Node { sum, lazy: None }
    }
}

impl SegTree {
    pub fn new(n: usize) -> Self {
        let nodes = vec![Node::default(); n * 4];
        SegTree { nodes, n }
    }

    pub fn from_slice(s: &[i64]) -> Self {
        let n = s.len();
        let mut nodes = vec![Node::default(); n * 4];
        fn build(nodes: &mut Vec<Node>, s: &[i64], i: usize, il: usize, ir: usize) {
            if il == ir {
                nodes[i] = Node {
                    sum: s[il],
                    lazy: None,
                };
                return;
            }
            let (lc, rc, md) = (i * 2, i * 2 + 1, (il + ir) / 2);
            build(nodes, s, lc, il, md);
            build(nodes, s, rc, md + 1, ir);
            nodes[i] = Node::merge(&nodes[lc], &nodes[rc]);
        }
        build(&mut nodes, s, 1, 0, n - 1);
        Self { nodes, n }
    }

    pub fn query(&mut self, l: usize, r: usize) -> i64 {
        self._query(1, 0, self.n - 1, l, r)
    }

    pub fn add(&mut self, l: usize, r: usize, v: i64) {
        self._add(1, 0, self.n - 1, l, r, v);
    }

    fn _query(&mut self, i: usize, il: usize, ir: usize, l: usize, r: usize) -> i64 {
        if r < il || ir < l {
            return 0;
        }
        if l <= il && ir <= r {
            return self.nodes[i].sum;
        }
        let (lc, rc, md) = (i * 2, i * 2 + 1, (il + ir) / 2);
        self.push_down(i, il, ir);
        self._query(lc, il, md, l, r) + self._query(rc, md + 1, ir, l, r)
    }

    fn _add(&mut self, i: usize, il: usize, ir: usize, l: usize, r: usize, v: i64) {
        if r < il || ir < l {
            return;
        }
        if l <= il && ir <= r {
            self.nodes[i].add_lazy(v, il, ir);
            return;
        }
        self.push_down(i, il, ir);
        let (lc, rc, md) = (i * 2, i * 2 + 1, (il + ir) / 2);
        self._add(lc, il, md, l, r, v);
        self._add(rc, md + 1, ir, l, r, v);
        self.nodes[i].sum = self.nodes[lc].sum + self.nodes[rc].sum;
    }

    fn push_down(&mut self, i: usize, il: usize, ir: usize) {
        if il < ir
            && let Some(lazy) = self.nodes[i].lazy
        {
            let (lc, rc, md) = (i * 2, i * 2 + 1, (il + ir) / 2);
            self.nodes[lc].add_lazy(lazy, il, md);
            self.nodes[rc].add_lazy(lazy, md + 1, ir);
            self.nodes[i].lazy.take();
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn naive_sum(s: &[i64], l: usize, r: usize) -> i64 {
        s[l..=r].iter().sum()
    }

    #[test]
    fn new_tree_sums_to_zero() {
        let mut t = SegTree::new(4);
        assert_eq!(0, t.query(0, 3));
        assert_eq!(0, t.query(1, 2));
        assert_eq!(0, t.query(2, 2));
    }

    #[test]
    fn from_slice_point_queries() {
        let s = [1, 2, 3, 4];
        let mut t = SegTree::from_slice(&s);
        for (i, want) in s.iter().enumerate() {
            assert_eq!(*want, t.query(i, i), "point query at {i}");
        }
    }

    #[test]
    fn from_slice_range_queries() {
        let s = [1, 2, 3];
        let mut t = SegTree::from_slice(&s);
        assert_eq!(naive_sum(&s, 0, 2), t.query(0, 2));
        assert_eq!(naive_sum(&s, 0, 1), t.query(0, 1));
        assert_eq!(naive_sum(&s, 1, 2), t.query(1, 2));
    }

    #[test]
    fn from_slice_matches_naive_on_larger_input() {
        let s: Vec<i64> = (0..10).map(|i| (i * 7 % 11) as i64).collect();
        let mut t = SegTree::from_slice(&s);
        for l in 0..s.len() {
            for r in l..s.len() {
                assert_eq!(naive_sum(&s, l, r), t.query(l, r), "range {l}..={r}");
            }
        }
    }

    #[test]
    fn add_updates_range_sum() {
        let s = [1, 2, 3, 4];
        let mut t = SegTree::from_slice(&s);
        t.add(1, 2, 10);
        assert_eq!(1, t.query(0, 0));
        assert_eq!(12, t.query(1, 1));
        assert_eq!(13, t.query(2, 2));
        assert_eq!(4, t.query(3, 3));
        assert_eq!(30, t.query(0, 3));
    }

    #[test]
    fn add_on_empty_tree_sets_values() {
        let mut t = SegTree::new(4);
        t.add(0, 3, 1);
        assert_eq!(4, t.query(0, 3));
        assert_eq!(1, t.query(2, 2));
    }

    #[test]
    fn add_matches_naive_model() {
        let n = 8;
        let mut model: Vec<i64> = (0..n as i64).collect();
        let mut t = SegTree::from_slice(&model);

        let ops = [(0, 3, 5), (2, 6, -2), (1, 1, 7), (0, 7, 1), (4, 5, 3)];
        for (l, r, v) in ops {
            t.add(l, r, v);
            for x in &mut model[l..=r] {
                *x += v;
            }
            for a in 0..n {
                for b in a..n {
                    assert_eq!(
                        naive_sum(&model, a, b),
                        t.query(a, b),
                        "after add({l},{r},{v}) on {a}..={b}"
                    );
                }
            }
        }
    }
}
