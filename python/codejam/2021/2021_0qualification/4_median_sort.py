def input_numbers(t=int):
    return map(t, input().split())


# T, = input_numbers()
T, N, Q = input_numbers()
for _c in range(1, T+1):
    def solve():
        cur = [i+1 for i in range(N)]
        for i in range(3, N+1):
            l, r = 0, i-2
            while l <= r:
                d = (r - l + 1) // 3
                a, b = l+d, r-d
                if a == b:
                    # a -= 1            # ----bug!!!
                    if a > 0:
                        a -= 1
                    else:
                        b += 1
                print(cur[a], cur[b], i)
                md, = input_numbers()
                if md == cur[a]:
                    r = a-1
                elif md == cur[b]:
                    l = b+1
                elif md == i:
                    l = a+1
                    r = b-1
                else:
                    raise RuntimeError("exceed limit")
            pos = l
            for j in range(i-1, pos, -1):
                cur[j] = cur[j-1]
            cur[pos] = i
            # print(cur[:i])
        return ' '.join(map(str, cur))
    print(solve())
    if '1' != input():
        raise RuntimeError("failed")

    # print(f"Case #{_c}: ")
