def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    N, K = input_numbers()
    P = set(input_numbers())
    P = list(sorted(P))
    one, two = [P[0]-1, P[0]-1]
    for i in range(1, len(P)):
        L = P[i] - P[i-1] - 1
        x = 1 + (L-1)//2
        one, two = max(x, one), max(two, one+x, L)
    last = K - P[-1]
    one, two = max(one, last), max(two, one+last, last)
    print(f"Case #{_c}: {max(one, two)/K}")
