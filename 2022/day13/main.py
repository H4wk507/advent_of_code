from functools import cmp_to_key
from math import prod


def parse(s: str) -> list:
    # advent of parsing
    i = 1
    n = len(s)
    lst = []
    while i < n:
        # parse number
        if s[i].isdecimal():
            j = i + 1
            while s[j].isdecimal():
                j += 1
            lst.append(int(s[i:j]))
            i = j
        # parse lists
        elif s[i] == "[":
            depth = 0
            for j in range(i, n):
                depth += s[j] == "["
                depth -= s[j] == "]"
                if depth == 0:
                    sblist = parse(s[i : j + 1])
                    lst.append(sblist)
                    i += j - 1
                    break
        # skip whitespace, commas and outer closing bracket
        else:
            i += 1
    return lst


def compare(l: int | list, r: int | list) -> int:
    if isinstance(l, int) and isinstance(r, int):
        return (l > r) - (l < r)
    elif isinstance(l, int) and isinstance(r, list):
        return compare([l], r)
    elif isinstance(l, list) and isinstance(r, int):
        return compare(l, [r])
    elif isinstance(l, list) and isinstance(r, list):
        # loop until you find different corresponding elements,
        # then return comparator, if whole lists are the same
        # return according to lists' length
        for a, b in zip(l, r):
            cmp = compare(a, b)
            if cmp != 0:
                return cmp
        return compare(len(l), len(r))


def solve(pairs: list[str]) -> None:
    pairs = [[*map(parse, pair.split("\n"))] for pair in pairs]
    sum_of_indices = sum(
        i for i, pair in enumerate(pairs, 1) if compare(*pair) == -1
    )
    print(f"Part #1: {sum_of_indices}")
    pairs = sorted(sum(pairs, [[2], [6]]), key=cmp_to_key(compare))
    decoder_key = prod(
        i for i, pair in enumerate(pairs, 1) if pair in [[2], [6]]
    )
    print(f"Part #2: {decoder_key}")


if __name__ == "__main__":
    with open("input.txt") as f:
        pairs = f.read().strip("\n").split("\n\n")
    solve(pairs)
