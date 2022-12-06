def extract_numbers(s: str) -> list[int]:
    """Extract numbers from s if s has a form of `move 6 from 2 to 1`."""
    return [int(x) for i, x in enumerate(s.split()) if i % 2 != 0]


def get_top_idx(arr: list[str]) -> int:
    """Get index of the top element"""
    for i, c in enumerate(arr):
        if not c.isspace():
            return i
    return 0


def get_last_free_space(arr: list[str]) -> int:
    """Get index of the last free (' ') space"""
    return get_top_idx(arr) - 1


def solve(boxes: str, moves: str, part: int) -> str:
    moves = moves.split("\n")
    boxes = boxes.split("\n")
    transposed_boxes: list[list[str]] = list(map(list, zip(*boxes)))
    # remove unnecessary (not letter) parts of boxes
    filtered_boxes: list[list[str]] = list(
        filter(lambda lst: any([c.isalpha() for c in lst]), transposed_boxes)
    )
    for move in moves:
        amount, src, dst = extract_numbers(move)
        src -= 1
        dst -= 1
        for _ in range(amount):
            top_idx = get_top_idx(filtered_boxes[src])
            # save top_element
            top_element = filtered_boxes[src][top_idx]
            # remove it from the top
            filtered_boxes[src][top_idx] = " "
            dst_idx = get_last_free_space(filtered_boxes[dst])
            # if no space left in dst, push it to the front
            if dst_idx == -1:
                filtered_boxes[dst].insert(0, top_element)
            else:
                filtered_boxes[dst][dst_idx] = top_element
        # for part 2 just reverse inserted boxes in that turn
        if part == 2:
            start = get_top_idx(filtered_boxes[dst])
            end = start + amount
            filtered_boxes[dst][start:end] = filtered_boxes[dst][start:end][
                ::-1
            ]

    return "".join([box[get_top_idx(box)] for box in filtered_boxes])


if __name__ == "__main__":
    file = "input.txt"
    with open(file) as f:
        boxes, moves = f.read().strip("\n").split("\n\n")
    print(f"Part #1: {solve(boxes, moves, 1)}")
    print(f"Part #2: {solve(boxes, moves, 2)}")
