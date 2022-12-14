import math


def out_of_bounds(
    x: int, y: int, xbounds: tuple[int, int], ybounds: tuple[int, int]
) -> bool:
    return x not in range(*xbounds) or y not in range(*ybounds)


def read_input(
    lines: list[str],
) -> tuple[dict[tuple[int, int], str], tuple[int, int], tuple[int, int]]:
    rocks = {}
    xbounds = math.inf, -math.inf
    ybounds = 0, -math.inf

    for line in lines:
        cords = [x.strip() for x in line.split("->")]
        x, y = [int(x) for x in cords[0].split(",")]
        for cord in cords[1:]:
            nx, ny = [int(_) for _ in cord.split(",")]
            if nx == x:
                low = min(y, ny)
                high = max(y, ny)
                for dy in range(low, high + 1):
                    ybounds = min(ybounds[0], dy), max(ybounds[1], dy)
                    rocks[(x, dy)] = "#"
            else:
                low = min(x, nx)
                high = max(x, nx)
                for dx in range(low, high + 1):
                    xbounds = min(xbounds[0], dx), max(xbounds[1], dx)
                    rocks[(dx, y)] = "#"
            x, y = nx, ny
    return rocks, xbounds, ybounds


def part1(
    rocks: dict[tuple[int, int], str],
    xbounds: tuple[int, int],
    ybounds: tuple[int, int],
) -> dict[tuple[int, int], str]:
    sand = {}
    sx, sy = 500, 0
    turn = 0
    while True:
        if (sx, sy + 1) in rocks:
            if (sx - 1, sy + 1) in rocks:
                if (sx + 1, sy + 1) in rocks:
                    rocks[(sx, sy)] = sand[(sx, sy)] = "o"
                    sx, sy = 500, 0
                else:
                    # fall right
                    sy += 1
                    sx += 1
                    if out_of_bounds(sx, sy, xbounds, ybounds):
                        break
            else:
                # fall left
                sy += 1
                sx -= 1
                if out_of_bounds(sx, sy, xbounds, ybounds):
                    break
        else:
            # fall down
            sy += 1
            if out_of_bounds(sx, sy, xbounds, ybounds):
                break
        turn += 1
    print(f"Part #1: {len(sand)} computed in {turn} turns")
    return sand


def part2(
    rocks: dict[tuple[int, int], str],
    sand: dict[tuple[int, int], str],
    xbounds: tuple[int, int],
    ybounds: tuple[int, int],
) -> None:
    floory = ybounds[1] + 2
    sx, sy = 500, 0
    turn = 0
    while True:
        if (sx, sy) in rocks:
            break
        elif (sx, sy + 1) in rocks or sy + 1 == floory:
            if (sx - 1, sy + 1) in rocks or sy + 1 == floory:
                if (sx + 1, sy + 1) in rocks or sy + 1 == floory:
                    rocks[(sx, sy)] = sand[(sx, sy)] = "o"
                    sx, sy = 500, 0
                else:
                    # fall right
                    sy += 1
                    sx += 1
            else:
                # fall left
                sy += 1
                sx -= 1
        else:
            # fall down
            sy += 1
        turn += 1
    print(f"Part #2: {len(sand)} computed in {turn} turns")


if __name__ == "__main__":
    with open("input.txt") as f:
        lines = f.read().strip("\n").split("\n")

    rocks, xbounds, ybounds = read_input(lines)
    sand = part1(rocks, xbounds, ybounds)
    part2(rocks, sand, xbounds, ybounds)
    # Computed in 0.845s using python3.11
