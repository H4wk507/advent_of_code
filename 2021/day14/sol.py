import sys
from collections import defaultdict

fn = sys.argv[1] if len(sys.argv) > 1 else "test.txt"

template = ""
pairins = defaultdict(str)
for line in open(fn):
    line = line.strip()
    if not template:
        template = line
    elif line:
        k, v = line.split(" -> ")
        pairins[k] = v

for k in range(40):
    i = 0
    while i < len(template) - 1:
        pair = template[i] + template[i + 1]
        if pair in pairins:
            # to usprawnic
            template = template[: i + 1] + pairins[pair] + template[i + 1 :]
            i += 1
        i += 1
    print(len(template))

occur = defaultdict(int)
for letter in template:
    occur[letter] += 1

# print(max(occur.values()) - min(occur.values()))
