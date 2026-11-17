pub fn rotate_min(a: &str) -> String {
    let cs: Vec<char> = a.chars().collect();
    let i = rotate_min_pos(&cs);
    cs[i..].iter().chain(&cs[..i]).collect()
}

pub fn rotate_min_pos(a: &[char]) -> usize {
    let (mut i, mut j, mut k) = (0, 1, 0);
    let n = a.len();
    while i < n && j < n && k < n {
        let vi = a[(i + k) % n];
        let vj = a[(j + k) % n];
        if vi == vj {
            k += 1;
            continue;
        }
        if vi < vj {
            // i win
            j += k + 1;
        } else {
            i += k + 1;
        }
        k = 0;
        if i == j {
            j += 1
        }
    }
    i.min(j)
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::string::test_util::strs;

    fn brute_min(s: &str) -> String {
        let cs: Vec<char> = s.chars().collect();
        (0..cs.len())
            .map(|i| cs[i..].iter().chain(&cs[..i]).collect::<String>())
            .min()
            .unwrap_or_default()
    }

    #[test]
    fn matches_brute_force() {
        for s in strs(&['a', 'b', 'c'], 7) {
            assert_eq!(brute_min(&s), rotate_min(&s), "s = {s:?}");

            let cs: Vec<char> = s.chars().collect();
            let i = rotate_min_pos(&cs);
            let rot: String = cs[i..].iter().chain(&cs[..i]).collect();
            assert_eq!(brute_min(&s), rot, "pos {i} of {s:?}");
        }
    }

    #[test]
    fn edge_cases() {
        assert_eq!("", rotate_min(""));
        assert_eq!("a", rotate_min("a"));
        assert_eq!("aaaa", rotate_min("aaaa"));
        assert_eq!("ab", rotate_min("ba"));
        assert_eq!("abb", rotate_min("bba"));
        assert_eq!("acb", rotate_min("cba"));
        assert_eq!("你好", rotate_min("好你"));
    }
}
