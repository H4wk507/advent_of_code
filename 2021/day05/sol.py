#!/usr/bin/python3

import sys


def main():
    if len(sys.argv) < 2:
        sys.exit("Error: Provide a filename.")

    arr = [
        line.replace(" -> ", ",").split(",")
        for line in open(sys.argv[1], "r").read().splitlines()
    ]

    x0, y0, xk, yk = list(), list(), list(), list()
    for el in arr:
        x0.append(int(el[0]))
        y0.append(int(el[1]))
        xk.append(int(el[2]))
        yk.append(int(el[3]))

    solve(x0, y0, xk, yk, arr, 1)
    solve(x0, y0, xk, yk, arr, 2)


def solve(x0, y0, xk, yk, arr, part):
    board = [[0 for _ in range(1001)] for _ in range(1001)]
    for i in range(len(arr)):

        # vertical
        if x0[i] == xk[i]:
            start = min(yk[i], y0[i])
            for j in range(start, abs(yk[i] - y0[i]) + start + 1):
                board[j][x0[i]] += 1

        # horizontal
        elif y0[i] == yk[i]:
            start = min(xk[i], x0[i])
            for j in range(start, abs(xk[i] - x0[i]) + start + 1):
                board[y0[i]][j] += 1

        # diagonal
        elif part == 2:
            startx = min(x0[i], xk[i])
            # 9, 7 -> 7, 9 wieksze mniejsze minx maxy +x -y
            # 0, 0 -> 8, 8 wieksze wieksze minx miny +x +y
            # 5, 5 -> 8, 2 mniejsze wieksze minx maxy +x -y
            # 6, 4 -> 2, 0 mniejsze mniejsze minx miny +x +y
            if (xk[i] > x0[i] and yk[i] > y0[i]) or (
                xk[i] < x0[i] and yk[i] < y0[i]
            ):

                starty = min(y0[i], yk[i])
                for j in range(startx, startx + abs(x0[i] - xk[i]) + 1):
                    board[starty][j] += 1
                    starty += 1
            else:
                starty = max(y0[i], yk[i])
                for j in range(startx, startx + abs(x0[i] - xk[i]) + 1):
                    board[starty][j] += 1
                    starty -= 1

    cnt = sum([board[i][j] > 1 for i in range(1001) for j in range(1001)])
    print(f"Part #{part}: {cnt}")


main()
