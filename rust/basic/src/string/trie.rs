use std::collections::HashMap;

#[derive(Debug)]
pub struct Trie {
    root: Node,
}

impl Trie {
    pub fn new() -> Self {
        Self { root: Node::new() }
    }

    pub fn count(&self, s: &str) -> usize {
        let mut u = &self.root;
        for ch in s.chars() {
            if let Some(p) = u.nxt.get(&ch) {
                u = p;
            } else {
                return 0;
            }
        }
        u.cnt
    }

    fn get_node(&mut self, s: &str) -> &mut Node {
        let mut u = &mut self.root;
        for ch in s.chars() {
            u = u.nxt.entry(ch).or_insert_with(Node::new);
        }
        u
    }

    pub fn remove(&mut self, s: &str) -> bool {
        if self.count(s) > 0 {
            self.get_node(s).cnt -= 1;
            true
        } else {
            false
        }
    }

    pub fn add(&mut self, s: &str) {
        self.get_node(s).cnt += 1;
    }
}

impl Default for Trie {
    fn default() -> Self {
        Self { root: Node::new() }
    }
}

#[derive(Debug)]
struct Node {
    nxt: HashMap<char, Node>,
    cnt: usize,
}

impl Node {
    fn new() -> Node {
        Self {
            nxt: HashMap::new(),
            cnt: 0,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn empty_trie_counts_nothing() {
        let mut t = Trie::new();
        assert_eq!(0, t.count(""));
        assert_eq!(0, t.count("abc"));
        assert!(!t.remove("abc"));
    }

    #[test]
    fn add_and_remove_counts_per_word() {
        let mut t = Trie::new();
        t.add("go");
        t.add("go");
        t.add("golang");

        // Exact matches only: prefixes and extensions are different keys.
        assert_eq!(2, t.count("go"));
        assert_eq!(1, t.count("golang"));
        assert_eq!(0, t.count("g"));
        assert_eq!(0, t.count("gol"));

        // Removing all copies of `go` leaves `golang` intact.
        assert!(t.remove("go"));
        assert_eq!(1, t.count("go"));
        assert!(t.remove("go"));
        assert_eq!(0, t.count("go"));
        assert_eq!(1, t.count("golang"));

        // Absent words are reported, not decremented.
        assert!(!t.remove("go"));
        assert!(!t.remove("rust"));
        assert_eq!(1, t.count("golang"));
    }

    #[test]
    fn prefix_word_is_counted_on_its_own() {
        let mut t = Trie::new();
        t.add("abc");
        t.add("ab");
        assert_eq!(1, t.count("ab"));
        assert_eq!(1, t.count("abc"));
    }

    #[test]
    fn empty_string_is_a_key() {
        let mut t = Trie::new();
        t.add("");
        assert_eq!(1, t.count(""));
        assert!(t.remove(""));
        assert_eq!(0, t.count(""));
    }

    #[test]
    fn splits_on_chars_not_bytes() {
        let mut t = Trie::new();
        t.add("中文");
        t.add("中");
        assert_eq!(1, t.count("中文"));
        assert_eq!(1, t.count("中"));
        assert_eq!(0, t.count("中国"));
        assert!(t.remove("中文"));
        assert_eq!(0, t.count("中文"));
        assert_eq!(1, t.count("中"));
    }
}
