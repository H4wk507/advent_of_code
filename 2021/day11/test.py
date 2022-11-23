def neighbors(arr, x, y):
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
        if 0 <= i <= len(arr) and 0 <= j <= len(arr[0]):
            arr[i][j] += 1


def test():
    arr = [[0] * 3 for _ in range(3)]
    print(arr)
    neighbors(arr, 1, 1)
    print(arr)
