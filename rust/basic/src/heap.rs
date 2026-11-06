#[derive(Debug)]
pub struct Heap<T> {
    data: Vec<T>,
}

// min heap
impl<T: Ord> Heap<T> {
    pub fn new() -> Self {
        Self { data: vec![] }
    }

    pub fn from_vec(v: Vec<T>) -> Self {
        let n = v.len();
        let mut h = Self { data: v };
        for i in (0..n/2).rev() {
            h.sift_down(i);
        }
        h
    }

    fn sift_up(&mut self, i: usize) {
        let mut i = i;
        while i > 0 && self.data[i] < self.data[(i - 1) / 2] {
            self.data.swap(i, (i - 1) / 2);
            i = (i - 1) / 2;
        }
    }

    fn sift_down(&mut self, i: usize) {
        let mut i = i;
        while i * 2 + 1 < self.data.len() {
            let (lc, rc) = (i * 2 + 1, i * 2 + 2);
            let ch = if rc < self.data.len() && self.data[rc] < self.data[lc] {
                rc
            } else {
                lc
            };
            if self.data[i] <= self.data[ch] {
                break;
            }
            self.data.swap(i, ch);
            i = ch;
        }
    }

    pub fn len(&self) -> usize {
        self.data.len()
    }

    pub fn is_empty(&self) -> bool {
        self.len() == 0
    }

    pub fn top(&self) -> Option<&T> {
        self.data.first()
    }

    pub fn push(&mut self, v: T) {
        self.data.push(v);
        self.sift_up(self.data.len() - 1);
    }

    pub fn pop(&mut self) -> Option<T> {
        if self.is_empty() {
            return None;
        }
        let last = self.data.len() - 1;
        self.data.swap(0, last);
        let res = self.data.pop();
        self.sift_down(0);
        res
    }
}

impl<T: Ord> Default for Heap<T> {
    fn default() -> Self {
        Self::new()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn heap(v: &[i32]) -> Heap<i32> {
        let mut h = Heap::new();
        for &x in v {
            h.push(x);
        }
        h
    }

    fn drain(h: &mut Heap<i32>) -> Vec<i32> {
        let mut out = Vec::new();
        while let Some(x) = h.pop() {
            out.push(x);
        }
        out
    }

    #[test]
    fn new_heap_is_empty() {
        let mut h = Heap::<i32>::new();
        assert_eq!(0, h.len());
        assert!(h.is_empty());
        assert_eq!(None, h.top());
        assert_eq!(None, h.pop());
    }

    #[test]
    fn top_is_minimum() {
        let mut h = Heap::<i32>::new();
        h.push(5);
        assert_eq!(Some(&5), h.top());
        h.push(3);
        assert_eq!(Some(&3), h.top());
        h.push(9);
        assert_eq!(Some(&3), h.top());
        assert_eq!(3, h.len());
    }

    #[test]
    fn pop_yields_ascending_order() {
        let mut h = heap(&[5, 3, 8, 1, 9, 2]);
        assert_eq!(vec![1, 2, 3, 5, 8, 9], drain(&mut h));
        assert!(h.is_empty());
    }

    #[test]
    fn pop_returns_minimum_each_time() {
        let mut h = heap(&[4, 1, 3]);
        assert_eq!(Some(1), h.pop());
        assert_eq!(2, h.len());
        assert_eq!(Some(&3), h.top());
        assert_eq!(Some(3), h.pop());
        assert_eq!(Some(4), h.pop());
        assert_eq!(None, h.pop());
    }

    #[test]
    fn handles_duplicates() {
        let mut h = heap(&[2, 2, 1, 1, 3]);
        assert_eq!(vec![1, 1, 2, 2, 3], drain(&mut h));
    }

    #[test]
    fn single_element() {
        let mut h = heap(&[7]);
        assert_eq!(Some(&7), h.top());
        assert_eq!(Some(7), h.pop());
        assert!(h.is_empty());
        assert_eq!(None, h.top());
    }

    #[test]
    fn sift_down_picks_smaller_child() {
        // Root has two children; the smaller one must become the new root.
        let mut h = Heap::<i32>::new();
        h.push(1);
        h.push(3);
        h.push(2);
        assert_eq!(Some(1), h.pop());
        assert_eq!(Some(&2), h.top());
        assert_eq!(Some(2), h.pop());
        assert_eq!(Some(3), h.pop());
    }

    #[test]
    fn sift_down_with_only_left_child() {
        // Two elements: the root has a single left child and no right child.
        let mut h = heap(&[3, 1]);
        assert_eq!(Some(1), h.pop());
        assert_eq!(Some(&3), h.top());
        assert_eq!(Some(3), h.pop());
        assert!(h.is_empty());
    }

    #[test]
    fn from_vec_yields_sorted_order() {
        assert!(Heap::<i32>::from_vec(vec![]).is_empty());

        let mut single = Heap::from_vec(vec![7]);
        assert_eq!(Some(&7), single.top());
        assert_eq!(Some(7), single.pop());
        assert!(single.is_empty());

        let mut h = Heap::from_vec(vec![5, 3, 8, 1, 9, 2]);
        assert_eq!(6, h.len());
        assert_eq!(vec![1, 2, 3, 5, 8, 9], drain(&mut h));
    }

    #[test]
    fn from_vec_matches_push_based_build() {
        let input: Vec<i32> = (0..50).map(|i| (i * 31) % 53).collect();
        let mut from_vec = Heap::from_vec(input.clone());
        let mut from_push = heap(&input);
        assert_eq!(drain(&mut from_push), drain(&mut from_vec));
    }

    #[test]
    fn large_input_matches_sorted() {
        let input: Vec<i32> = (0..100).map(|i| (i * 37) % 101).collect();
        let mut expected = input.clone();
        expected.sort_unstable();

        let mut h = heap(&input);
        assert_eq!(input.len(), h.len());
        assert_eq!(expected, drain(&mut h));
    }

    #[test]
    fn works_with_non_copy_values() {
        let mut h = Heap::<String>::new();
        h.push(String::from("c"));
        h.push(String::from("a"));
        h.push(String::from("b"));
        assert_eq!(Some("a"), h.top().map(String::as_str));
        assert_eq!(Some(String::from("a")), h.pop());
        assert_eq!(Some(String::from("b")), h.pop());
        assert_eq!(Some(String::from("c")), h.pop());
    }
}
