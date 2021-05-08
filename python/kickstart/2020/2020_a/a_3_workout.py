import heapq


def input_numbers(t=int):
    return map(t, input().split())


def possible(diffs, k, x):
    for d in diffs:
        k -= (d+x-1)//x - 1
        if k < 0:
            return False
    return True


T, = input_numbers()
for _c in range(1, T+1):
    N, K = input_numbers()
    A = list(input_numbers())
    # wrong greedy!!: maybe better to cut into 3 parts than 2 parts
    #
    # F = []
    # for i in range(len(A)-1):
    #     heapq.heappush(F, -(A[i+1] - A[i]))
    #
    # for _ in range(K):
    #     x = -heapq.heappop(F)
    #     heapq.heappush(F, -(x//2))
    #     heapq.heappush(F, -(x - x//2))
    # print(f"Case #{_c}: {-F[0]}")

    D = [A[i+1] - A[i] for i in range(N-1)]
    l, r = 1, A[-1]
    while l <= r:
        md = (l+r)//2
        if possible(D, K, md):
            r = md - 1
        else:
            l = md + 1
    print(f"Case #{_c}: {l}")
