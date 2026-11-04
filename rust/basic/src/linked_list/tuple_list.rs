#[derive(Debug, Default)]
pub enum Lst<T> {
    Con(T, Box<Lst<T>>),
    #[default]
    Nil,
}

use std::{fmt::Display, iter, mem::{take,replace}};

use Lst::*;

impl<T> Lst<T> {
    pub fn from_slice(s: &[T]) -> Self
    where
        T: Clone,
    {
        let mut l = Nil;
        for t in s.iter().rev() {
            l = Con(t.clone(), Box::new(l));
        }
        l
    }

    pub fn peek(&self) -> Option<&T> {
        match self {
            Con(v, _) => Some(v),
            Nil => None,
        }
    }

    fn as_con(&self) -> Option<&Self> {
        match self {
            Con(..) => Some(self),
            Nil => None,
        }
    }

    pub fn push_front(&mut self, v: T) {
        let pre = take(self);
        *self = Con(v, Box::new(pre));
    }

    pub fn pop_front(&mut self) -> Option<T> {
        let Con(v, nxt) = replace(self, Nil) else { return None };
        *self = *nxt;
        Some(v)

    }

    pub fn iter(&self) -> impl Iterator<Item = &T> {
        iter::successors(self.as_con(), |x| match x {
            Con(_, nxt) => nxt.as_con(),
            Nil => None,
        })
        .map(|x| match x {
            Con(v, _) => v,
            Nil => unreachable!("as_con only yields Con nodes"),
        })
    }
}

impl<T: Display> Display for Lst<T> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for x in self.iter() {
            write!(f, "{} -> ", x)?;
        }
        write!(f, "None")
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn list(v: &[i32]) -> Lst<i32> {
        Lst::from_slice(v)
    }

    #[test]
    fn empty_list_prints_none() {
        let l = list(&[]);
        assert!(matches!(l, Nil));
        assert_eq!(None, l.peek());
        assert_eq!(0, l.iter().count());
        assert_eq!("None", l.to_string());
    }

    #[test]
    fn from_slice_builds_list_in_order() {
        let l = list(&[1, 2, 3]);
        assert_eq!(Some(&1), l.peek());
        assert_eq!(vec![1, 2, 3], l.iter().copied().collect::<Vec<_>>());
        assert_eq!("1 -> 2 -> 3 -> None", l.to_string());

        assert_eq!(Some(&7), list(&[7]).peek());
        assert_eq!("7 -> None", list(&[7]).to_string());
    }

    #[test]
    fn push_front_prepends() {
        let mut l = list(&[]);
        l.push_front(1);
        assert_eq!("1 -> None", l.to_string());

        l.push_front(0);
        assert_eq!(Some(&0), l.peek());
        assert_eq!(vec![0, 1], l.iter().copied().collect::<Vec<_>>());

        l.push_front(-1);
        assert_eq!("-1 -> 0 -> 1 -> None", l.to_string());
    }

    #[test]
    fn pop_front_removes_head() {
        let mut l = list(&[1, 2, 3]);
        assert_eq!(Some(1), l.pop_front());
        assert_eq!(vec![2, 3], l.iter().copied().collect::<Vec<_>>());

        assert_eq!(Some(2), l.pop_front());
        assert_eq!(Some(3), l.pop_front());
        assert!(matches!(l, Nil));

        // Popping past the end is a no-op, not a panic.
        assert_eq!(None, l.pop_front());
        assert!(matches!(l, Nil));
        assert_eq!("None", l.to_string());
    }

    #[test]
    fn works_with_non_copy_values() {
        let l = Lst::from_slice(&[String::from("a"), String::from("b")]);
        assert_eq!(Some("a"), l.peek().map(String::as_str));
        assert_eq!("a -> b -> None", l.to_string());
    }
}
