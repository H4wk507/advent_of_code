import math


def handle_elevation(grid, cx, cy, nx, ny):
    if (cx, cy) == (0, 0):
        return True

    if grid[cy][cx] == "S":
        return ord("a") + 1 >= ord(grid[ny][nx])

    if grid[ny][nx] == "E":
        return ord(grid[cy][cx]) + 1 >= ord("z")
    return ord(grid[cy][cx]) + 1 >= ord(grid[ny][nx])


def get_start_cords(grid, R, C):
    for y in range(R):
        for x in range(C):
            if grid[y][x] == "S":
                return (y, x)
    return (-1, -1)


def get_a_cords(grid, R, C, lst):
    for y in range(R):
        for x in range(C):
            if grid[y][x] == "a":
                lst.append((y, x))


def solve(rows: list[str], part):
    R = len(rows)
    C = len(rows[0])
    grid = [[c for c in row] for row in rows]

    S = get_start_cords(grid, R, C)
    start_indices = [S]
    if part == 2:
        get_a_cords(grid, R, C, start_indices)
    global_min = math.inf
    for ay, ax in start_indices:
        q = []
        q.append((ay, ax))
        # min dist to [i][j]
        dist = [[math.inf for _ in range(C)] for _ in range(R)]
        dist[ay][ax] = 0
        while q:
            v = q.pop(0)
            vy, vx = v[0], v[1]
            if grid[vy][vx] == "E":
                global_min = min(global_min, dist[vy][vx])
                break
            for dx, dy in [(-1, 0), (0, -1), (1, 0), (0, 1)]:
                ny = vy + dy
                nx = vx + dx

                if nx not in range(C) or ny not in range(R):
                    continue

                curr_dist = dist[vy][vx] + 1
                if dist[ny][nx] > curr_dist and handle_elevation(
                    grid, vx, vy, nx, ny
                ):
                    dist[ny][nx] = curr_dist
                    q.append((ny, nx))
    return global_min


if __name__ == "__main__":
    file = "input.txt"
    with open(file) as f:
        rows = f.read().strip("\n").split("\n")

    print(f"Part #1: {solve(rows, 1)}")
    print(f"Part #2: {solve(rows, 2)}")
