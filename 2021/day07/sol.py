import sys


def main(lines: [int], part: int) -> int:
    # ------------------------------------------
    # Part #1

    if part == 1:
        n = len(lines)
        mediana = (
            lines[n // 2]
            if n % 2
            else (lines[n // 2] + lines[n // 2 - 1]) // 2
        )
        return int(sum([abs(x - mediana) for x in lines]))

    # ------------------------------------------
    # Part #2

    elif part == 2:
        res = []
        for i in range(min(lines), max(lines) + 1):
            curr = 0
            for el in lines:
                curr += (1 + abs(el - i)) / 2 * abs(el - i)
            res.append(curr)

        return int(min(res))


if __name__ == "__main__":
    if len(sys.argv) < 2:
        exit("Error: provide a filename.")

    lines = [int(x) for x in open(sys.argv[1], "r").read().split(",")]
    lines.sort()
    print(f"Part #1: {main(lines, part=1)}")
    print(f"Part #2: {main(lines, part=2)}")
