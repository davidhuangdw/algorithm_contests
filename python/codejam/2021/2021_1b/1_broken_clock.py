def input_numbers(t=int):
    return map(t, input().split())


def comb(arr, d):
    if d >= 2:
        yield arr
    else:
        for i in range(d, 3):
            arr[d], arr[i] = arr[i], arr[d]
            for x in comb(arr, d+1):
                yield x
            arr[d], arr[i] = arr[i], arr[d]


NA = 10**9
V = 12 * 60 * 60 * NA
# MAP = {}
# for v in range(V):
#     h = v
#     m = (v * 12) % V
#     s = (v * 12 * 60) % V
#     a, b, c = sorted((h, m, s))
#     MAP[(b - a, c - a)] = v
#     MAP[(c - b, (a + V - b) % V)] = v
#     MAP[((a + V - c) % V, (b + V - c) % V)] = v

T, = input_numbers()
for _c in range(1, T+1):
    A, B, C = sorted(list(input_numbers()))
    # v = MAP[(B-A)//NA, (C-A)//NA]
    v = V
    ok = False
    for a, b, c in comb([A, B, C], 0):
        d = (b - a) % V
        # 11*v = d (mod V), 0<= 11*v < 11*V
        for k in range(12):
            if (d + k*V) % 11 != 0:
                continue
            v = ((d + k*V) // 11) % V
            if (c-a - 719*v) % V == 0:
                ok = True
                break
        if ok:
            break
    assert ok
    time = []
    for bt in (NA, 60, 60, 12):
        time.append(str(v % bt))
        v = v // bt
    print(f"Case #{_c}: {' '.join(reversed(time))}")
