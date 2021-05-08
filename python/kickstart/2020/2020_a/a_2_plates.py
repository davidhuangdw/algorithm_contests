def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    N, K, P = input_numbers()
    F = [0] * (P+1)
    for _ in range(N):
        A = list(input_numbers())
        for x in range(P, 0, -1):
            s = 0
            for i in range(min(K, x)):
                s += A[i]
                F[x] = max(F[x], s + F[x-(i+1)])
    print(f"Case #{_c}: {F[P]}")
