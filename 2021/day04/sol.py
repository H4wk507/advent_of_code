def solve(bingo, winners, drawn):
    for board in bingo:
        for i in range(len(board)):
            if board[i] in drawn:
                board[i] = 0

    for board in bingo:
        for i in range(5):
            if sum(board[i * 5 : i * 5 + 5]) == 0:
                winners.append(board)
                bingo.remove(board)
                break
            if (
                board[i]
                + board[i + 5]
                + board[i + 10]
                + board[i + 15]
                + board[i + 20]
                == 0
            ):
                winners.append(board)
                bingo.remove(board)
                break

    return winners, bingo


def main():
    nums = [int(x) for x in open("input.txt", "r").read().split(",")]

    data = [line.strip() for line in open("input2.txt", "r").readlines()]
    bingo, i = [], 0

    data.append("")

    while data:
        bingo.append([])
        for _ in range(6):
            for val in data.pop(0).split():
                bingo[i].append(int(val))
        i += 1

    # ------------------------------------------------
    # Part1

    winners, drawn = [], []
    while not winners:
        drawn.append(nums.pop(0))
        winners, bingo = solve(bingo, winners, drawn)

    print(f"Part #1: {drawn[-1] * sum(winners[0])}")

    # -------------------------------------------------
    # Part 2

    winners, drawn = [], []
    while nums and bingo:
        drawn.append(nums.pop(0))
        winners, bingo = solve(bingo, winners, drawn)

    print(f"Part #2: {drawn[-1] * sum(winners[-1])}")


main()
