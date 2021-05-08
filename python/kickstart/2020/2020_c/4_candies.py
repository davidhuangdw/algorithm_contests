# https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff43/0000000000337b4d#problem

def input_numbers(t=int):
    return map(t, input().split())


class Seg:
    DEFAULT = 0

    def __init__(self, l, r, v=DEFAULT, lc=None, rc=None):
        assert(l < r)
        self.l, self.r, self.v, self.lc, self.rc = l, r, v, lc, rc

    @staticmethod
    def _calc(lv, rv):
        return lv + rv

    def query(self, x, y):
        l, r = self.l, self.r
        if y < l or r-1 < x:
            return self.DEFAULT
        if x <= l and r-1 <= y:
            return self.v
        return self._calc(self.lc.query(x, y), self.rc.query(x, y))

    def update(self, x, v):
        l, r = self.l, self.r
        if not l <= x < r:
            return
        if r-l == 1:
            self.v = v
            return
        md = (l+r) >> 1
        if not self.lc:
            self.lc = Seg(l, md, 0)
            self.rc = Seg(md, r, 0)
        ch = self.lc if x < md else self.rc
        ch.update(x, v)
        self.v = self._calc(self.lc.v, self.rc.v)


def _sign(i):
    return 1 if i & 1 else -1


T, = input_numbers()
for _c in range(1, T+1):
    N, Q = input_numbers()
    A = list(input_numbers())
    V, W = Seg(1, N+1), Seg(1, N+1)
    ans = 0
    for i, a in enumerate(A):
        k = i+1
        V.update(k, a * _sign(k))
        W.update(k, a * k * _sign(k))
    for _ in range(Q):
        q, a, b = input().split()
        a, b = int(a), int(b)
        if q == 'U':
            V.update(a, b * _sign(a))
            W.update(a, b * a * _sign(a))
        else:
            w, v = W.query(a, b), V.query(a, b)
            ans += (w - (a-1)*v) * _sign(a)
    print(f"Case #{_c}: {ans}")
