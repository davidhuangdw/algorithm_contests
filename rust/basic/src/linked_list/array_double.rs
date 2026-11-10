use crate::linked_list::Seq;
use std::fmt::Display;

#[derive(Debug)]
pub struct ANode<T> {
    val: T,
    prev: usize,
    next: usize,
}

#[derive(Debug)]
pub struct AList<T> {
    nodes: Vec<Option<ANode<T>>>,
    holes: Vec<usize>,
    head: usize, // dummy node
    tail: usize,
}

impl<T> ANode<T> {
    pub fn new(val: T) -> Self {
        Self {
            val,
            prev: 0,
            next: 0,
        }
    }
}

impl<T> AList<T> {
    fn alloc(&mut self, v: T) -> usize {
        let node = ANode::new(v);
        if let Some(i) = self.holes.pop() {
            self.nodes[i] = Some(node);
            i
        } else {
            self.nodes.push(Some(node));
            self.nodes.len() - 1
        }
    }
    fn free(&mut self, i: usize) -> Option<ANode<T>> {
        let node = self.nodes[i].take()?;
        self.holes.push(i);
        Some(node)
    }
    fn get_mut(&mut self, i: usize) -> &mut ANode<T> {
        self.nodes[i].as_mut().unwrap()
    }
    fn get(&self, i: usize) -> &ANode<T> {
        self.nodes[i].as_ref().unwrap()
    }
    pub fn get_valid(&self, i: usize) -> Option<&ANode<T>> {
        if i == self.head || i == self.tail {
            None
        } else {
            self.nodes.get(i)?.as_ref()
            // Some(self.get(i))
        }
    }

    pub fn new() -> Self
    where
        T: Default,
    {
        let mut list = Self {
            nodes: vec![],
            holes: vec![],
            head: 0,
            tail: 0,
        };
        let (head, tail) = (list.alloc(T::default()), list.alloc(T::default()));
        (list.head, list.tail) = (head, tail);
        list.get_mut(head).next = tail;
        list.get_mut(tail).prev = head;
        list
    }

    pub fn iter(&self) -> impl Iterator<Item = &ANode<T>> {
        let first = self.get(self.head).next;
        std::iter::successors(self.get_valid(first), |node| self.get_valid(node.next))
    }

    pub fn len(&self) -> usize {
        self.nodes.len() - self.holes.len() - 2 // exclude head&tail dummy-nodes
    }
    pub fn is_empty(&self) -> bool {
        self.len() == 0
    }

    fn insert_after(&mut self, i: usize, v: T) {
        let cur = self.alloc(v);
        let prev = i;
        let next = self.get(prev).next;

        let node = self.get_mut(cur);
        (node.next, node.prev) = (next, prev);

        self.get_mut(prev).next = cur;
        self.get_mut(next).prev = cur;
    }

    pub fn push_front(&mut self, v: T) {
        self.insert_after(self.head, v);
    }

    pub fn push_back(&mut self, v: T) {
        self.insert_after(self.get(self.tail).prev, v);
    }

    fn remove(&mut self, i: usize) -> Option<T> {
        let node = self.get_valid(i)?;
        let (prev, next) = (node.prev, node.next);
        self.get_mut(prev).next = next;
        self.get_mut(next).prev = prev;
        self.free(i).map(|x| x.val)
    }

    pub fn pop_front(&mut self) -> Option<T> {
        self.remove(self.get(self.head).next)
    }
    pub fn pop_back(&mut self) -> Option<T> {
        self.remove(self.get(self.tail).prev)
    }
}

impl<T> Seq<T> for AList<T> {
    fn new() -> Self
    where
        T: Default,
    {
        AList::new()
    }

    fn len(&self) -> usize {
        AList::len(self)
    }

    fn push_front(&mut self, v: T) {
        AList::push_front(self, v);
    }

    fn pop_front(&mut self) -> Option<T> {
        AList::pop_front(self)
    }
    
    fn push_back(&mut self, v: T) {
        AList::push_back(self, v);
    }
    
    fn pop_back(&mut self) -> Option<T> {
        AList::pop_back(self)
    }

    fn iter(&self) -> Box<dyn Iterator<Item = &T> + '_> {
        Box::new(AList::iter(self).map(|node| &node.val))
    }
}

impl<T: Default> Default for AList<T> {
    fn default() -> Self {
        Self::new()
    }
}

impl<T: Display> Display for AList<T> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for p in self.iter() {
            write!(f, "{} -> ", p.val)?;
        }
        write!(f, "None")?;
        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn list(v: &[i32]) -> AList<i32> {
        let mut l = AList::new();
        for &x in v {
            l.push_back(x);
        }
        l
    }

    fn values(l: &AList<i32>) -> Vec<i32> {
        l.iter().map(|node| node.val).collect()
    }

    #[test]
    fn new_list_is_empty() {
        let mut l = AList::<i32>::new();
        assert_eq!(0, l.len());
        assert!(l.is_empty());
        assert_eq!(Vec::<i32>::new(), values(&l));
        assert_eq!("None", l.to_string());
        assert_eq!(None, l.pop_front());
        assert_eq!(None, l.pop_back());
    }

    #[test]
    fn satisfies_the_front_seq_contract() {
        use crate::linked_list::test_util::check;

        let mut l = AList::<i32>::new();
        for v in [2, 3] {
            l.push_back(v);
        }
        check(&l, &[2, 3]);

        Seq::push_front(&mut l, 1);
        check(&l, &[1, 2, 3]);

        assert_eq!(Some(1), Seq::pop_front(&mut l));
        check(&l, &[2, 3]);
        assert_eq!(Some(2), Seq::pop_front(&mut l));
        check(&l, &[3]);
    }

    #[test]
    fn push_prepends_or_appends() {
        let mut l = AList::<i32>::new();
        l.push_back(2);
        l.push_front(1);
        l.push_back(3);
        l.push_front(0);
        assert_eq!(vec![0, 1, 2, 3], values(&l));
        assert_eq!(4, l.len());
        assert_eq!("0 -> 1 -> 2 -> 3 -> None", l.to_string());

        // Negative values must not be confused with sentinel indices.
        assert_eq!("-1 -> 0 -> 7 -> None", list(&[-1, 0, 7]).to_string());
    }

    #[test]
    fn pop_removes_from_the_requested_end() {
        let mut l = list(&[1, 2, 3]);
        assert_eq!(Some(1), l.pop_front());
        assert_eq!(Some(3), l.pop_back());
        assert_eq!(vec![2], values(&l));
        assert_eq!(1, l.len());

        assert_eq!(Some(2), l.pop_front());
        assert!(l.is_empty());
        assert_eq!(None, l.pop_front());
        assert_eq!(None, l.pop_back());
        assert_eq!("None", l.to_string());
    }

    #[test]
    fn freed_slots_are_reused_without_growing_arena() {
        let mut l = list(&[1, 2, 3]);
        let arena_len = l.nodes.len();

        l.pop_front();
        assert_eq!(2, l.len());
        l.pop_back();
        l.pop_back();
        assert_eq!(0, l.len());
        assert!(l.is_empty());

        l.push_front(9);
        assert_eq!(vec![9], values(&l));
        assert_eq!(1, l.len());
        assert_eq!(arena_len, l.nodes.len());
    }

    #[test]
    fn get_valid_skips_dummies_only() {
        let l = list(&[5]);
        assert!(l.get_valid(l.head).is_none());
        assert!(l.get_valid(l.tail).is_none());
        let first = l.get(l.head).next;
        assert_eq!(Some(&5), l.get_valid(first).map(|node| &node.val));
    }

    #[test]
    fn get_valid_out_of_range_returns_none() {
        assert!(list(&[1]).get_valid(999).is_none());
    }

    #[test]
    fn get_valid_returns_none_for_freed_slot() {
        let mut l = list(&[1, 2]);
        l.pop_front();
        assert!(l.get_valid(l.holes[0]).is_none());
    }

    #[test]
    fn mixed_operations_match_vec_model() {
        let mut l = AList::<i32>::new();
        let mut model: Vec<i32> = Vec::new();

        for (i, val) in (0..12).enumerate() {
            let front = i.is_multiple_of(2);
            if front {
                l.push_front(val);
                model.insert(0, val);
            } else {
                l.push_back(val);
                model.push(val);
            }
            assert_eq!(model, values(&l));
            assert_eq!(model.len(), l.len());
        }

        while !model.is_empty() {
            let pop_back = model.len().is_multiple_of(2);
            let expected = if pop_back {
                model.pop()
            } else {
                Some(model.remove(0))
            };
            let actual = if pop_back {
                l.pop_back()
            } else {
                l.pop_front()
            };
            assert_eq!(expected, actual);
            assert_eq!(model, values(&l));
            assert_eq!(model.len(), l.len());
        }
    }

    #[test]
    fn works_with_non_copy_values() {
        let mut l = AList::<String>::new();
        l.push_back(String::from("a"));
        l.push_front(String::from("b"));
        assert_eq!("b -> a -> None", l.to_string());
        assert_eq!(Some(String::from("b")), l.pop_front());
    }
}
