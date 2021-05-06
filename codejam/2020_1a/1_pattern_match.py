def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    N, = input_numbers()
    All = []
    for i in range(N):
        All.append(input().split('*'))

    conflict = False
    L, R = [], []
    for i in range(N):
        l, r = list(All[i][0]), list(reversed(All[i][-1]))
        if any(L[j] != l[j] for j in range(min(len(l), len(L)))) \
                or any(R[j] != r[j] for j in range(min(len(r), len(R)))):
            conflict = True
            break
        if len(L) < len(l):
            L = l
        if len(R) < len(r):
            R = r
    res = ''
    if conflict:
        res = '*'
    else:
        res = L
        for i in range(N):
            for s in All[i][1:-1]:
                res.append(s)
        for x in reversed(R):
            res.append(x)
    print(f"Case #{_c}: {''.join(res)}")
