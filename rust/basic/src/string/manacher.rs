pub fn max_palindrome(s: &str) -> String {
    if s.len() == 0 {
        return "".to_string();
    }
    let cs: Vec<char> = s.chars().collect();
    let dis = manacher(&cs);
    let i = (0..dis.len()).max_by_key(|&i| dis[i]).unwrap();
    let d = dis[i];
    cs[(i / 2 - d / 2)..=(i / 2 + (d - 1) / 2)].iter().collect()
}

/// return palindrom-len d[] when each s[i] as
///   center-right(d[i*2]) & center(d[i*2+1])
///   d[i] mapping to `s[(i / 2 - d / 2)..=(i / 2 + (d - 1) / 2)]`
pub fn manacher(chars: &[char]) -> Vec<usize> {
    let mut s = vec!['#'; chars.len() * 2 + 1];
    for (i, ch) in chars.iter().enumerate() {
        s[i * 2 + 1] = *ch;
    }
    let mut dis = vec![0; s.len()];

    let (mut md, mut rt) = (0, 0);
    for i in 1..s.len() - 1 {
        let mut r = i + 1;
        if rt > r {
            r = (i + dis[2 * md - i] + 1).min(rt);
        }
        let mut l = 2 * i - r + 1;
        while 0 < l && r < s.len() && s[l - 1] == s[r] {
            l -= 1;
            r += 1;
        }
        dis[i] = r - 1 - i;
        if r > rt {
            md = i;
            rt = r;
        }
    }
    dis[..s.len() - 1].to_vec()
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::string::test_util::strs;

    /// Radii of the virtual `#c0#c1#...#` string, computed naively.
    fn brute_dis(cs: &[char]) -> Vec<usize> {
        let v: Vec<char> = (0..2 * cs.len() + 1)
            .map(|j| if j % 2 == 0 { '#' } else { cs[(j - 1) / 2] })
            .collect();
        (0..2 * cs.len())
            .map(|p| {
                (1..)
                    .take_while(|&t| p >= t && p + t < v.len() && v[p - t] == v[p + t])
                    .count()
            })
            .collect()
    }

    fn brute_longest(cs: &[char]) -> usize {
        (0..cs.len())
            .flat_map(|i| (i..cs.len()).map(move |j| (i, j)))
            .filter(|&(i, j)| cs[i..=j].iter().eq(cs[i..=j].iter().rev()))
            .map(|(i, j)| j - i + 1)
            .max()
            .unwrap_or(0)
    }

    #[test]
    fn dis_matches_brute_force() {
        for s in strs(&['a', 'b', 'c'], 7) {
            let cs: Vec<char> = s.chars().collect();
            assert_eq!(brute_dis(&cs), manacher(&cs), "s = {s:?}");
        }
    }

    #[test]
    fn longest_palindrome_matches_brute_force() {
        for s in strs(&['a', 'b', 'c'], 7).iter().filter(|s| !s.is_empty()) {
            let out = max_palindrome(s);
            let got: Vec<char> = out.chars().collect();
            assert_eq!(
                brute_longest(&s.chars().collect::<Vec<_>>()),
                got.len(),
                "s = {s:?}"
            );
            assert!(
                got.iter().eq(got.iter().rev()),
                "{out:?} is not a palindrome"
            );
            assert!(s.contains(&out), "{out:?} is not a substring of {s:?}");
        }
    }

    #[test]
    fn empty_input_returns_empty() {
        assert_eq!("", max_palindrome(""));
    }

    #[test]
    fn sentinel_chars_in_input() {
        assert_eq!("@@", max_palindrome("@@"));
    }
}
