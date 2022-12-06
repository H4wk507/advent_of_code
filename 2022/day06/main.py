def is_unique_string(s: str) -> bool:
    return len(set(s)) == len(s)


def solve(code: str, part: int) -> int:
    code_length = 4 if part == 1 else 14

    if is_unique_string(code[:code_length]):
        return code_length

    for idx in range(1, len(code) - code_length):
        packet = code[idx : idx + code_length]
        if is_unique_string(packet):
            return idx + code_length

    return -1


if __name__ == "__main__":
    with open("input.txt") as f:
        code = f.read().strip("\n")
    print(f"Part #1: {solve(code, 1)}")
    print(f"Part #2: {solve(code, 2)}")
