def to_dec(s):
    res = 0
    pot = 1
    for i in s[::-1]:
        res = res + (pot * i)
        pot *= 2

    return res


def solve(arr, idx, minmax):
    if len(arr) == 1:
        return to_dec(arr[0])

    column = [num[idx] for num in arr]
    bits = [0, 0]

    for i in column:
        bits[i] += 1

    if bits[0] == bits[1]:
        v = 1 if minmax == max else 0
    else:
        v = bits.index(minmax(bits))

    arr = list(filter(lambda x: x[idx] == v, arr))
    return solve(arr, idx + 1, minmax)


def main():
    arr = [
        [int(x) for x in line.strip()]
        for line in open("input.txt", "r").read().splitlines()
    ]

    idx = 0
    print(solve(arr, idx, min) * solve(arr, idx, max))


main()
