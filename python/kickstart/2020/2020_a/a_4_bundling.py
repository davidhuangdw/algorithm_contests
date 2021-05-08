def input_numbers(t=int):
    return map(t, input().split())


OFFSET = ord('A')
T, = input_numbers()
for _c in range(1, T+1):
    N, K = input_numbers()
    Nodes = [[0, {}]]
    for _ in range(N):
        cur = Nodes[0]
        word = input()
        for c in word:
            # i = ord(c) - ord('A')
            if c not in cur[1]:
                Nodes.append([0, {}])
                cur[1][c] = len(Nodes)-1
            cur = Nodes[cur[1][c]]
        cur[0] += 1
    score = 0

    def dfs(i, h):
        global score
        node = Nodes[i]
        sub = 0
        for _, j in node[1].items():
            sub += dfs(j, h+1)
        sub += node[0]
        score += sub//K * h
        return sub % K
    dfs(0, 0)
    print(f"Case #{_c}: {score}")
