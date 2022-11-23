import sys
import math

fn = sys.argv[1] if len(sys.argv) > 1 else "test.txt"


with open(fn) as f:
    matrix = f.read()

r = len(matrix.strip().split("\n"))
c = len(matrix.split("\n")[0])

matrix = [[int(x) for x in y] * 5 for y in matrix.strip().split("\n")] * 5
R = len(matrix)
C = len(matrix[0])

print(r, c)
print(R, C)
dat = {}  # (x, y) -> weight
for y in range(R):
    for x in range(C):
        if matrix[y][x] + y // r + x // c <= 9:
            dat[(x, y)] = matrix[y][x] + y // r + x // c
        else:
            dat[(x, y)] = (matrix[y][x] + y // r + x // c) % 10 + 1

dist = [[math.inf for j in range(C)] for i in range(R)]
q = [(0, 0, 0)]  # row, column, weight
dist[0][0] = dat[(0, 0)]

while len(q) != 0:

    k = q[0]
    q.remove(k)

    dx = [-1, 0, 1, 0]
    dy = [0, 1, 0, -1]

    for i in range(4):
        x = k[0] + dx[i]
        y = k[1] + dy[i]

        if x < 0 or y < 0 or x >= R or y >= C:
            continue

        if dist[x][y] > dist[k[0]][k[1]] + dat[(x, y)]:
            if dist[x][y] != math.inf:
                if (x, y, dist[x][y]) in q:
                    q.remove((x, y, dist[x][y]))

            dist[x][y] = dist[k[0]][k[1]] + dat[(x, y)]
            q.append((x, y, dist[x][y]))


print(dist[R - 1][C - 1] - dat[(0, 0)])
