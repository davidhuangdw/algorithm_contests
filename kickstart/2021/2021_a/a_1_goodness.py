def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    N, K = input_numbers()
    S = input()
    diff = 0
    for i in range(N//2):
        if S[i] != S[N-1-i]:
            diff += 1
    print(f"Case #{_c}: {abs(K-diff)}")
