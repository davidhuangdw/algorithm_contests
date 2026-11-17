pub mod kmp;
pub mod manacher;
pub mod rotate_min;
pub mod trie;

#[cfg(test)]
pub(crate) mod test_util {
    /// All strings over `alpha` up to `max` chars.
    pub fn strs(alpha: &[char], max: usize) -> Vec<String> {
        (0..=max)
            .flat_map(move |n| {
                (0..alpha.len().pow(n as u32)).map(move |mut code| {
                    (0..n)
                        .map(|_| {
                            let c = alpha[code % alpha.len()];
                            code /= alpha.len();
                            c
                        })
                        .collect::<String>()
                })
            })
            .collect()
    }
}
