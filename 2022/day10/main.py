def run_cycle(sig_sum: int, cycles: int, x: int) -> tuple[int, int]:
    cycles += 1
    if (cycles - 20) % 40 == 0:
        return sig_sum + x * cycles, cycles
    else:
        return sig_sum, cycles


def draw_pixel(img: str, cycles: int, x: int) -> tuple[str, int]:
    # reset cycles for new row
    if cycles % 40 == 0:
        cycles = 0
        img += "\n"

    if abs(x - cycles) <= 1:
        img += "#"
    else:
        img += "."
    cycles += 1

    return img, cycles


def solve(lines: list[str], part: int) -> int | str:
    lines = [line.split() for line in lines]
    cycles, x = 0, 1
    sig_sum = 0
    img = ""

    for line in lines:
        match line:
            case ["noop"]:
                if part == 1:
                    sig_sum, cycles = run_cycle(sig_sum, cycles, x)
                elif part == 2:
                    img, cycles = draw_pixel(img, cycles, x)
            case ["addx", v]:
                for _ in range(2):
                    if part == 1:
                        sig_sum, cycles = run_cycle(sig_sum, cycles, x)
                    elif part == 2:
                        img, cycles = draw_pixel(img, cycles, x)
                x += int(v)

    return sig_sum if part == 1 else img


if __name__ == "__main__":
    with open("input.txt") as f:
        lines = f.read().strip("\n").split("\n")
    print(f"Part #1: {solve(lines, 1)}")
    print(f"Part #2: {solve(lines, 2)}")
