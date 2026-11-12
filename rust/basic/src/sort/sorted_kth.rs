// O(log(k))
pub fn two_sorted_kth<'t, T: Ord>(a: &'t [T], b: &'t [T], k: usize) -> &'t T {
    assert!(k < a.len() + b.len());
    let (mut a, mut b, mut k) = (a, b, k);

    while k > 0 {
        let la = k.div_ceil(2); // (k+1)/2
        let lb = k + 1 - la;
        if lb > b.len() || (la <= a.len() && a[la - 1] < b[lb - 1]) {
            k -= la;
            a = &a[la..];
        } else {
            k -= lb;
            b = &b[lb..];
        }
    }
    if b.is_empty() || (!a.is_empty() && a[0] <= b[0]) {
        &a[0]
    } else {
        &b[0]
    }
}

// O(log(i64_max)*m*log(n))
pub fn multi_sorted_kth(vs: &[Vec<i64>], k: usize) -> i64 {
    assert!(k < vs.iter().map(|v| v.len()).sum());

    let count_lt = |v: i64| {
        let mut cnt: usize = 0;
        for vec in vs {
            cnt += vec.partition_point(|x| *x < v);
        }
        cnt
    };
    let (mut l, mut r) = (i64::MIN + 1, i64::MAX);
    // binary-search: count(<me)<=k | count(<me)>k
    while l <= r {
        let md = l.midpoint(r);
        if count_lt(md) <= k {
            let Some(nxt) = md.checked_add(1) else { break };
            l = nxt;
        } else {
            r = md - 1;
        }
    }
    r
}

#[cfg(test)]
mod tests {
    use super::*;

    /// Sorted merge of every input slice.
    fn merged(vs: &[&[i64]]) -> Vec<i64> {
        let mut all: Vec<i64> = vs.iter().flat_map(|v| v.iter().copied()).collect();
        all.sort_unstable();
        all
    }

    /// `two_sorted_kth` must equal the merged array at every index.
    fn check_two(a: &[i64], b: &[i64]) {
        for (k, want) in merged(&[a, b]).iter().enumerate() {
            assert_eq!(want, two_sorted_kth(a, b, k), "a={a:?} b={b:?} k={k}");
        }
    }

    /// `multi_sorted_kth` must equal the merged array at every index.
    fn check_multi(vs: &[Vec<i64>]) {
        let refs: Vec<&[i64]> = vs.iter().map(Vec::as_slice).collect();
        for (k, want) in merged(&refs).iter().enumerate() {
            assert_eq!(*want, multi_sorted_kth(vs, k), "vs={vs:?} k={k}");
        }
    }

    #[test]
    fn two_sorted_kth_matches_merged_array() {
        check_two(&[1, 4, 9], &[2, 3, 10, 11]);
        check_two(&[1, 1, 1], &[1, 2, 2]);
        check_two(&[], &[5, 6]);
        check_two(&[5, 6], &[]);
        check_two(&[0], &[0]);
        check_two(&[-4, 8], &[1, 3, 5]);
    }

    #[test]
    #[should_panic]
    fn two_sorted_kth_rejects_k_out_of_range() {
        two_sorted_kth(&[1], &[2], 2);
    }

    #[test]
    fn multi_sorted_kth_matches_merged_array() {
        check_multi(&[vec![1, 4, 9], vec![2, 3, 10], vec![-1, 7]]);
        check_multi(&[vec![1, 1, 2], vec![1, 2, 2]]);
        check_multi(&[vec![5]]);
        check_multi(&[vec![], vec![5, 6]]);
        check_multi(&[vec![], vec![], vec![3]]);
    }

    #[test]
    #[should_panic]
    fn multi_sorted_kth_rejects_k_out_of_range() {
        multi_sorted_kth(&[vec![1], vec![2]], 2);
    }

    #[test]
    fn multi_sorted_kth_handles_i64_bounds() {
        check_multi(&[vec![i64::MIN]]);
        check_multi(&[vec![i64::MIN, i64::MIN + 5]]);
        check_multi(&[vec![i64::MAX]]);
        check_multi(&[vec![i64::MIN], vec![i64::MAX]]);
    }
}
