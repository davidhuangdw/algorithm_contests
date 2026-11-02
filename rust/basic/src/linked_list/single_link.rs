use std::fmt::Display;

type Link = Option<Box<ListNode>>;

#[derive(Debug, PartialEq, Eq)]
pub struct ListNode {
    pub val: i32,
    pub next: Link,
}

pub struct List {
    pub head: Link,
}

impl List {
    pub fn from_slice(vec: &[i32]) -> Self {
        let mut head = None;
        for v in vec.iter().rev() {
            head = Some(Box::new(ListNode {
                val: *v,
                next: head,
            }));
        }
        Self { head }
    }
    pub fn to_vec(&self) -> Vec<i32> {
        self.iter().map(|node| node.val).collect()
    }
    pub fn iter(&self) -> impl Iterator<Item = &ListNode> {
        std::iter::successors(self.head.as_deref(), |node| node.next.as_deref())
    }
}

impl Display for List {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut p = &self.head;
        while let Some(node) = p {
            write!(f, "{} -> ", node.val)?;
            p = &node.next;
        }
        write!(f, "None")?;
        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn from_vec_handles_empty_input() {
        let list = List::from_slice(&[]);
        assert!(list.head.is_none());
        assert!(list.iter().next().is_none());
        assert_eq!(Vec::<i32>::new(), list.to_vec());
    }

    #[test]
    fn from_vec_builds_list_in_order() {
        let v = vec![-2, 3, 3, 5, 0];
        let list = List::from_slice(&v);
        assert_eq!(v, list.to_vec());
        assert_eq!(-2, list.head.as_ref().unwrap().val);
        // Every suffix of the input is a valid sub-list.
        assert_eq!(
            vec![3, 3, 5, 0],
            list.iter().skip(1).map(|n| n.val).collect::<Vec<_>>()
        );

        let single = List::from_slice(&[7]);
        assert_eq!(vec![7], single.to_vec());
        assert!(single.head.as_ref().unwrap().next.is_none());
    }

    #[test]
    fn iter_yields_nodes_in_order() {
        let list = List::from_slice(&[1, 2, 3]);
        let vals: Vec<i32> = list.iter().map(|node| node.val).collect();
        assert_eq!(vec![1, 2, 3], vals);
    }

    #[test]
    fn iter_ends_after_last_node() {
        let single = List::from_slice(&[9]);
        let mut it = single.iter();
        assert_eq!(Some(9), it.next().map(|node| node.val));
        assert!(it.next().is_none());
    }

    #[test]
    fn display_joins_values_with_arrow() {
        assert_eq!(
            "1 -> 2 -> 3 -> None",
            List::from_slice(&[1, 2, 3]).to_string()
        );
        assert_eq!("9 -> None", List::from_slice(&[9]).to_string());
        assert_eq!("None", List::from_slice(&[]).to_string());
    }
}
