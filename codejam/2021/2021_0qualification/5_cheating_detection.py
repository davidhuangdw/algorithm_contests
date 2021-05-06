def input_numbers(t=int):
    return map(t, input().split())


K, N = 100, 10000
T, = input_numbers()
P, = input_numbers()
for _c in range(1, T+1):
    A = []
    for i in range(K):
        A.append(input())
    prob = sorted(range(N), key=lambda i: sum(A[k][i] == '0' for k in range(K)))

    def metric(k):
        l = sum(1 for i in range(500) if A[k][prob[i]] == '0')
        r = sum(1 for i in range(N-500, N) if A[k][prob[i]] == '1')
        return (l+1) * (r+1)

    def inv(k):
        s = 0
        pre0 = 0
        for i in range(N):
            if A[k][prob[i]] == '0':
                pre0 += 1
            else:
                s += pre0
        return s * 1.0 / pre0 / (N-pre0)    # why?????

    res = max(range(K), key=lambda k: inv(k))
    print(f"Case #{_c}: {res+1}")
    # for k in range(100):
    #     print(''.join(A[i][k] for i in range(100)))
    # print(''.join(A[res][prob[i]] for i in range(N)))
    # print(''.join(A[58][prob[i]] for i in range(N)))
    # print(inv(res), metric(res))
    # print(inv(58), metric(58))
