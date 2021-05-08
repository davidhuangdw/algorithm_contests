import heapq


def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    R, C = input_numbers()
    H = []
    for _ in range(R):
        H.append(list(input_numbers()))
    que = []
    for i in range(R):
        for j in range(C):
            heapq.heappush(que, (-H[i][j], i, j))

    cnt = 0
    while que:
        h, i, j = heapq.heappop(que)
        h = -h
        if H[i][j] > h:
            continue
        for x, y in ((i-1, j), (i+1, j), (i, j-1), (i, j+1)):
            if 0 <= x < R and 0 <= y < C and H[x][y] < h-1:
                cnt += h-1 - H[x][y]
                H[x][y] = h-1
                heapq.heappush(que, (-H[x][y], x, y))
    print(f"Case #{_c}: {cnt}")
