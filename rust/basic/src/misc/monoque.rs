use std::collections::VecDeque;

pub fn window_min<T: Ord>(a: &[T], win_size: usize) -> Vec<&T> {
    assert!(win_size > 0);
    let mut ans = vec![];
    let mut cand = VecDeque::new();
    for (i, v) in a.iter().enumerate() {
        if let Some(&j) = cand.front()
            && win_size <= i - j
        {
            cand.pop_front();
        }
        while let Some(&j) = cand.back()
            && a[j] > *v
        {
            cand.pop_back();
        }
        ans.push(if let Some(&j) = cand.front() {
            &a[j]
        } else {
            v
        });
        cand.push_back(i);
    }
    ans
}

#[cfg(test)]
mod tests {
    use super::*;

    fn naive(a: &[i32], win: usize) -> Vec<i32> {
        (0..a.len())
            .map(|i| *a[i.saturating_sub(win - 1)..=i].iter().min().unwrap())
            .collect()
    }

    fn got(a: &[i32], win: usize) -> Vec<i32> {
        window_min(a, win).into_iter().copied().collect()
    }

    fn check(a: &[i32], win: usize) {
        assert_eq!(naive(a, win), got(a, win), "a = {a:?}, win = {win}");
    }

    #[test]
    fn matches_model() {
        // Exhaustive over {0,1,2} up to len 4.
        for len in 0..=4usize {
            for mut code in 0..3usize.pow(len as u32) {
                let a: Vec<i32> = (0..len)
                    .map(|_| {
                        let v = (code % 3) as i32;
                        code /= 3;
                        v
                    })
                    .collect();
                for win in 1..=len + 2 {
                    check(&a, win);
                }
            }
        }
        let mut seed = 0x2545_f491_4f6c_dd1du64;
        for _ in 0..100 {
            seed = seed
                .wrapping_mul(6364136223846793005)
                .wrapping_add(1442695040888963407);
            let a: Vec<i32> = (0..(seed % 25) as usize)
                .map(|i| ((seed >> (i % 17)) % 11) as i32 - 5)
                .collect();
            for win in 1..=a.len() + 2 {
                check(&a, win);
            }
        }
    }

    #[test]
    fn results_alias_the_input() {
        let a = [3, 1, 4, 1, 5, 9, 2, 6];
        let out = window_min(&a, 3);
        assert_eq!(a.len(), out.len());
        for (i, r) in out.iter().enumerate() {
            assert!(a.iter().any(|x| std::ptr::eq(x, *r)), "out[{i}] aliases a");
        }
    }

    #[test]
    fn edge_cases() {
        assert!(got(&[], 3).is_empty());
        assert_eq!(vec![4, 2, 8], got(&[4, 2, 8], 1)); // window of one
        assert_eq!(vec![4, 2, 2, 2], got(&[4, 2, 8, 9], 99)); // win > len
        assert_eq!(vec![5, 5, 5], got(&[5, 5, 5], 2)); // equal values
        let s = ["b", "a", "c"]; // non-Copy, ordered by `Ord`
        assert_eq!(
            vec!["b", "a", "a"],
            window_min(&s, 2).iter().map(|s| **s).collect::<Vec<_>>()
        );
    }

    #[test]
    #[should_panic]
    fn rejects_zero_window() {
        window_min(&[1, 2, 3], 0);
    }

    /// Regression: `j + win_size` used to overflow near `usize::MAX`.
    #[test]
    fn huge_window_means_whole_array() {
        assert_eq!(vec![3, 1, 1, 1, 1], got(&[3, 1, 4, 1, 5], usize::MAX));
    }
}
