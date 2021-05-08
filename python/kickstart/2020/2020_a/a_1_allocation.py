def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    N, B = input_numbers()
    A = list(input_numbers())
    A.sort()
    cnt = 0
    for v in A:
        if v <= B:
            B -= v
            cnt += 1
        else:
            break
    print(f"Case #{_c+1}: {cnt}")
