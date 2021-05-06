def input_numbers(t=int):
    return map(t, input().split())


T, = input_numbers()
for _c in range(1, T+1):
    print(f"Case #{_c}: ")
