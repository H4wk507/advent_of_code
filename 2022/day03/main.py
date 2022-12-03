from collections import defaultdict


def get_intersection(*args: str) -> str:
    seen: defaultdict[str, int] = defaultdict(int)
    for s in args:
        for c in set(s):
            seen[c] += 1
    return "".join([k for k, v in seen.items() if v == len(args)])


def sum_priorities(intersections: list[str]) -> int:
    total = 0
    for intersection in intersections:
        if intersection.islower():
            total += ord(intersection) - ord("a") + 1
        else:
            total += ord(intersection) - ord("A") + 27
    return total


def solve(items_list: list[str], part: int) -> int:
    intersections = []
    if part == 1:
        for items in items_list:
            mid = len(items) // 2
            left_half = items[:mid]
            right_half = items[mid:]
            intersection = get_intersection(left_half, right_half)
            intersections.append(intersection)
    if part == 2:
        for i in range(0, len(items_list), 3):
            intersection = get_intersection(
                items_list[i], items_list[i + 1], items_list[i + 2]
            )
            intersections.append(intersection)
    return sum_priorities(intersections)


if __name__ == "__main__":
    with open("input.txt") as f:
        items_list = f.read().split("\n")[:-1]

    print(f"Part #1: {solve(items_list, 1)}")
    print(f"Part #2: {solve(items_list, 2)}")
