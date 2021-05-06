def input_numbers(t=int):
    return map(t, input().split())

def J(n):
    x = 1
    for i in range(2, n+1):
        x *= i
    return x

print(J(40))
print(J(40) // J(20) // J(20))



def gcd(x, y):
    if y == 0:
        return x
    return gcd(y, x % y)


T, = input_numbers()
for _c in range(1, T+1):
    N, Q = input_numbers()
    A = []
    for _ in range(N):
        a, k = input().split()
        A.append((int(k), sum(1 << i if a[i] == 'T' else 0 for i in range(Q))))
    cnt = [[0, 0] for i in range(Q)]
    for i in range(1<<Q):
        poss = True
        for (n, a) in A:
            xr = a ^ i
            if n != sum(0 if (1 << k) & xr else 1 for k in range(Q)):
                poss = False
                break
        if poss:
            for j in range(Q):
                cnt[j][1 if (1 << j) & i else 0] += 1

    ans = ''
    z, w = 0, 1
    for i in range(Q):
        if cnt[i][0] >= cnt[i][1]:
            ans += 'F'
        else:
            ans += 'T'
        a, b = cnt[i]
        # if a > 0 and b > 0:
        a, b = max(a, b), a+b
        z, w = z*b + a*w, b*w
        d = gcd(z, w)
        z, w = z//d, w//d

    print(f"Case #{_c}: {ans} {z}/{w}")
