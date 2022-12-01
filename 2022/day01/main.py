def solve(calories_list: list[list[int]], part: int) -> int:
    sorted_calories_sum = sorted(
        [sum(calories) for calories in calories_list], reverse=True
    )
    if part == 1:
        return sum(sorted_calories_sum[:1])
    if part == 2:
        return sum(sorted_calories_sum[:3])


if __name__ == "__main__":
    with open("input.txt") as f:
        splitted_calories = f.read().split("\n\n")[:-1]
        calories_list = [
            list(map(int, calories.split("\n")))
            for calories in splitted_calories
        ]

        print(f"Part #1: {solve(calories_list, 1)}")
        print(f"Part #2: {solve(calories_list, 2)}")
