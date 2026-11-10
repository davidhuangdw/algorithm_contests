#[derive(Debug)]
pub struct BITree {
    sub: Vec<i64>,
}

impl BITree {
    // interval [0, n)
    pub fn new(n: usize) -> Self {
        Self { sub: vec![0; n] }
    }

    pub fn from_slice(s: &[i64]) -> Self {
        let mut sub: Vec<i64> = s
            .iter()
            .scan(0, |s, v| {
                *s += v;
                Some(*s)
            })
            .collect();
        for i in (1..=s.len()).rev() {
            let j = i & (i - 1);
            if j > 0 {
                sub[i - 1] -= sub[j - 1];
            }
        }
        Self { sub }
    }

    pub fn len(&self) -> usize {
        self.sub.len()
    }

    pub fn is_empty(&self) -> bool {
        self.len() == 0
    }

    // sum(a[0..=i])
    pub fn pre_sum(&self, mut i: usize) -> i64 {
        assert!(i < self.len());
        let mut sum = 0;
        i += 1;
        while i > 0 {
            sum += self.sub[i - 1];
            i &= i - 1;
        }
        sum
    }

    pub fn sum(&self, l: usize, r: usize) -> i64 {
        assert!(l <= r && r < self.len());
        self.pre_sum(r) - (if l > 0 { self.pre_sum(l - 1) } else { 0 })
    }

    pub fn get(&self, i: usize) -> i64 {
        self.sum(i, i)
    }

    // a[i]+=v
    pub fn add(&mut self, mut i: usize, v: i64) {
        assert!(i < self.sub.len());
        i += 1;
        while i <= self.sub.len() {
            self.sub[i - 1] += v;
            i += i & (!i + 1)
        }
    }

    // O(n) linear iter all
    pub fn iter(&self) -> std::vec::IntoIter<i64> {
        let mut val = self.sub.clone();
        for i in 0..self.len() {
            let i = i + 1;
            let p = i + (i & (!i + 1));
            if p <= self.len() {
                val[p - 1] -= self.sub[i - 1]
            }
        }
        val.into_iter()
    }
}

impl IntoIterator for &BITree {
    type Item = i64;
    type IntoIter = std::vec::IntoIter<i64>;

    fn into_iter(self) -> Self::IntoIter {
        self.iter()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    /// Checks `t` against `model` through every read API it exposes.
    fn check(t: &BITree, model: &[i64]) {
        assert_eq!(model.len(), t.len());
        assert_eq!(model.is_empty(), t.is_empty());
        assert_eq!(model, t.iter().collect::<Vec<_>>());

        for i in 0..model.len() {
            assert_eq!(
                model[..=i].iter().sum::<i64>(),
                t.pre_sum(i),
                "pre_sum({i})"
            );
            assert_eq!(model[i], t.get(i), "get({i})");
            let mut acc = 0;
            for (r, v) in model.iter().enumerate().skip(i) {
                acc += v;
                assert_eq!(acc, t.sum(i, r), "sum({i},{r})");
            }
        }
    }

    #[test]
    fn new_tree_is_all_zeros() {
        check(&BITree::new(5), &[0; 5]);
    }

    #[test]
    fn from_slice_matches_values() {
        // 8 is a power of two, so include other lengths as well.
        for s in [
            vec![],
            vec![7],
            vec![1, 2, 3],
            vec![3, -1, 4, -1, 5],
            vec![3, -1, 4, -1, 5, 9, -2, 6],
        ] {
            check(&BITree::from_slice(&s), &s);
        }
    }

    #[test]
    fn add_matches_model() {
        let mut model: Vec<i64> = (0..8).collect();
        let mut t = BITree::from_slice(&model);
        for (idx, v) in [(0, 3), (2, 6), (1, 1), (7, 7), (4, 5)] {
            t.add(idx, v);
            model[idx] += v;
            check(&t, &model);
        }
    }
}
