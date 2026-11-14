use std::{
    cell::RefCell,
    fmt::{Debug, Display},
    rc::{Rc, Weak},
};

#[derive(Debug, Default)]
pub struct DNode<T = i32> {
    pub val: T,
    pub next: Option<Rc<RefCell<DNode<T>>>>,
    pub prev: Option<Weak<RefCell<DNode<T>>>>,
}
type DLink<T> = Rc<RefCell<DNode<T>>>;

#[derive(Debug)]
pub struct DList<T = i32> {
    head: Rc<RefCell<DNode<T>>>, // dummy
    tail: Rc<RefCell<DNode<T>>>,
    size: usize,
}

impl<T> DList<T> {
    pub fn new() -> Self
    where
        T: Default,
    {
        let head = Rc::new(RefCell::new(DNode::default()));
        let tail = Rc::new(RefCell::new(DNode::default()));
        head.borrow_mut().next = Some(tail.clone());
        tail.borrow_mut().prev = Some(Rc::downgrade(&head));
        let size = 0;
        DList { head, tail, size }
    }

    pub fn from_slice(s: &[T]) -> Self
    where
        T: Clone + Default,
    {
        let mut l = DList::new();
        for val in s.iter().rev() {
            l.push_front(val.clone());
        }
        l
    }
    pub fn iter(&self) -> impl Iterator<Item = DLink<T>> {
        let first = next_link(&self.head).filter(|node| !Rc::ptr_eq(node, &self.tail));
        std::iter::successors(first, |node| {
            next_link(node).filter(|nxt| !Rc::ptr_eq(nxt, &self.tail))
        })
    }

    pub fn len(&self) -> usize {
        self.size
    }

    pub fn is_empty(&self) -> bool {
        self.size == 0
    }

    pub fn push_front(&mut self, val: T) {
        insert_after(&self.head, val);
        self.size += 1;
    }

    pub fn pop_front(&mut self) -> Option<T>
    where
        T: Debug,
    {
        if self.is_empty() {
            return None;
        }
        let mut cur = next_link(&self.head).unwrap();
        self.pop_link(&mut cur);
        Some(take_link(cur))
    }

    pub fn push_back(&mut self, val: T) {
        let last = prev_link(&self.tail).unwrap();
        insert_after(&last, val);
        self.size += 1;
    }

    pub fn pop_back(&mut self) -> Option<T>
    where
        T: Debug,
    {
        if self.is_empty() {
            return None;
        }
        let mut cur = prev_link(&self.tail).unwrap();
        self.pop_link(&mut cur);
        Some(take_link(cur))
    }

    fn pop_link(&mut self, cur: &mut DLink<T>) {
        let prev = prev_link(&cur).unwrap();
        let next = next_link(&cur).unwrap();
        prev.borrow_mut().next = Some(next.clone());
        next.borrow_mut().prev = Some(Rc::downgrade(&prev));
        self.size -= 1;
    }
}

fn insert_after<T>(link: &DLink<T>, val: T) {
    let prev = link;
    let next = next_link(link).unwrap();
    let cur = Rc::new(RefCell::new(DNode {
        val,
        next: Some(next.clone()),
        prev: Some(Rc::downgrade(link)),
    }));
    next.borrow_mut().prev = Some(Rc::downgrade(&cur));
    prev.borrow_mut().next = Some(cur);
}

fn next_link<T>(l: &DLink<T>) -> Option<DLink<T>> {
    l.borrow().next.clone()
}

fn prev_link<T>(l: &DLink<T>) -> Option<DLink<T>> {
    l.borrow().prev.as_ref()?.upgrade().clone()
}

fn take_link<T>(l: DLink<T>) -> T
where
    T: Debug,
{
    match Rc::try_unwrap(l) {
        Ok(cell) => cell.into_inner().val,
        Err(err) => panic!("take_dlink: {:?}", err),
    }
}

impl<T: Display> Display for DList<T> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for p in self.iter() {
            write!(f, "{} -> ", p.borrow().val)?;
        }
        write!(f, "None")?;
        Ok(())
    }
}

impl<T: Default> Default for DList<T> {
    fn default() -> Self {
        Self::new()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn list(v: &[i32]) -> DList<i32> {
        DList::from_slice(v)
    }

    fn values(l: &DList<i32>) -> Vec<i32> {
        l.iter().map(|node| node.borrow().val).collect()
    }

    fn pop_val(l: &mut DList<i32>) -> Option<i32> {
        l.pop_front()
    }

    #[test]
    fn empty_list_has_no_nodes() {
        let mut l = DList::<i32>::new();
        assert_eq!(0, l.len());
        assert!(l.is_empty());
        assert_eq!(0, l.iter().count());
        assert_eq!("None", l.to_string());
        assert_eq!(None, pop_val(&mut l));
    }

    #[test]
    fn from_slice_links_both_directions() {
        let l = list(&[1, 2, 3]);
        assert_eq!(3, l.len());
        assert!(!l.is_empty());
        assert_eq!(vec![1, 2, 3], values(&l));
        assert_eq!("1 -> 2 -> 3 -> None", l.to_string());

        // Forward links skip the tail sentinel, backward links reach the head.
        let first = l.iter().next().unwrap();
        assert_eq!(1, first.borrow().val);
        let second = first.borrow().next.clone().unwrap();
        assert_eq!(2, second.borrow().val);
        assert_eq!(
            1,
            second
                .borrow()
                .prev
                .clone()
                .unwrap()
                .upgrade()
                .unwrap()
                .borrow()
                .val
        );
        let last = l.iter().last().unwrap();
        assert_eq!(3, last.borrow().val);
        assert_eq!(
            2,
            last.borrow()
                .prev
                .clone()
                .unwrap()
                .upgrade()
                .unwrap()
                .borrow()
                .val
        );
    }

    #[test]
    fn push_front_prepends() {
        let mut l = DList::<i32>::new();
        l.push_front(1);
        assert_eq!(vec![1], values(&l));
        l.push_front(0);
        assert_eq!(vec![0, 1], values(&l));
        assert_eq!(Some(0), l.iter().next().map(|n| n.borrow().val));
        assert_eq!("0 -> 1 -> None", l.to_string());
    }

    #[test]
    fn push_back_appends() {
        let mut l = DList::<i32>::new();
        l.push_back(1);
        assert_eq!(vec![1], values(&l));
        l.push_back(2);
        assert_eq!(vec![1, 2], values(&l));
        assert_eq!(2, l.len());
        assert_eq!("1 -> 2 -> None", l.to_string());

        // The appended node still points back to its predecessor.
        let last = l.iter().last().unwrap();
        assert_eq!(
            1,
            last.borrow()
                .prev
                .clone()
                .unwrap()
                .upgrade()
                .unwrap()
                .borrow()
                .val
        );
    }

    #[test]
    fn pop_front_removes_nodes_until_empty() {
        let mut l = list(&[1, 2, 3]);
        assert_eq!(Some(1), pop_val(&mut l));
        assert_eq!(vec![2, 3], values(&l));
        assert_eq!(2, l.len());

        assert_eq!(Some(2), pop_val(&mut l));
        assert_eq!(Some(3), pop_val(&mut l));
        assert!(l.is_empty());
        assert!(l.iter().next().is_none());
        assert_eq!("None", l.to_string());

        // Popping an empty list is a no-op.
        assert_eq!(None, pop_val(&mut l));
        assert_eq!(0, l.len());
    }

    #[test]
    fn mixes_front_and_back_operations() {
        let mut l = list(&[2]);
        l.push_front(1);
        l.push_back(3);
        assert_eq!(vec![1, 2, 3], values(&l));
        assert_eq!(Some(1), pop_val(&mut l));
        assert_eq!(vec![2, 3], values(&l));
        l.push_back(4);
        assert_eq!(vec![2, 3, 4], values(&l));
        assert_eq!(3, l.len());
    }

    #[test]
    fn works_with_non_copy_values() {
        let mut l = DList::<String>::new();
        l.push_back(String::from("a"));
        l.push_front(String::from("b"));
        assert_eq!(
            vec!["b", "a"],
            l.iter().map(|n| n.borrow().val.clone()).collect::<Vec<_>>()
        );
        assert_eq!("b -> a -> None", l.to_string());

        let popped = l.pop_front().expect("list is not empty");
        assert_eq!("b", popped);
    }
}
