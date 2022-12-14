import math


def handle_elevation(
    grid: list[list[str]], cx: int, cy: int, nx: int, ny: int
) -> bool:
    if grid[ny][nx] == "S":
        return ord(grid[cy][cx]) - 1 <= ord("a")
    else:
        return ord(grid[cy][cx]) - 1 <= ord(grid[ny][nx])


def get_start_cords(grid: list[list[str]], R: int, C: int) -> tuple[int, int]:
    return [(y, x) for x in range(C) for y in range(R) if grid[y][x] == "E"][0]


def solve(grid: list[list[str]], R: int, C: int) -> None:
    E = get_start_cords(grid, R, C)
    part1 = part2 = math.inf
    q = []
    q.append(E)
    # min dist to [i][j]
    dist = [[math.inf for _ in range(C)] for _ in range(R)]
    dist[E[0]][E[1]] = 0
    grid[E[0]][E[1]] = "z"
    while q:
        v = q.pop(0)
        vy, vx = v[0], v[1]
        if grid[vy][vx] == "S":
            part1 = min(part1, dist[vy][vx])
        if grid[vy][vx] in ("a", "S"):
            part2 = min(part2, dist[vy][vx])
        for dx, dy in [(-1, 0), (0, -1), (1, 0), (0, 1)]:
            ny = vy + dy
            nx = vx + dx
            if nx not in range(C) or ny not in range(R):
                continue
            curr_dist = dist[vy][vx] + 1
            if curr_dist < dist[ny][nx] and handle_elevation(
                grid, vx, vy, nx, ny
            ):
                dist[ny][nx] = curr_dist
                q.append((ny, nx))

    print(f"Part #1: {part1}")
    print(f"Part #2: {part2}")


if __name__ == "__main__":
    with open("input.txt") as f:
        rows = f.read().strip("\n").split("\n")
    R = len(rows)
    C = len(rows[0])
    grid = [[c for c in row] for row in rows]
    solve(grid, R, C)
