from collections import Counter
from math import sqrt


def input_numbers(t=int):
    return map(t, input().split())


Sq = [i*i for i in range(int(sqrt(10**7)))]

T, = input_numbers()
for _c in range(1, T+1):
    N, = input_numbers()
    A = list(input_numbers())
    cnt = Counter()
    ans = 0
    pre = 0
    cnt[0] = 1
    mi = 0
    for v in A:
        pre += v
        for d in Sq:
            if pre - d < mi:
                break
            ans += cnt[pre - d]
        mi = min(mi, pre)
        cnt[pre] += 1
    print(f"Case #{_c}: {ans}")
