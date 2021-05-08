from collections import deque


def input_numbers(t=int):
    return map(t, input().split())


LIMIT = 10 ** 5


def b_to_i(x):
    res = 0
    for c in x:
        res <<= 1
        if c == '1':
            res += 1
    return res


def bsize(x):
    l = 0
    while x > 0:
        l += 1
        x >>= 1
    return max(l, 1)


T, = input_numbers()
for _c in range(1, T+1):
    S, E = list(map(b_to_i, input().split()))
    done = {S: 0}
    if E == S:
        print(f"Case #{_c}: {0}")
        continue
    que = deque([(S, 0)])
    found = False
    while que and len(done) < LIMIT and not found:
        x, d = que.popleft()
        l = bsize(x)
        for y in ((1 << l) - 1 - x, x << 1):
            if y not in done:
                done[y] = d+1
                que.append((y, d+1))
                if y == E:
                    found = True
    print(f"Case #{_c}: {done.get(E, 'IMPOSSIBLE')}")
