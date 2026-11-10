use super::multiset::Multiset;
use rand::RngExt;

type Link<T> = Option<Box<Node<T>>>;

#[derive(Debug)]
struct Node<T: Ord> {
    l: Option<Box<Node<T>>>,
    r: Option<Box<Node<T>>>,
    v: T,
    pri: u32,
    rep: usize,
    size: usize,
}

impl<T: Ord> Node<T> {
    fn new(v: T, pri: Option<u32>) -> Self {
        let pri = pri.unwrap_or_else(|| rand::rng().random());
        Node {
            l: None,
            r: None,
            v,
            pri,
            rep: 1,
            size: 1,
        }
    }
    fn update(&mut self) {
        self.size = self.rep;
        if let Some(l) = self.l.as_ref() {
            self.size += l.size;
        }
        if let Some(r) = self.r.as_ref() {
            self.size += r.size;
        }
    }
}

fn merge<T: Ord>(l: Link<T>, r: Link<T>) -> Link<T> {
    if l.is_none() {
        return r;
    }
    if r.is_none() {
        return l;
    }
    let (mut l, mut r) = (l?, r?);
    if l.pri <= r.pri {
        l.r = merge(l.r.take(), Some(r));
        l.update();
        Some(l)
    } else {
        r.l = merge(Some(l), r.l.take());
        r.update();
        Some(r)
    }
}

fn split<T: Ord>(root: Link<T>, on_left: impl Fn(&T) -> bool) -> (Link<T>, Link<T>) {
    if root.is_none() {
        return (None, None);
    }
    let mut root = root.unwrap();
    let (mut l, mut r);
    if on_left(&root.v) {
        (l, r) = split(root.r.take(), on_left);
        root.r = l.take();
        root.update();
        l = Some(root);
    } else {
        (l, r) = split(root.l.take(), on_left);
        root.l = r.take();
        root.update();
        r = Some(root);
    }
    (l, r)
}

// split into 3 parts: <x, =x, >x
fn split_by<T: Ord>(node: Link<T>, x: &T) -> (Link<T>, Link<T>, Link<T>) {
    let (le, gt) = split(node, |v| *v <= *x);
    let (lt, eq) = split(le, |v| *v < *x);
    (lt, eq, gt)
}

fn size<T: Ord>(node: &Link<T>) -> usize {
    node.as_ref().map_or(0, |x| x.size)
}

fn kth<T: Ord>(node: &Link<T>, mut k: usize) -> Option<&T> {
    if node.is_none() || k >= size(node) {
        return None;
    }
    let mut u = node.as_ref();
    while let Some(nd) = u {
        if k < size(&nd.l) {
            u = nd.l.as_ref();
        } else if k < nd.size - size(&nd.r) {
            return Some(&nd.v);
        } else {
            k -= nd.size - size(&nd.r);
            u = nd.r.as_ref();
        }
    }
    None
}

#[derive(Debug)]
pub struct Treap<T: Ord> {
    root: Link<T>,
}

impl<T: Ord> Treap<T> {
    pub fn new() -> Self {
        Self { root: None }
    }

    pub fn len(&self) -> usize {
        size(&self.root)
    }

    pub fn is_empty(&self) -> bool {
        self.len() == 0
    }

    pub fn insert(&mut self, x: T) {
        let (lt, mut eq, gt) = split_by(self.root.take(), &x);
        if let Some(mut nd) = eq {
            nd.rep += 1;
            nd.size += 1;
            eq = Some(nd);
        } else {
            eq = Some(Box::new(Node::new(x, None)));
        }
        self.root = merge(merge(lt, eq), gt);
    }

    /// Removes one occurrence of `x`; returns whether one was present.
    pub fn remove(&mut self, x: &T) -> bool {
        let (lt, mut eq, gt) = split_by(self.root.take(), x);
        let removed = eq.is_some();
        if let Some(mut nd) = eq {
            if nd.rep > 1 {
                nd.rep -= 1;
                nd.size -= 1;
                eq = Some(nd);
            } else {
                eq = merge(nd.l, nd.r);
            }
        }
        self.root = merge(merge(lt, eq), gt);
        removed
    }

    // count(< *x)
    pub fn rank(&self, x: &T) -> usize {
        let mut u = self.root.as_ref();
        let mut cnt = 0;
        while let Some(nd) = u {
            if nd.v < *x {
                cnt += nd.size - size(&nd.r);
                u = nd.r.as_ref();
            } else {
                u = nd.l.as_ref();
            }
        }
        cnt
    }

    // count(<= *x)
    pub fn count_le(&self, x: &T) -> usize {
        let mut u = self.root.as_ref();
        let mut cnt = 0;
        while let Some(nd) = u {
            if nd.v <= *x {
                cnt += nd.size - size(&nd.r);
                u = nd.r.as_ref();
            } else {
                u = nd.l.as_ref();
            }
        }
        cnt
    }

    pub fn kth(&self, k: usize) -> Option<&T> {
        kth(&self.root, k)
    }
}

impl<T: Ord> Default for Treap<T> {
    fn default() -> Self {
        Self::new()
    }
}

impl<T: Ord> Multiset<T> for Treap<T> {
    fn len(&self) -> usize {
        Treap::len(self)
    }

    fn insert(&mut self, x: T) {
        Treap::insert(self, x);
    }

    fn remove(&mut self, x: &T) -> bool {
        Treap::remove(self, x)
    }

    fn rank(&self, x: &T) -> usize {
        Treap::rank(self, x)
    }

    fn count_le(&self, x: &T) -> usize {
        Treap::count_le(self, x)
    }

    fn kth(&self, k: usize) -> Option<&T> {
        Treap::kth(self, k)
    }
}

#[cfg(test)]
mod tests {
    use super::super::multiset::test_util::check;
    use super::*;

    #[test]
    fn new_treap_is_empty() {
        check(&Treap::new(), &[]);
    }

    /// `prev`/`next` are the strict neighbours of `x` in the sorted model.
    fn check_neighbours(t: &Treap<i32>, model: &[i32]) {
        let lo = model.first().copied().unwrap_or(0) - 1;
        let hi = model.last().copied().unwrap_or(0) + 2;
        for x in lo..=hi {
            let want_prev = model.iter().filter(|v| **v < x).max();
            let want_next = model.iter().filter(|v| **v > x).min();
            assert_eq!(want_prev, t.prev(&x), "prev({x}) in {model:?}");
            assert_eq!(want_next, t.next(&x), "next({x}) in {model:?}");
        }
    }

    #[test]
    fn prev_next_on_empty_treap() {
        let t = Treap::new();
        assert_eq!(None, t.prev(&0));
        assert_eq!(None, t.next(&0));
    }

    #[test]
    fn prev_next_are_strict_neighbours() {
        let mut t = Treap::new();
        for x in [1, 3, 5, 7] {
            t.insert(x);
        }
        // Below the minimum / above the maximum.
        assert_eq!(None, t.prev(&1));
        assert_eq!(None, t.prev(&-5));
        assert_eq!(None, t.next(&7));
        assert_eq!(None, t.next(&100));

        // Gaps use the in-between values.
        assert_eq!(Some(&1), t.prev(&2));
        assert_eq!(Some(&3), t.next(&2));

        // A present value never returns itself.
        assert_eq!(Some(&1), t.prev(&3));
        assert_eq!(Some(&5), t.next(&3));

        check_neighbours(&t, &[1, 3, 5, 7]);
    }

    #[test]
    fn prev_next_with_duplicates_and_removals() {
        let mut t = Treap::new();
        for x in [2, 2, 5, 9] {
            t.insert(x);
        }
        assert_eq!(None, t.prev(&2));
        assert_eq!(Some(&5), t.next(&2));
        assert_eq!(Some(&2), t.prev(&5));
        assert_eq!(Some(&5), t.prev(&9));
        check_neighbours(&t, &[2, 2, 5, 9]);

        // Removing one duplicate keeps `prev`/`next` pointing at values, not at
        // a removed slot.
        t.remove(&2);
        check_neighbours(&t, &[2, 5, 9]);
        t.remove(&2);
        check_neighbours(&t, &[5, 9]);
        assert_eq!(None, t.prev(&5));
        assert_eq!(Some(&9), t.next(&5));

        // Drained: nothing left to walk to.
        t.remove(&5);
        t.remove(&9);
        check_neighbours(&t, &[]);
        assert_eq!(None, t.next(&5));
    }

    #[test]
    fn insert_keeps_all_elements() {
        let mut t = Treap::new();
        for x in [5, 1, 3, 2, 4] {
            t.insert(x);
        }
        check(&t, &[1, 2, 3, 4, 5]);
    }

    #[test]
    fn insert_counts_duplicates() {
        let mut t = Treap::new();
        for x in [2, 2, 2, 1] {
            t.insert(x);
        }
        check(&t, &[1, 2, 2, 2]);

        let mut u = Treap::new();
        for _ in 0..3 {
            u.insert(7);
        }
        check(&u, &[7, 7, 7]);
    }

    #[test]
    fn remove_drops_one_occurrence() {
        let mut t = Treap::new();
        for x in [1, 2, 2, 3] {
            t.insert(x);
        }
        t.remove(&2);
        check(&t, &[1, 2, 3]);
        t.remove(&2);
        check(&t, &[1, 3]);
        t.remove(&2); // absent: no change
        check(&t, &[1, 3]);
        t.remove(&1);
        t.remove(&3);
        check(&t, &[]);
    }

    #[test]
    fn mixed_ops_match_sorted_model() {
        let mut model: Vec<i32> = Vec::new();
        let mut t = Treap::new();

        let ops = [
            (true, 4),
            (true, 2),
            (true, 7),
            (true, 2),
            (false, 7),
            (true, 1),
            (true, 9),
            (false, 2),
            (true, 5),
            (false, 4),
        ];
        for (insert, x) in ops {
            if insert {
                t.insert(x);
                let pos = model.partition_point(|v| *v < x);
                model.insert(pos, x);
            } else {
                if let Some(pos) = model.iter().position(|v| *v == x) {
                    model.remove(pos);
                }
                t.remove(&x);
            }
            check(&t, &model);
        }
    }
}
