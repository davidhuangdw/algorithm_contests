def input_numbers(t=int):
    return map(t, input().split())


MAX = 1000*1000
T, = input_numbers()
for _c in range(1, T+1):
    C = {(0, 0): 0, (1, 1): 0}
    X, Y, S = input().split()
    C[(0, 1)] = int(X)
    C[(1, 0)] = int(Y)
    # pre = [0, 0]    # bug!!: there's no C, J with zero pre-string
    pre = [0, 0] if S[0] == '?' else ([0, MAX] if S[0] == 'C' else [MAX, 0])
    for ch in S[1:]:
        cur = [MAX, MAX]
        if ch in ('?', 'C'):
            cur[0] = min(cur[0], pre[0], pre[1] + C[(1, 0)])
        if ch in ('?', 'J'):
            cur[1] = min(cur[1], pre[1], pre[0] + C[(0, 1)])
        pre = cur
    print(f"Case #{_c}: {min(pre)}")
