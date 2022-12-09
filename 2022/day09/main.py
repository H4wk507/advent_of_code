def follow(tail: tuple[int, int], head: tuple[int, int]) -> tuple[int, int]:
    """Make tail follow head, return new tail coordinates."""
    dx, dy = (head[0] - tail[0], head[1] - tail[1])
    abs_dx = abs(dx)
    abs_dy = abs(dy)
    # if tail touches the head, return the same coordinates
    if abs_dx <= 1 and abs_dy <= 1:
        return tail
    else:
        return (
            tail[0] + (0 if dx == 0 else dx // abs_dx),
            tail[1] + (0 if dy == 0 else dy // abs_dy),
        )


def solve(lines: list[str], part: int) -> int:
    directions = {"R": (1, 0), "L": (-1, 0), "U": (0, 1), "D": (0, -1)}
    lines = [line.split() for line in lines]
    lines = [(directions[c], int(dist)) for c, dist in lines]

    knots = 2 if part == 1 else 10
    # every knot coordinates, first knot is head, last is tail
    rope = [(0, 0) for _ in range(knots)]
    # visited coordinates by tail
    tail_visited = set({rope[-1]})
    for vec, dist in lines:
        for _ in range(dist):
            rope[0] = (rope[0][0] + vec[0], rope[0][1] + vec[1])
            for i in range(1, knots):
                rope[i] = follow(rope[i], rope[i - 1])
            tail_visited.add(rope[-1])
    return len(tail_visited)


if __name__ == "__main__":
    file = "input.txt"
    with open(file) as f:
        lines = f.read().strip("\n").split("\n")
    print(f"Part #1: {solve(lines, 1)}")
    print(f"Part #2: {solve(lines, 2)}")
