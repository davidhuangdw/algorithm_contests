def input_numbers(t=int):
    return map(t, input().split())


def rev(a, l, r):
    while l < r:
        a[l], a[r] = a[r], a[l]
        l += 1
        r -= 1


def compute(a):
    a = [x for x in a]
    cnt = 0
    for i in range(N - 1):
        mi = i
        for j in range(i + 1, N):
            if a[j] < a[mi]:
                mi = j
        cnt += mi - i + 1
        rev(a, i, mi)
    return cnt


T, = input_numbers()
for _c in range(1, T+1):
    N, C = input_numbers()

    def solve():
        if not N <= C + 1 <= N*(N+1)/2:
            return "IMPOSSIBLE"

        costs = [0] * (N-1)
        rem = C - (N-1)
        for i in range(N-1):
            t = min(rem, N-1-i)
            rem -= t
            costs[i] += t
        res = [i+1 for i in range(N)]
        for i in range(N-2, -1, -1):
            rev(res, i, i+costs[i])
        return ' '.join(map(str, res))
    print(f"Case #{_c}: {solve()}")
