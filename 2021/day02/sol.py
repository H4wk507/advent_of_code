import sys

infile = sys.argv[1] if len(sys.argv) > 1 else "test.txt"

lines = [
    (d, int(v))
    for line in open(infile).read().splitlines()
    for d, v in [line.split()]
]


def solve():
    x, y, aim = [0] * 3

    for pos, val in lines:
        if pos == "forward":
            x += val
            y += aim * val
        else:
            aim += val * (-1 if pos == "up" else 1)

    return x * y


print(f"Part #2: {solve()}")
