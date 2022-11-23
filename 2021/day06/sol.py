import sys
from collections import Counter
from functools import reduce


def main(input: [int], days: int) -> int:
    def helper(counter: dict, _):
        timerToCnt = {
            timer - 1: population for timer, population in counter.items()
        }
        if -1 in timerToCnt:
            timerToCnt[8] = timerToCnt.pop(-1)
            timerToCnt[6] = timerToCnt[8] + timerToCnt.get(6, 0)
        return timerToCnt

    return sum(reduce(helper, range(0, days), dict(Counter(arr))).values())


if __name__ == "__main__":
    if len(sys.argv) < 2:
        exit("Error: provide a filename.")

    arr = [
        int(x)
        for x in open(sys.argv[1], "r").read().replace("\n", "").split(",")
    ]

    print(f"Part #1: {main(arr, 80)}")
    print(f"Part #2: {main(arr, 256)}")
