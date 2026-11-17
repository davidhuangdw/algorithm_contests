pub struct KMP {
    pattern: Vec<char>,
    nxt: Vec<isize>,
}

impl KMP {
    pub fn new(pattern: &str) -> Self {
        let pattern: Vec<char> = pattern.chars().collect();
        let nxt = get_nxt(&pattern);
        Self { pattern, nxt }
    }

    pub fn search(&self, s: &str) -> Option<usize> {
        let n = self.pattern.len();
        if n == 0 {
            return Some(0);
        }
        let mut j = 0;
        for (i, ch) in s.chars().enumerate() {
            while j >= 0 && self.pattern[j as usize] != ch {
                j = self.nxt[j as usize];
            }
            j += 1;
            if n == j as usize {
                return Some(i + 1 - n);
            }
        }
        None
    }
}

fn get_prev_len(p: &[char]) -> Vec<usize> {
    let mut prev_len = vec![0; p.len()];
    for i in 1..p.len() {
        let mut j = prev_len[i - 1];
        while j > 0 && p[j] != p[i] {
            j = prev_len[j - 1];
        }
        prev_len[i] = if p[j] == p[i] { j + 1 } else { 0 };
    }
    prev_len
}

fn get_nxt(p: &[char]) -> Vec<isize> {
    let prev = get_prev_len(p);
    let mut nxt = vec![-1; p.len()];
    for i in 1..p.len() {
        let j = prev[i - 1];
        nxt[i] = if p[j] == p[i] { nxt[j] } else { j as isize };
    }
    nxt
}

#[cfg(test)]
mod tests {
    use crate::string::test_util::strs;

    use super::*;

    /// cross-checked against a naive `windows` scan.
    #[test]
    fn matches_naive_search() {
        let strs = strs(&['a', 'b', 'c'], 6);
        for p in strs.iter().filter(|p| (1..=4).contains(&p.len())) {
            let m = KMP::new(p);
            for s in &strs {
                let want = s
                    .chars()
                    .collect::<Vec<_>>()
                    .windows(p.chars().count())
                    .position(|w| w.iter().copied().eq(p.chars()));
                assert_eq!(want, m.search(s), "pattern {p:?} in {s:?}");
            }
        }
    }

    #[test]
    fn edge_cases() {
        assert_eq!(Some(0), KMP::new("a").search("a"));
        assert_eq!(None, KMP::new("abc").search("ab"));
        assert_eq!(None, KMP::new("ab").search(""));
        assert_eq!(Some(1), KMP::new("aa").search("baaa")); // leftmost match
        // Returns a char index, not a byte offset.
        assert_eq!(Some(2), KMP::new("世界").search("你好世界"));
    }

    #[test]
    fn empty_pattern_matches_at_zero() {
        assert_eq!(Some(0), KMP::new("").search("abc"));
    }
}
