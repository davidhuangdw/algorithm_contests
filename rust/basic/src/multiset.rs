//! Ordered multiset: a set that keeps duplicate occurrences.
//!
//! Any ordered structure (treap, splay tree, sorted vec, ...) can expose the
//! same API by implementing the primitive methods below; the rest are derived.

/// Ordered multiset interface.
///
/// Implementors provide the primitives; `is_empty`, `count`, `contains`,
/// `min`, `max`, `prev` and `next` come for free.
pub trait Multiset<T: Ord> {
    /// Number of stored occurrences.
    fn len(&self) -> usize;

    /// Inserts one occurrence of `x`.
    fn insert(&mut self, x: T);

    /// Removes one occurrence of `x`; returns whether one was present.
    fn remove(&mut self, x: &T) -> bool;

    /// Number of occurrences strictly less than `x`.
    fn rank(&self, x: &T) -> usize;

    /// Number of occurrences less than or equal to `x`.
    fn count_le(&self, x: &T) -> usize;

    /// The `k`-th smallest occurrence, 0-based; `None` when `k >= len()`.
    fn kth(&self, k: usize) -> Option<&T>;

    fn is_empty(&self) -> bool {
        self.len() == 0
    }

    fn count(&self, x: &T) -> usize {
        self.count_le(x) - self.rank(x)
    }

    fn contains(&self, x: &T) -> bool {
        self.count(x) > 0
    }

    fn min(&self) -> Option<&T> {
        self.kth(0)
    }

    fn max(&self) -> Option<&T> {
        self.len().checked_sub(1).and_then(|k| self.kth(k))
    }

    // max <x
    fn prev(&self, x: &T) -> Option<&T> {
        self.rank(x).checked_sub(1).and_then(|k| self.kth(k))
    }

    // min >x
    fn next(&self, x: &T) -> Option<&T> {
        self.kth(self.count_le(x))
    }
}

#[cfg(test)]
pub(crate) mod test_util {
    use super::Multiset;

    /// Cross-checks a multiset against a sorted `model` through every read API.
    pub fn check<M: Multiset<i32>>(m: &M, model: &[i32]) {
        assert_eq!(model.len(), m.len(), "len");
        assert_eq!(model.is_empty(), m.is_empty(), "is_empty");
        assert_eq!(model.first(), m.min(), "min");
        assert_eq!(model.last(), m.max(), "max");

        for (i, want) in model.iter().enumerate() {
            assert_eq!(Some(want), m.kth(i), "kth({i})");
        }
        assert_eq!(None, m.kth(model.len()), "kth(len)");

        let lo = model.first().copied().unwrap_or(0) - 1;
        let hi = model.last().copied().unwrap_or(0) + 2;
        for x in lo..=hi {
            let below = model.iter().filter(|v| **v < x).count();
            let not_above = model.iter().filter(|v| **v <= x).count();
            assert_eq!(below, m.rank(&x), "rank({x})");
            assert_eq!(not_above, m.count_le(&x), "count_le({x})");
            assert_eq!(not_above - below, m.count(&x), "count({x})");
            assert_eq!(not_above > below, m.contains(&x), "contains({x})");

            let want_prev = model.iter().filter(|v| **v < x).max();
            let want_next = model.iter().filter(|v| **v > x).min();
            assert_eq!(want_prev, m.prev(&x), "prev({x})");
            assert_eq!(want_next, m.next(&x), "next({x})");
        }
    }
}
