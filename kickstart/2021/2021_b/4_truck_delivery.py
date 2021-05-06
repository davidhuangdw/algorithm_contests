from collections import defaultdict


def input_numbers(t=int):
    return map(t, input().split())


def gcd(a, b):
    while b != 0:
        a, b = b, a % b
    return a


class Node:
    def __init__(self, l, r, g=0, lc=None, rc=None):
        assert(l != r)
        self.l, self.r, self.g, self.lc, self.rc = l, r, g, lc, rc

    @staticmethod
    def _calc(lv, rv):
        return gcd(lv, rv)

    def make_children(self):
        l, r = self.l, self.r
        md = (l + r) // 2
        if l == md:
            raise Exception("l == md")

    def update(self, x, v):
        l, r = self.l, self.r
        if not l <= x < r:
            return
        if r - l == 1:
            self.g = v
        else:
            md = (l + r)//2
            if self.lc is None:
                self.lc, self.rc = Node(l, md), Node(md, r)
            ch = self.lc if x < md else self.rc
            ch.update(x, v)
            self.g = self._calc(self.lc.g, self.rc.g)
        return self.g

    def query(self, x, y):  # [x, y]
        l, r = self.l, self.r
        if y < l or r <= x:
            return 0
        if self.lc is None or (x <= l and r-1 <= y):        # bug: self.lc could be None
            return self.g
        return self._calc(self.lc.query(x, y), self.rc.query(x, y))


T, = input_numbers()
for _c in range(1, T+1):
    N, Q = input_numbers()
    Adj = defaultdict(list)
    E = {}
    Queries = defaultdict(list)
    ml = 0
    for _ in range(N-1):
        x, y, l, a = input_numbers()
        Adj[x].append(y)
        Adj[y].append(x)
        E[(x, y)] = E[(y, x)] = (l, a)
        ml = max(ml, l)
    for i in range(Q):
        c, w = input_numbers()
        Queries[c].append((w, i))
    Root = Node(1, ml+1)
    Ans = [0] * Q

    # def dfs(parent, v):
    #     load, toll = E[(parent, v)]
    #     Root.update(load, toll)
    #     for w, i in Queries[v]:
    #         Ans[i] = Root.query(1, w)
    #     for ch in Adj[v]:
    #         if ch != parent:
    #             dfs(v, ch)
    #     Root.update(load, 0)
    #
    # for ch in Adj[1]:
    #     dfs(1, ch)

    # Stack = [(0, 1, 0)]
    Stack = [(0, 0, 0)]
    parent, v, k = 0, 1, 0
    while v:
        if v != 1 and k == 0:
            load, toll = E[(parent, v)]
            Root.update(load, toll)
            for w, i in Queries[v]:
                Ans[i] = Root.query(1, w)
        if k < len(Adj[v]) and parent == Adj[v][k]:
            k += 1
        if k < len(Adj[v]):
            ch = Adj[v][k]
            k += 1
            Stack.append((parent, v, k))
            parent, v, k = v, ch, 0
        else:
            if v != 1:
                load, toll = E[(parent, v)]
                Root.update(load, 0)
            parent, v, k = Stack.pop()

    print(f"Case #{_c}: {' '.join(map(str, Ans))}")
