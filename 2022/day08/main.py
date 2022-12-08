from enum import Enum


class Direction(Enum):
    LEFT = 1
    UP = 2
    RIGHT = 3
    DOWN = 4


def check_row(
    grid: list[str], start: int, end: int, nrow: int, height: str
) -> bool:
    return all(grid[nrow][i] < height for i in range(start, end + 1))


def check_column(
    grid: list[str], start: int, end: int, ncol: int, height: str
) -> bool:
    return all(grid[i][ncol] < height for i in range(start, end + 1))


def is_visible(grid: list[str], r: int, c: int, nrow: int, ncol: int) -> bool:
    return (
        check_row(grid, 0, ncol - 1, nrow, grid[nrow][ncol])
        or check_row(grid, ncol + 1, c - 1, nrow, grid[nrow][ncol])
        or check_column(grid, 0, nrow - 1, ncol, grid[nrow][ncol])
        or check_column(grid, nrow + 1, r - 1, ncol, grid[nrow][ncol])
    )


def get_view_range(
    grid: list[str], r: int, c: int, nrow: int, ncol: int, direction: Direction
) -> int:
    total = 0
    if direction == Direction.LEFT:
        start = ncol - 1
        for i in range(start, -1, -1):
            if grid[nrow][i] < grid[nrow][ncol]:
                total += 1
            else:
                return total + 1
    elif direction == Direction.RIGHT:
        start = ncol + 1
        for i in range(start, c):
            if grid[nrow][i] < grid[nrow][ncol]:
                total += 1
            else:
                return total + 1
    elif direction == Direction.UP:
        start = nrow - 1
        for i in range(start, -1, -1):
            if grid[i][ncol] < grid[nrow][ncol]:
                total += 1
            else:
                return total + 1
    elif direction == Direction.DOWN:
        start = nrow + 1
        for i in range(start, r):
            if grid[i][ncol] < grid[nrow][ncol]:
                total += 1
            else:
                return total + 1
    return total


def get_scenic_score(
    grid: list[str], r: int, c: int, nrow: int, ncol: int
) -> int:
    return (
        get_view_range(grid, r, c, nrow, ncol, Direction.LEFT)
        * get_view_range(grid, r, c, nrow, ncol, Direction.UP)
        * get_view_range(grid, r, c, nrow, ncol, Direction.RIGHT)
        * get_view_range(grid, r, c, nrow, ncol, Direction.DOWN)
    )


def solve(grid: list[str], part: int) -> int:
    r = len(grid)
    c = len(grid[0])
    if part == 1:
        return sum(
            rr in (0, r - 1)
            or cc in (0, c - 1)
            or is_visible(grid, r, c, rr, cc)
            for rr in range(r)
            for cc in range(c)
        )
    elif part == 2:
        return max(
            get_scenic_score(grid, r, c, rr, cc)
            for rr in range(r)
            for cc in range(c)
            if rr not in (0, r - 1) and cc not in (0, c - 1)
        )
    return -1


if __name__ == "__main__":
    file = "input.txt"
    with open(file) as f:
        grid = f.read().strip("\n").split("\n")
    print(f"Part #1: {solve(grid, 1)}")
    print(f"Part #2: {solve(grid, 2)}")
