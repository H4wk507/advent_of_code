import re
import sys
from functools import cache


def read_input(
    lines: list[str],
) -> tuple[dict[str, int], dict[str, list[str]]]:
    pattern = r"^Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? ([a-zA-Z, ]+)$"
    rates = {}
    links = {}
    for line in lines:
        src, rate, dst = re.match(pattern, line).groups()
        rates[src] = int(rate)
        links[src] = [x.strip() for x in dst.split(",")]
    return rates, links


def solve(
    rates: dict[str, int],
    links: dict[str, list[str]],
    start: str,
    minutes_left: int,
) -> int:
    @cache
    def dfs(valve: str, minutes_left: int, visited: tuple[str]) -> int:
        if minutes_left <= 0:
            return 0
        res = 0
        for link in links[valve]:
            res = max(res, dfs(link, minutes_left - 1, visited))

        if valve not in visited and rates[valve] > 0:
            visited = tuple(sorted(visited + (valve,)))
            res = max(
                res,
                dfs(valve, minutes_left - 1, visited)
                + rates[valve] * (minutes_left - 1),
            )
        return res

    return dfs(start, minutes_left, ())


if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        lines = f.read().strip("\n").split("\n")
    rates, links = read_input(lines)
    print(f"Part #1: {solve(rates, links, 'AA', 30)}")
