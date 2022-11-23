import sys

fn = sys.argv[1] if len(sys.argv) > 1 else "test.txt"

before, after = [], []
flag = 0
ymax, xmax = [0, 0]
with open(fn, "r") as f:
    for line in f:
        line = line.strip("\n")
        if not line:
            flag = 1
        elif flag:
            after.append(line.split()[-1])
        else:
            x, y = list(map(int, line.split(",")))
            if x > xmax:
                xmax = x
            if y > ymax:
                ymax = y
            before.append((x, y))


matrix = [[0 for _ in range(xmax + 1)] for _ in range(ymax + 1)]
for x, y in before:
    matrix[y][x] = 1

folds = 0
for line in after:
    cord = line[0]
    v = int(line[2:])
    if cord == "x":
        matrix = [
            [
                matrix[i][j] | matrix[i][k]
                for j, k in zip(range(v), range(xmax, v - 1, -1))
            ]
            for i in range(ymax + 1)
        ]
        xmax = v - 1
    elif cord == "y":
        matrix = [
            [matrix[i][j] | matrix[k][j] for j in range(xmax + 1)]
            for i, k in zip(range(v), range(ymax, v - 1, -1))
        ]
        ymax = v - 1

print(
    "\n".join(
        [
            "".join(["#" if matrix[i][j] else " " for j in range(xmax + 1)])
            for i in range(ymax + 1)
        ]
    )
)
