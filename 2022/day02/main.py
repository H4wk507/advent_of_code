def get_results_idx(enemy_choice: str, my_choice: str) -> int:
    return (
        (ord(enemy_choice) - ord("A")) * 2 % 3 + (ord(my_choice) - ord("X"))
    ) % 3


def get_my_choices_idx(enemy_choice: str, game_result: str) -> int:
    return (ord(enemy_choice) - ord("A") + ord(game_result) - ord("X")) % 3


def solve(strats: list[list[str]], part: int) -> int:
    choosing = {
        "X": 1,
        "Y": 2,
        "Z": 3,
    }
    total = 0
    game_results = [3, 6, 0]
    if part == 1:
        for enemy_choice, my_choice in strats:
            results_idx = get_results_idx(enemy_choice, my_choice)
            total += game_results[results_idx] + choosing[my_choice]
    if part == 2:
        my_choices = ["Z", "X", "Y"]
        for enemy_choice, game_result in strats:
            my_choices_idx = get_my_choices_idx(enemy_choice, game_result)
            my_choice = my_choices[my_choices_idx]
            results_idx = get_results_idx(enemy_choice, my_choice)
            total += game_results[results_idx] + choosing[my_choice]
    return total


if __name__ == "__main__":
    with open("input.txt") as f:
        strats = [strat.split() for strat in f.read().split("\n")[:-1]]
    print(f"Part #1: {solve(strats, 1)}")
    print(f"Part #2: {solve(strats, 2)}")
