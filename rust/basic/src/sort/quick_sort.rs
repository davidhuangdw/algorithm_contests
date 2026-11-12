use rand::random_range;

fn partition<T: Ord>(a: &mut [T]) -> usize {
    let n = a.len();
    a.swap(0, random_range(..n));

    let (mut l, mut r) = (1, n - 1);
    while l <= r {
        while l <= r && a[l] <= a[0] {
            l += 1;
        }
        while l <= r && a[r] > a[0] {
            r -= 1;
        }
        if l < r {
            a.swap(l, r);
            l += 1;
            r -= 1;
        }
    }
    a.swap(0, r);
    r
}

// (short-half, longer-half)
fn split<T: Ord>(a: &mut [T], i: usize) -> (&mut [T], &mut [T]) {
    let (a, b) = a.split_at_mut(i);
    let b = &mut b[1..];
    if a.len() < b.len() { (a, b) } else { (b, a) }
}

pub fn quick_sort<T: Ord>(a: &mut [T]) {
    if a.len() <= 1 {
        return;
    }
    let i = partition(a);

    let (a, b) = split(a, i);
    quick_sort(a); // short half first for less recur-depth
    quick_sort(b);
}

#[cfg(test)]
mod tests {
    use super::*;

    /// Sorts `input` and checks it against `sort_unstable`.
    fn assert_sorts(input: &[i32]) {
        let mut got = input.to_vec();
        quick_sort(&mut got);

        let mut want = input.to_vec();
        want.sort_unstable();

        assert_eq!(want, got, "input {input:?}");
    }

    #[test]
    fn sorts_edge_cases() {
        assert_sorts(&[]);
        assert_sorts(&[7]);
        assert_sorts(&[2, 1]);
        assert_sorts(&[1, 2, 3]);
        assert_sorts(&[3, 2, 1]);
        assert_sorts(&[3, 1, 2]);
        assert_sorts(&[4, 4, 4, 4]);
        assert_sorts(&[3, -1, 0, 3, -5, 0]);
    }

    #[test]
    fn sorts_random_input() {
        let mut seed = 12345u64;
        for _ in 0..50 {
            seed = seed.wrapping_mul(6364136223846793005).wrapping_add(1);
            let n = (seed % 40) as usize;
            let input: Vec<i32> = (0..n)
                .map(|i| ((i as u64 * 7919 + seed) % 17) as i32)
                .collect();
            assert_sorts(&input);
        }
    }

    #[test]
    fn sorts_non_copy_values() {
        let mut words: Vec<String> = ["c", "a", "b"].iter().map(|s| s.to_string()).collect();
        quick_sort(&mut words);
        assert_eq!(vec!["a", "b", "c"], words);
    }
}
