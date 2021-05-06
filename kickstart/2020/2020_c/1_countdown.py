def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    N, K = input_numbers()
    pre = -1
    cnt = 0
    for v in input_numbers():
        pre = v if v == K or v == pre - 1 else -1
        if pre == 1:
            cnt += 1
    print(f"Case #{_c}: {cnt}")
