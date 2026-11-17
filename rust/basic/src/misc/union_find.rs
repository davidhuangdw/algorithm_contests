#[derive(Debug)]
pub struct UnionFindSet {
    fa: Vec<usize>,
    size: Vec<usize>,
}

impl UnionFindSet {
    pub fn new(n: usize) -> Self {
        let fa = (0..n).collect();
        let size = vec![1; n];
        UnionFindSet { fa, size }
    }

    // find root
    pub fn find(&mut self, i: usize) -> usize {
        if self.fa[i] != i {
            self.fa[i] = self.find(self.fa[i]);
        }
        self.fa[i]
    }

    pub fn union(&mut self, i: usize, j: usize) {
        let ri = self.find(i);
        let rj = self.find(j);
        if ri == rj {
            return;
        }
        if self.size[ri] >= self.size[rj] {
            self.fa[rj] = ri;
            self.size[ri] += self.size[rj];
        } else {
            self.fa[ri] = rj;
            self.size[rj] += self.size[ri];
        }
    }

    pub fn same_set(&mut self, i: usize, j: usize) -> bool {
        self.find(i) == self.find(j)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn basic_ops() {
        let mut u = UnionFindSet::new(5);
        assert_eq!((0..5).collect::<Vec<_>>(), u.fa);
        assert!(u.same_set(2, 2) && !u.same_set(0, 4));
        for k in 1..5 {
            u.union(k, k - 1);
        }
        let r = u.find(0);
        for i in 0..5 {
            assert_eq!(r, u.find(i));
            assert_eq!(r, u.fa[i], "fa[{i}] not compressed");
        }
        assert!(u.same_set(0, 4));
    }

    #[test]
    fn matches_naive_model() {
        let mut seed = 1u64;
        for _ in 0..20 {
            seed = seed.wrapping_mul(6364136223846793005).wrapping_add(1);
            let n = (seed % 12) as usize + 1;
            let (mut u, mut lbl) = (UnionFindSet::new(n), (0..n).collect::<Vec<_>>());
            for _ in 0..2 * n {
                seed = seed.wrapping_mul(6364136223846793005).wrapping_add(1);
                let (i, j) = ((seed % n as u64) as usize, (seed >> 32) as usize % n);
                u.union(i, j);
                let (old, new) = (lbl[j], lbl[i]);
                for l in &mut lbl {
                    if *l == old {
                        *l = new;
                    }
                }
            }
            for i in 0..n {
                for j in 0..n {
                    assert_eq!(lbl[i] == lbl[j], u.same_set(i, j), "{i},{j} in {n}");
                }
            }
        }
    }

    #[test]
    fn union_keeps_the_larger_root() {
        let mut u = UnionFindSet::new(6);
        for (a, b) in [(0, 1), (0, 2), (3, 4)] {
            u.union(a, b);
        }
        let big = u.find(0); // {0, 1, 2}
        u.union(1, 3); // 3 nodes vs 2 nodes
        assert_eq!(big, u.find(1));
        let r = u.find(1);
        assert_eq!(5, u.size[r]);
    }
}
