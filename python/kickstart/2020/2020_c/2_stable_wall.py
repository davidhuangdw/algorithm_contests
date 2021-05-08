from collections import defaultdict


def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    R, C = input_numbers()
    Adj = defaultdict(set)
    V = set()
    pre = None
    for _ in range(R):
        cur = list(input())
        V |= set(cur)
        if pre:
            for fr, to in zip(cur, pre):
                if fr != to:
                    Adj[fr].add(to)
        pre = cur
    ans = []
    loop = False
    done = set()

    def has_loop(a):
        vis.add(a)
        done.add(a)
        for b in Adj[a]:
            if b in vis:
                return True
            if b in done:
                continue
            if has_loop(b):
                return True
        ans.append(a)
        vis.remove(a)           # bug!!: remove
        return False

    for v in V:
        if v not in done:
            vis = set()
            if has_loop(v):
                loop = True
                break
            # done |= vis           # bug!!: shared-variable: should be mutated instead of assigned
    print(f"Case #{_c}: {-1 if loop else ''.join(reversed(ans))}")
