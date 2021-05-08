def input_numbers(t=int):
    return map(t, input().split())


def nxt(st, y):
    # if st >= y:
    #     raise Exception("not allowed")
    cur = st
    while cur <= y:
        st = st+1
        cur = cur * (10**(len(str(st)))) + st
    return cur


T, = input_numbers()
for _c in range(1, T+1):
    Y, = input_numbers()
    S = str(Y)
    mi = nxt(1, Y)
    for l in range(1, (len(S)+1)//2+1):
        over = 10**l
        for d in range(30):
            if over - d >= 1 and over-d <= Y:
                mi = min(mi, nxt(over-d, Y))
        st = int(S[:l])
        for d in range(30):
            if st+d <= Y:
                mi = min(mi, nxt(st+d, Y))
    print(f"Case #{_c}: {mi}")
