def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    N, A, B = input_numbers()
    U = list(input_numbers()) + [0]*100
    i = 0
    while i < N-1 or not(U[i+1] == 0 and U[i] == 1):
        d = U[i] - U[i+1]
        if d > 0:
            U[i] -= (d + 1) // 2
            U[i+1] += (d + 1) // 2
        U[i+2] += U[i]
        U[i+1] -= U[i]
        U[i] = 0
        i += 1
    print(f"Case #{_c}: {i+1}")
