#!/usr/bin/python3
import sys

infile = sys.argv[1] if len(sys.argv) > 1 else "test.txt"

matrix = [[int(y) for y in x] for x in open(infile, "r").read().split("\n")]

R = len(matrix)
C = len(matrix[0])


def flash(x, y):
    global suma
    suma += 1
    matrix[x][y] = -1  # flashed
    for i, j in (
        (x, y - 1),
        (x - 1, y - 1),
        (x - 1, y),
        (x - 1, y + 1),
        (x, y + 1),
        (x + 1, y + 1),
        (x + 1, y),
        (x + 1, y - 1),
    ):

        if 0 <= i < R and 0 <= j < C and matrix[i][j] != -1:
            matrix[i][j] += 1
            if matrix[i][j] > 9:
                flash(i, j)


suma = 0
steps = 100
while 1:
    for i in range(R):
        for j in range(C):
            matrix[i][j] += 1

    for i in range(R):
        for j in range(C):
            if matrix[i][j] > 9:
                flash(i, j)

    for i in range(R):
        for j in range(C):
            if matrix[i][j] == -1:
                matrix[i][j] = 0

print(suma)
