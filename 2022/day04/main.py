def solve(assigment_list: list[list[list[int]]], part: int) -> int:
    cnt = 0
    for assig in assigment_list:
        first_start, first_end = assig[0][0], assig[0][1]
        second_start, second_end = assig[1][0], assig[1][1]
        if part == 1:
            if first_start <= second_start and first_end >= second_end:
                cnt += 1
            elif first_start >= second_start and first_end <= second_end:
                cnt += 1
        elif part == 2:
            if second_start <= first_end and second_end >= first_start:
                cnt += 1
    return cnt


if __name__ == "__main__":
    with open("input.txt") as f:
        # XD
        assigment_list = [
            [[int(z) for z in y.split("-")] for y in x.split(",")]
            for x in f.read().split("\n")[:-1]
        ]

    print(f"Part #1: {solve(assigment_list, 1)}")
    print(f"Part #2: {solve(assigment_list, 2)}")
