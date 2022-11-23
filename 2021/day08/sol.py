import sys
from collections import defaultdict

infile = sys.argv[1] if len(sys.argv) > 1 else "test.txt"

ans = 0
for line in open(infile, "r"):
    before, after = line.split("|")
    before = before.split()
    after = after.split()

    dct = defaultdict(list)  # word length -> list of words
    for b in before:
        dct[len(b)].append(b)
    for a in after:
        if len(dct[len(a)]) == 1:
            ans += 1
print(ans)
