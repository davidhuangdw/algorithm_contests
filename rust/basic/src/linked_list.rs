pub mod array_double;
pub mod double_link;
pub mod single_link;
pub mod tuple_list;

pub use array_double::AList;
pub use double_link::DList;
pub use single_link::List;
pub use tuple_list::Lst;

pub trait Seq<T> {
    fn new() -> Self
    where
        T: Default,
        Self: Sized;

    fn len(&self) -> usize;
    fn push_front(&mut self, v: T);
    fn pop_front(&mut self) -> Option<T>;
    fn push_back(&mut self, v: T);
    fn pop_back(&mut self) -> Option<T>;
    fn iter(&self) -> Box<dyn Iterator<Item = &T> + '_>;

    fn is_empty(&self) -> bool {
        self.len() == 0
    }

    fn peek(&self) -> Option<&T> {
        self.iter().next()
    }

    fn get(&self, k: usize) -> Option<&T> {
        self.iter().nth(k)
    }

    fn from_slice(s: &[T]) -> Self
    where
        T: Clone + Default,
        Self: Sized,
    {
        let mut out = Self::new();
        for v in s.iter().rev() {
            out.push_front(v.clone());
        }
        out
    }

    fn to_vec(&self) -> Vec<T>
    where
        T: Clone,
    {
        self.iter().cloned().collect()
    }
}

#[cfg(test)]
pub(crate) mod test_util {
    use super::Seq;

    /// Cross-checks a sequence against `model` (front first) through every
    /// read API the trait exposes.
    pub fn check<Q: Seq<i32>>(q: &Q, model: &[i32]) {
        assert_eq!(model.len(), q.len(), "len");
        assert_eq!(model.is_empty(), q.is_empty(), "is_empty");
        assert_eq!(model.first(), q.peek(), "peek");
        assert_eq!(model, q.iter().copied().collect::<Vec<_>>(), "iter");
        assert_eq!(model, q.to_vec(), "to_vec");
        for (k, want) in model.iter().enumerate() {
            assert_eq!(Some(want), q.get(k), "get({k})");
        }
        assert_eq!(None, q.get(model.len()), "get(len)");
    }

    /// Exercises `new` / `from_slice` / `push_back` / `pop_front` / `pop_back`
    /// of a `Seq` implementation and checks each intermediate state.
    pub fn check_built<Q: Seq<i32>>(model: &[i32]) {
        check(&Q::new(), &[]);

        let built = Q::from_slice(model);
        check(&built, model);

        let mut drained = Q::from_slice(model);
        let mut got = Vec::new();
        while let Some(v) = drained.pop_front() {
            got.push(v);
        }
        assert_eq!(model, got.as_slice(), "pop_front drain");
        assert!(drained.is_empty());

        let mut back = Q::new();
        for &v in model {
            back.push_back(v);
        }
        check(&back, model);
        let mut got = Vec::new();
        while let Some(v) = back.pop_back() {
            got.push(v);
        }
        got.reverse();
        assert_eq!(model, got.as_slice(), "pop_back drain");
        assert!(back.is_empty());
    }
}

#[cfg(test)]
mod tests {
    use super::test_util::check_built;
    use super::{AList, List, Lst, Seq};

    /// The whole point of the trait: one contract test drives every impl.
    #[test]
    fn every_impl_satisfies_the_seq_contract() {
        for model in [&[][..], &[0][..], &[1, 2, 3][..], &[5, 5, -1, 0, 7][..]] {
            check_built::<List<i32>>(model);
            check_built::<Lst<i32>>(model);
            check_built::<AList<i32>>(model);
        }
    }

    #[test]
    fn trait_from_slice_is_reachable_generically() {
        fn build<Q: Seq<i32>>(s: &[i32]) -> Vec<i32> {
            Q::from_slice(s).to_vec()
        }

        let s = [4, 2, 9];
        assert_eq!(s, build::<List<i32>>(&s).as_slice());
        assert_eq!(s, build::<Lst<i32>>(&s).as_slice());
        assert_eq!(s, build::<AList<i32>>(&s).as_slice());
    }
}
