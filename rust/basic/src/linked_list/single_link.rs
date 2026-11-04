use std::fmt::Display;

type Link<T> = Option<Box<ListNode<T>>>;

#[derive(Debug, PartialEq, Eq)]
pub struct ListNode<T = i32> {
    pub val: T,
    pub next: Link<T>,
}

pub struct List<T = i32> {
    head: Link<T>,
}

impl<T> List<T> {
    pub fn from_slice(vec: &[T]) -> Self
    where
        T: Clone,
    {
        let mut head = None;
        for v in vec.iter().rev() {
            head = Some(Box::new(ListNode {
                val: v.clone(),
                next: head,
            }));
        }
        Self { head }
    }

    pub fn merge(a: Self, b: Self) -> Self
    where
        T: PartialOrd,
    {
        let (mut a, mut b) = (a.head, b.head);
        let mut head = None;
        let mut tail = &mut head;
        while a.is_some() && b.is_some() {
            let mut node;
            if a.as_ref().unwrap().val <= b.as_ref().unwrap().val {
                node = a.take().unwrap();
                a = node.next.take();
            } else {
                node = b.take().unwrap();
                b = node.next.take();
            }
            *tail = Some(node);
            tail = &mut tail.as_mut().unwrap().next;
        }
        *tail = a.take().or_else(|| b.take());
        List { head }
    }

    pub fn kth_link(&mut self, k: usize) -> Option<&mut Link<T>> {
        let mut p = &mut self.head;
        for _ in 0..k {
            p = &mut p.as_mut()?.next
        }
        Some(p)
    }

    pub fn split_at(&mut self, k: usize) -> Option<Link<T>> {
        Some(self.kth_link(k)?.take())
    }

    pub fn rotate(&mut self, k: usize) -> bool {
        if k == 0 {
            return true;
        }
        if let Some(mut r) = self.split_at(k)
            && r.is_some()
        {
            let l = self.head.take();
            let mut tail = &mut r;
            while let Some(node) = tail {
                tail = &mut node.next;
            }
            *tail = l;
            self.head = r;
            true
        } else {
            false
        }
    }

    pub fn middle(&self) -> Option<&T> {
        let mut a = self.head.as_ref()?;
        let mut b = a.next.as_ref();
        while let Some(bx) = b
            && bx.next.is_some()
        {
            a = a.next.as_ref()?;
            b = bx.next.as_ref()?.next.as_ref();
        }
        Some(&a.val)
    }

    pub fn iter(&self) -> impl Iterator<Item = &ListNode<T>> {
        std::iter::successors(self.head.as_deref(), |node| node.next.as_deref())
    }

    pub fn to_vec(&self) -> Vec<T>
    where
        T: Clone,
    {
        self.iter().map(|node| node.val.clone()).collect()
    }

    pub fn reverse(&mut self) {
        let mut p = self.head.take();
        while let Some(mut first) = p {
            p = first.next.take();
            first.next = self.head.take();
            self.head = Some(first);
        }
    }

    pub fn is_empty(&self) -> bool {
        self.head.is_none()
    }

    pub fn len(&self) -> usize {
        self.iter().count()
    }

    pub fn peek(&self) -> Option<&T> {
        self.head.as_deref().map(|x| &x.val)
    }

    pub fn clear(&mut self) {
        self.head = None;
    }

    pub fn push_front(&mut self, v: T) {
        self.head = Some(Box::new(ListNode {
            val: v,
            next: self.head.take(),
        }));
    }

    pub fn pop_front(&mut self) -> Option<T> {
        let head = self.head.take()?;
        self.head = head.next;
        Some(head.val)
    }

    pub fn push_back(&mut self, v: T) {
        let mut p = &mut self.head;
        while let Some(node) = p {
            p = &mut node.next;
        }
        p.replace(Box::new(ListNode { val: v, next: None }));
    }

    pub fn pop_back(&mut self) -> Option<T> {
        let mut p = &mut self.head;
        while p.as_ref()?.next.is_some() {
            p = &mut p.as_mut()?.next;
        }
        p.take().map(|p| p.val)
    }

    pub fn remove(&mut self, k: usize) -> Option<T> {
        let kth = self.kth_link(k)?;
        let cut = kth.take()?;
        *kth = cut.next;
        Some(cut.val)
    }

    pub fn insert(&mut self, k: usize, v: T) -> bool {
        let Some(link) = self.kth_link(k) else {
            return false;
        };
        let next = link.take();
        *link = Some(Box::new(ListNode { val: v, next }));
        true
    }
}

impl<T: Display> Display for List<T> {
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

    fn list(v: &[i32]) -> List {
        List::from_slice(v)
    }

    #[test]
    fn empty_list_reports_nothing() {
        let mut l = list(&[]);
        assert!(l.is_empty());
        assert_eq!(0, l.len());
        assert_eq!(None, l.peek());
        assert_eq!(None, l.pop_front());
        assert_eq!(None, l.pop_back());
        assert_eq!(None, l.remove(0));
        assert_eq!("None", l.to_string());
    }

    #[test]
    fn from_slice_builds_in_order() {
        let l = list(&[-2, 3, 3, 5, 0]);
        assert_eq!(vec![-2, 3, 3, 5, 0], l.to_vec());
        assert_eq!(5, l.len());
        assert!(!l.is_empty());
        assert_eq!(Some(&-2), l.peek());
    }

    #[test]
    fn push_and_pop_front_are_lifo() {
        let mut l = list(&[2, 3]);
        l.push_front(1);
        assert_eq!(vec![1, 2, 3], l.to_vec());
        assert_eq!(Some(1), l.pop_front());
        assert_eq!(vec![2, 3], l.to_vec());
    }

    #[test]
    fn push_and_pop_back_are_fifo() {
        let mut l = list(&[]);
        l.push_back(1);
        l.push_back(2);
        assert_eq!(vec![1, 2], l.to_vec());
        assert_eq!(Some(2), l.pop_back());
        assert_eq!(Some(1), l.pop_back());
        assert!(l.pop_back().is_none());
    }

    #[test]
    fn insert_and_remove_by_index() {
        let mut l = list(&[1, 3]);
        assert!(l.insert(1, 2));
        assert!(l.insert(3, 4));
        assert!(!l.insert(99, 0));
        assert_eq!(vec![1, 2, 3, 4], l.to_vec());

        assert_eq!(Some(2), l.remove(1));
        assert_eq!(Some(1), l.remove(0));
        assert_eq!(None, l.remove(99));
        assert_eq!(vec![3, 4], l.to_vec());
    }

    #[test]
    fn remove_covers_tail_and_empties_list() {
        let mut l = list(&[1, 2, 3]);
        assert_eq!(Some(3), l.remove(2));
        assert_eq!(vec![1, 2], l.to_vec());

        let mut single = list(&[7]);
        assert_eq!(Some(7), single.remove(0));
        assert!(single.is_empty());

        let mut drained = list(&[1, 2, 3]);
        let mut removed = Vec::new();
        while let Some(v) = drained.remove(0) {
            removed.push(v);
        }
        assert_eq!(vec![1, 2, 3], removed);
        assert!(drained.is_empty());
    }

    #[test]
    fn insert_covers_head_and_empty_list() {
        let mut l = list(&[2, 3]);
        assert!(l.insert(0, 1));
        assert_eq!(vec![1, 2, 3], l.to_vec());

        let mut empty = list(&[]);
        assert!(empty.insert(0, 9));
        assert_eq!(vec![9], empty.to_vec());

        // One past the end is the first rejected index.
        let mut single = list(&[1]);
        assert!(!single.insert(2, 9));
        assert_eq!(vec![1], single.to_vec());
    }

    #[test]
    fn merge_keeps_sorted_order() {
        let merged = List::merge(list(&[1, 3, 5]), list(&[2, 3, 4]));
        assert_eq!(vec![1, 2, 3, 3, 4, 5], merged.to_vec());

        let merged = List::merge(list(&[]), list(&[1, 2]));
        assert_eq!(vec![1, 2], merged.to_vec());

        let merged = List::merge(list(&[1, 2]), list(&[]));
        assert_eq!(vec![1, 2], merged.to_vec());

        assert!(List::merge(list(&[]), list(&[])).is_empty());
    }

    #[test]
    fn kth_link_reaches_slot_up_to_len() {
        let mut l = list(&[1, 2, 3]);
        assert_eq!(Some(1), l.kth_link(0).unwrap().as_deref().map(|n| n.val));
        assert_eq!(Some(3), l.kth_link(2).unwrap().as_deref().map(|n| n.val));
        // The one-past-end slot still exists, which is what makes appending possible.
        assert!(l.kth_link(3).is_some_and(|link| link.is_none()));
        assert!(l.kth_link(4).is_none());
    }

    #[test]
    fn split_at_detaches_tail() {
        let mut l = list(&[1, 2, 3, 4]);
        let tail = List {
            head: l.split_at(2).unwrap(),
        };
        assert_eq!(vec![1, 2], l.to_vec());
        assert_eq!(vec![3, 4], tail.to_vec());

        // Splitting at 0 keeps everything in the detached part.
        let mut l = list(&[1, 2]);
        let all = List {
            head: l.split_at(0).unwrap(),
        };
        assert!(l.is_empty());
        assert_eq!(vec![1, 2], all.to_vec());

        assert!(l.split_at(99).is_none());
    }

    #[test]
    fn rotate_moves_prefix_to_back() {
        let mut l = list(&[1, 2, 3, 4, 5]);
        assert!(l.rotate(2));
        assert_eq!(vec![3, 4, 5, 1, 2], l.to_vec());

        // Rotating by 0 is a no-op that still succeeds.
        assert!(l.rotate(0));
        assert_eq!(vec![3, 4, 5, 1, 2], l.to_vec());

        // Rotating by the full length detaches nothing, so it reports failure
        // and leaves the list untouched.
        assert!(!l.rotate(5));
        assert_eq!(vec![3, 4, 5, 1, 2], l.to_vec());

        // Beyond the end fails and leaves the list untouched.
        assert!(!l.rotate(6));
        assert_eq!(vec![3, 4, 5, 1, 2], l.to_vec());

        let mut empty = list(&[]);
        assert!(empty.rotate(0));
        assert!(empty.is_empty());
    }

    #[test]
    fn works_with_non_copy_values() {
        let mut l = List::from_slice(&[String::from("b"), String::from("c")]);
        l.push_front(String::from("a"));
        assert_eq!("a -> b -> c -> None", l.to_string());
        assert_eq!(Some("a"), l.peek().map(String::as_str));
        assert_eq!(Some(String::from("a")), l.pop_front());
        assert_eq!(vec![String::from("b"), String::from("c")], l.to_vec());

        let merged = List::merge(l, List::from_slice(&[String::from("a")]));
        assert_eq!(
            vec!["a", "b", "c"],
            merged.iter().map(|n| n.val.as_str()).collect::<Vec<_>>()
        );
    }

    #[test]
    fn middle_returns_lower_middle() {
        assert_eq!(None, list(&[]).middle());
        assert_eq!(Some(&1), list(&[1]).middle());
        assert_eq!(Some(&1), list(&[1, 2]).middle());
        assert_eq!(Some(&2), list(&[1, 2, 3]).middle());
        assert_eq!(Some(&2), list(&[1, 2, 3, 4]).middle());
        assert_eq!(Some(&3), list(&[1, 2, 3, 4, 5]).middle());
    }

    #[test]
    fn clear_resets_list() {
        let mut l = list(&[1, 2, 3]);
        assert_eq!("1 -> 2 -> 3 -> None", l.to_string());
        l.clear();
        assert!(l.is_empty());
        assert_eq!("None", l.to_string());
    }
    #[test]
    fn reverse_flips_order_in_place() {
        let mut l = list(&[1, 2, 3]);
        l.reverse();
        assert_eq!(vec![3, 2, 1], l.to_vec());
        assert_eq!(3, l.len());
        assert_eq!(Some(&3), l.peek());
        // The old head is now the tail and must not dangle.
        assert!(l.iter().last().is_some_and(|node| node.next.is_none()));

        l.reverse();
        assert_eq!(vec![1, 2, 3], l.to_vec()); // Reversing twice restores the original order.

        let mut even = list(&[1, 2, 3, 4]);
        even.reverse();
        assert_eq!(vec![4, 3, 2, 1], even.to_vec()); // Even length.

        let mut single = list(&[7]);
        single.reverse();
        assert_eq!(vec![7], single.to_vec()); // Single node.

        let mut empty = list(&[]);
        empty.reverse();
        assert!(empty.is_empty());
        assert_eq!("None", empty.to_string()); // None of the above may panic.
    }
}
