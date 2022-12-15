import sys
import re

def print_map(grid, xbounds, ybounds):
    x_start, x_end = xbounds
    y_start, y_end = ybounds
    for y in range(y_start, y_end+1):
        for x in range(x_start, x_end+1):
            if (x, y) in grid:
                print(grid[(x, y)], end="")
            else:
                print(".", end="")
        print()

if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        lines = f.read().strip("\n").split("\n")

    pattern = r"^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at " \
              r"x=(-?\d+), y=(-?\d+)$"
    xpos = set() 
    for line in lines:
        sx, sy, bx, by = map(int, re.match(pattern, line).groups())
        y_distance = abs(2_000_000 - sy)
        beacon_distance = abs(sx - bx) + abs(sy - by)
        if y_distance <= beacon_distance:
            for x in range(sx-(beacon_distance-y_distance),sx+(beacon_distance-y_distance)):
                xpos.add(x)

    # tuning frequency = x * 4000000 + y
    print(len(xpos))
