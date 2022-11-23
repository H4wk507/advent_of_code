import sys


def solve(lines: [int], part: int) -> int:
    if part == 2:
        return sum([lines[i] < lines[i + 3] for i in range(len(lines) - 3)])
    elif part == 1:
        return sum([lines[i] < lines[i + 1] for i in range(len(lines) - 1)])


if __name__ == "__main__":
    if len(sys.argv) < 2:
        exit("Error: provide a filename.")

    lines = [int(line[:-1]) for line in open(sys.argv[1], "r")]
    print(f"Part #1: {solve(lines, part=1)}")
    print(f"Part #2: {solve(lines, part=2)}")
