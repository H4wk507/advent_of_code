import sys

infile = sys.argv[1] if len(sys.argv) > 1 else "test.txt"
lines = open(infile, "r").read().splitlines()

R = len(lines)
C = len(lines[0])


def check_lowest(rr, cc):
    # check every neighbor
    # so if current cell is [y][x]
    # check [y-1][x], [y+1][x], [y][x-1], [y][x+1]
    for y, x in (rr - 1, cc), (rr + 1, cc), (rr, cc - 1), (rr, cc + 1):
        # check out of bounds
        if 0 <= x < C and 0 <= y < R:
            # if neighbor is smaller return false
            if lines[y][x] <= lines[rr][cc]:
                return False

    return True


def solve(rr, cc, minval):
    # amount of basins is equal to amount of low points
    if (
        0 <= rr < R
        and 0 <= cc < C
        and (rr, cc) not in checked
        and lines[rr][cc] >= minval
        and lines[rr][cc] != "9"
    ):
        checked.append((rr, cc))
        solve(rr - 1, cc, minval)
        solve(rr + 1, cc, minval)
        solve(rr, cc - 1, minval)
        solve(rr, cc + 1, minval)


basins = []
for rr in range(R):
    for cc in range(C):
        checked = []
        if check_lowest(rr, cc):
            solve(rr, cc, lines[rr][cc])
            basins.append(len(checked))

basins.sort()
print(basins[-1] * basins[-2] * basins[-3])
