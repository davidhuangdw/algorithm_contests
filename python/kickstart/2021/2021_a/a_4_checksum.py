def input_numbers(t=int):
    return map(t, input().split())


def poss(bits, a):
    row = [set() for _ in range(N)]
    col = [set() for _ in range(N)]
    all = []
    for i in range(N):
        for j in range(N):
            x = 1 << (i*N + j)
            if not (x & bits) and a[i][j] == 1:
                row[i].add((i, j))
                col[j].add((i, j))
                all.add((i, j))
    while all:
        found = False
        for t in range(N):
            if len(row[t]) == 1:
                found = True
                i, j = row[t].pop()
                all.remove((i, j))
                col[j].remove((i, j))
        for t in range(N):
            if len(col[t]) == 1:
                found = True
                i, j = col[t].pop()
                all.remove((i, j))
                row[i].remove((i, j))
        if not found:
            return False
    return True


def get_cost(bits, b):
    s = 0
    for i in range(N):
        for j in range(N):
            x = 1 << (i*N + j)
            if x & bits:
                s += b[i][j]
    return s


T, = input_numbers()
for _c in range(1, T+1):
    N, = input_numbers()
    A, B = [], []
    for i in range(N):
        A.append(list(input_numbers()))
        for j in range(N):
            A[i][j] = 1 if A[i][j] == -1 else 0
    for i in range(N):
        B.append(list(input_numbers()))
    input_numbers()
    input_numbers()
    Cost = 1000 * N * N
    for Bits in range(1 << (N*N)):
        if poss(Bits, A):
            Cost = min(Cost, get_cost(Bits, B))
    print(f"Case #{_c}: {Cost}")
