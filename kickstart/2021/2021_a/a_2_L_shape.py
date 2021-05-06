def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    R, C = input_numbers()
    S = []
    Up = [[0] * C for _ in range(R)]
    Down = [[0] * C for _ in range(R)]
    Left = [[0] * C for _ in range(R)]
    Right = [[0] * C for _ in range(R)]
    for i in range(R):
        S.append(list(input_numbers()))
        for j in range(C):
            if S[i][j]:
                Up[i][j] = 1 + (Up[i-1][j] if i-1 >= 0 else 0)
                Left[i][j] = 1 + (Left[i][j-1] if j-1 >= 0 else 0)
        for j in range(C-1, -1, -1):
            if S[i][j]:
                Right[i][j] = 1 + (Right[i][j+1] if j+1 < C else 0)
    for i in range(R-1, -1, -1):
        for j in range(C):
            if S[i][j]:
                Down[i][j] = 1 + (Down[i+1][j] if i+1 < R else 0)

    cnt = 0
    CNT = [[0] * C for _ in range(R)]
    for i in range(R):
        for j in range(C):
            if not S[i][j]:
                continue
            for a in (Left[i][j], Right[i][j]):
                for b in (Up[i][j], Down[i][j]):
                    cnt += max(min(a, b//2) - 1, 0)
                    cnt += max(min(b, a//2) - 1, 0)
    print(f"Case #{_c}: {cnt}")
