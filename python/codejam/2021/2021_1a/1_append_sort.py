def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    N, = input_numbers()
    A = list(input_numbers())
    cnt = 0
    pre = A[0]
    for x in A[1:]:
        pre += 1
        if x >= pre:
            pre = x
        else:
            sx = str(x)
            sp = str(pre)
            k = len(sx)
            if sx == sp[:k]:
                cnt += len(sp) - k
                continue
            elif sx > sp[:k]:
                add = len(sp) - k
            else:
                add = len(sp) - k + 1
            cnt += add
            pre = int(sx + '0'*add)
    print(f"Case #{_c}: {cnt}")
