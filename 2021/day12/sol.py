import sys
from collections import defaultdict

infile = sys.argv[1] if len(sys.argv) > 1 else "test.txt"

# dictionary node -> array of possible paths
# small caves only 1 time

ans = 0


def solve(begin, visited, twice) -> None:
    global ans

    if begin == "end":
        ans += 1
        return

    if visited[begin] > 0 and twice:
        return

    if begin in small_caves:
        visited[begin] += 1
        if visited[begin] == 2:
            twice = True

    for node in graph[begin]:
        if node != "start":
            solve(node, visited, twice)

    visited[begin] -= 1


graph = defaultdict(list)
small_caves = set()

for line in open(infile, "r").read().splitlines():
    start, end = line.split("-")
    graph[start].append(end)
    graph[end].append(start)

    if start.islower():
        small_caves.add(start)
    if end.islower():
        small_caves.add(end)


solve("start", defaultdict(int), False)
print(ans)
