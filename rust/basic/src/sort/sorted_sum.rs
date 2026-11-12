type ValCnt = (i64, usize);

// return (max(a+b<=target), count) -- O(n)
pub fn two_sorted_sum(a: &[i64], b: &[i64], target: i64) -> (i64, usize) {
    let (a, b) = (uniq_cnt_sorted(a), uniq_cnt_sorted(b));
    let mut max_cnt = (i64::MIN, 0);
    let mut j = b.len();
    for va in a {
        while j > 0 && b[j - 1].0 + va.0 > target {
            j -= 1;
        }
        if j > 0 {
            max_cnt = max_vcnt(max_cnt, add_vcnt(va, b[j - 1]));
        }
    }
    max_cnt
}

// return (max(sum<=target), count) -- O(n^2)
pub fn three_sorted_sum(a: &[i64], b: &[i64], c: &[i64], target: i64) -> (i64, usize) {
    let mut max_cnt = (i64::MIN, 0);
    for va in uniq_cnt_sorted(a) {
        let bc = two_sorted_sum(b, c, target - va.0);
        max_cnt = max_vcnt(max_cnt, add_vcnt(va, bc));
    }
    max_cnt
}

// no need sorted, return (max(sum<=target), count) -- O(n^2*log)
pub fn four_sum(a: &[i64], b: &[i64], c: &[i64], d: &[i64], target: i64) -> (i64, usize) {
    let ab = sorted_sum(a, b);
    let cd = sorted_sum(c, d);
    two_sorted_sum(&ab, &cd, target)
}

fn sorted_sum(a: &[i64], b: &[i64]) -> Vec<i64> {
    let mut sums: Vec<i64> = a
        .iter()
        .flat_map(|va| b.iter().map(|vb| *va + *vb))
        .collect();
    sums.sort();
    sums
}

fn uniq_cnt_sorted(a: &[i64]) -> Vec<ValCnt> {
    let mut uniq: Vec<i64> = vec![];
    let mut cnt: Vec<usize> = vec![];
    for &v in a {
        if let Some(&last) = uniq.last()
            && last == v
        {
            *cnt.last_mut().unwrap() += 1;
        } else {
            uniq.push(v);
            cnt.push(1);
        }
    }
    uniq.into_iter().zip(cnt).collect()
}

fn max_vcnt(a: ValCnt, b: ValCnt) -> ValCnt {
    if a.0 == b.0 {
        (a.0, a.1 + b.1)
    } else if a.0 > b.0 {
        a
    } else {
        b
    }
}

fn add_vcnt(a: ValCnt, b: ValCnt) -> ValCnt {
    if b.1 == 0 {
        return b;
    }
    (a.0 + b.0, a.1 * b.1)
}

#[cfg(test)]
mod tests {
    use super::*;

    /// Brute force over the cartesian product of `vals`; `None` when nothing fits.
    fn brute(vals: &[&[i64]], target: i64) -> Option<(i64, usize)> {
        fn rec(vals: &[&[i64]], i: usize, sum: i64, model: &mut Model) {
            if i == vals.len() {
                model.offer(sum);
                return;
            }
            for &v in vals[i] {
                rec(vals, i + 1, sum + v, model);
            }
        }

        struct Model {
            target: i64,
            best: Option<(i64, usize)>,
        }

        impl Model {
            fn offer(&mut self, sum: i64) {
                if sum > self.target {
                    return;
                }
                match &mut self.best {
                    None => self.best = Some((sum, 1)),
                    Some((best, cnt)) if *best == sum => *cnt += 1,
                    Some((best, _)) if *best < sum => self.best = Some((sum, 1)),
                    Some(_) => {}
                }
            }
        }

        let mut model = Model { target, best: None };
        rec(vals, 0, 0, &mut model);
        model.best
    }

    /// Compares `got` with the model. A "no solution" case only has to report
    /// `count == 0`, the sentinel value itself does not matter.
    fn check(vals: &[&[i64]], target: i64, got: (i64, usize)) {
        match brute(vals, target) {
            Some(want) => assert_eq!(want, got, "{vals:?} t={target}"),
            None => assert_eq!(0, got.1, "{vals:?} t={target} must report no solution"),
        }
    }

    const TARGETS: [i64; 11] = [-3, -1, 0, 1, 2, 3, 4, 5, 6, 8, 12];

    #[test]
    fn two_sorted_sum_matches_brute_force() {
        let cases: [(&[i64], &[i64]); 3] = [
            (&[1, 2], &[3, 4]),
            (&[1, 1, 2], &[2, 2, 3]),
            (&[-4, -1, 5], &[0, 2]),
        ];
        for (a, b) in cases {
            for t in TARGETS {
                check(&[a, b], t, two_sorted_sum(a, b, t));
            }
        }
    }

    #[test]
    fn three_sorted_sum_matches_brute_force() {
        let cases: [(&[i64], &[i64], &[i64]); 2] =
            [(&[1, 2], &[3], &[4, 5]), (&[0, 0], &[1, 2], &[2])];
        for (a, b, c) in cases {
            for t in TARGETS {
                check(&[a, b, c], t, three_sorted_sum(a, b, c, t));
            }
        }
    }

    #[test]
    fn four_sum_matches_brute_force_and_needs_no_sorting() {
        for (a, b, c, d) in [
            (&[1][..], &[2][..], &[3][..], &[4][..]),
            (&[3, 1][..], &[2][..], &[4][..], &[0][..]),
            (&[1, 1][..], &[2][..], &[2][..], &[1][..]),
        ] {
            for t in TARGETS {
                check(&[a, b, c, d], t, four_sum(a, b, c, d, t));
            }
        }
    }

    #[test]
    fn unreachable_target_reports_no_solution() {
        check(&[&[5, 6], &[7]], 3, two_sorted_sum(&[5, 6], &[7], 3));
        check(
            &[&[5], &[6], &[7]],
            3,
            three_sorted_sum(&[5], &[6], &[7], 3),
        );
        // Negative values must also report "no solution" instead of crashing.
        check(
            &[&[-5], &[-5], &[-5]],
            -100,
            three_sorted_sum(&[-5], &[-5], &[-5], -100),
        );
        check(
            &[&[5], &[6], &[7], &[8]],
            3,
            four_sum(&[5], &[6], &[7], &[8], 3),
        );
    }
}
