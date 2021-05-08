def input_numbers(t=int):
    return map(t, input().split())


def rev(arr, a, b):
    while a < b:
        arr[a], arr[b] = arr[b], arr[a]
        a += 1
        b -= 1


T, = input_numbers()
for _c in range(1, T+1):
    N, = input_numbers()
    A = list(input_numbers())
    cnt = 0
    for i in range(N-1):
        mi = i
        for j in range(i+1, N):
            if A[j] < A[mi]:
                mi = j
        cnt += mi - i + 1
        rev(A, i, mi)

    print(f"Case #{_c}: {cnt}")
