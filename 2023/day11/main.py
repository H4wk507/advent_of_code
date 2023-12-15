import numpy as np


def expand_grid(grid):
    new_grid = np.array([], dtype=object)

    for i, r in enumerate(grid):
        print(i, r)
        new_grid = np.insert(new_grid, i, r, axis=0)
        if np.all(r == '.'):
            new_grid = np.insert(new_grid, i, r, axis=0)

    for i, c in enumerate(grid.T):
        new_grid = np.insert(new_grid, i, c, axis=1)
        if np.all(c == '.'):
            new_grid = np.insert(new_grid, i, c, axis=1)

    return new_grid


with open("example") as f:
    g = np.array([np.array([x for x in r]) for r in f.read().strip('\n').split('\n')])

eg = expand_grid(g)
print(eg)
