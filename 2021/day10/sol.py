# maybe not the best solution but it works

part1 = {")": 3, "]": 57, "}": 1197, ">": 25137}
part2 = {")": 1, "]": 2, "}": 3, ">": 4}
match = {"(": ")", "{": "}", "[": "]", "<": ">"}

lines = open("input.txt", "r").read().strip().split("\n")

ans = []
suma1 = 0


def solve(testcase):
    stack = []
    suma2 = 0
    global suma1
    for sym in testcase:

        if sym in "([{<":
            stack.append(sym)

        elif sym in ")]}>":
            if match[stack[-1]] == sym:
                stack.pop()  # pop last element

            else:  # not correct
                suma1 += part1[sym]  # corrupted line
                return

    while len(stack) > 0:
        suma2 = suma2 * 5 + part2[match[stack[-1]]]  # incompleted line
        stack.pop()

    ans.append(suma2)


for testcase in lines:
    solve(testcase)

print(f"Part #1 {suma1}")
print(f"Part #2 {sorted(ans)[len(ans) // 2]}")
