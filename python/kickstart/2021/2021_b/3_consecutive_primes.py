from math import sqrt


def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    Z, = input_numbers()
    x = int(sqrt(Z))
    l, r = max(2, x-2000), x+1000   # idea: prime gap < 1000

    k = int(sqrt(r))
    Pri = [1] * (k+1)
    SPri = [1] * (r-l+1)
    for i in range(2, k+1):
        if not Pri[i]:
            continue

        st = i*i
        while st <= k:
            Pri[st] = 0
            st += i

        st = (l+i-1)//i * i
        if st == i: st += i     # bug: st could be a prime
        while st <= r:
            SPri[st-l] = 0
            st += i

    primes = [v for v in range(l, r+1) if SPri[v-l]]
    for i in range(len(primes)-1):
        if primes[i] * primes[i+1] > Z:
            break
    print(f"Case #{_c}: {primes[i] * primes[i-1]}")
