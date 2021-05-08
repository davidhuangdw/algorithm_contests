from collections import deque


def input_numbers(t=int):
    return map(t, input().split())


# T, = input_numbers()
# for _c in range(1, T+1):
#     M, = input_numbers()
#     P = [0] * M
#     N = [0] * M
#     S = 0
#     for i in range(M):
#         P[i], N[i] = input_numbers()
#         S += P[i] * N[i]
#     ans = 0
#     que = deque([])
#     que.append((1, 0, [0] * M))
#     while que:
#         mul, s, cnt = que.pop()
#         for i in range(M-1, -1, -1):
#             if cnt[i]+1 > N[i] or S - s - P[i] <= ans:
#                 continue
#             cur = mul * P[i] + s + P[i]
#             if cur >= S:
#                 if cur == S:
#                     ans = S - s - P[i]
#                     # print([cnt[j] + (1 if j == i else 0) for j in range(M)])
#                 continue
#             else:
#                 que.appendleft((mul*P[i], s + P[i], [cnt[j] + (1 if j == i else 0) for j in range(M)]))
#     print(f"Case #{_c}: {ans}")


T, = input_numbers()
for _c in range(1, T+1):
    M, = input_numbers()
    P = [0] * M
    N = [0] * M
    S = 0
    for i in range(M):
        P[i], N[i] = input_numbers()
        S += P[i] * N[i]
    ans = 0
    for s in range(2, min(30000, S-1)):
        rem = S - s
        acc = 0
        for i in range(M):
            for j in range(N[i]):
                if rem % P[i] == 0:
                    rem //= P[i]
                    acc += P[i]
                else:
                    break
        if rem == 1 and acc == s:
            ans = S - s
            break
    print(f"Case #{_c}: {ans}")
