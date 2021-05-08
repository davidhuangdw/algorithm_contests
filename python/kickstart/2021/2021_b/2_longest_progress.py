def input_numbers(t=int):
    return map(t, input().split())


INF = 10**10
T, = input_numbers()
for _c in range(1, T+1):
    N, = input_numbers()
    A = list(input_numbers())
    A.append(INF)

    # bug cases:
    # 5: 5 5 3 2 1
    # 5: 9 1 3 5 7
    ans = 2
    R = [(0, 0)] * (N+1)
    R[N] = (0, INF)
    R[N-1] = (1, INF)
    for i in range(N-2, -1, -1):
        ans = max(ans, R[i+1][0]+1)
        nl, nd = R[i+1]
        l, d = 2, A[i+1]-A[i]
        if d == nd:
            l = nl + 1
        R[i] = (l, d)

    pl, pd = 1, INF
    for i in range(1, N):
        nl, nd = R[i+1]
        ans = max(ans, pl+1+(A[i-1] + pd*2 == A[i+1]), nl+1+(A[i-1]+nd*2 == A[i+1]))
        if pd == nd and A[i-1] + 2*pd == A[i+1]:
            ans = max(ans, pl + nl + 1)
        pl, pd = 1 + (pl if pd == A[i] - A[i-1] else 1), A[i] - A[i-1]
    print(f"Case #{_c}: {ans}")
