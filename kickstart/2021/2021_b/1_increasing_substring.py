def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    N, = input_numbers()
    S = input()
    S = 'Z' + S
    ans = [-1]
    for i in range(1, N+1):
        ans.append(1 + (ans[i-1] if S[i] > S[i-1] else 0))
    print(f"Case #{_c}: {' '.join(map(str, ans[1:]))}")
