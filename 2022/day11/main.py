import re
from math import prod
from copy import deepcopy


def extract_numbers(s: str) -> list[int]:
    """Extract numbers from string"""
    return [int(x) for x in re.findall(r"\b\d+\b", s)]


def get_operation(s: str) -> str:
    """Operation: new = old + 6"""
    idx = s.find("=")
    return s[idx + 1 :].strip()


def perform_operation(op: str, item: int, div: int) -> int:
    op = op.split()[1:]
    if op[0] == "+":
        return (
            item % div + (item % div if op[1] == "old" else int(op[1]) % div)
        ) % div
    elif op[0] == "*":
        return (
            item % div * (item % div if op[1] == "old" else int(op[1]) % div)
        ) % div


class Monke:
    def __init__(
        self,
        items: list[int],
        operation: str,
        divisor: int,
        passed_idx: int,
        failed_idx: int,
    ):
        self.items = items
        self.operation = operation
        self.divisor = divisor
        self.passed_idx = passed_idx
        self.failed_idx = failed_idx
        self.inspected = 0


def populate_monkes(monkes: list[str], monkes_list: list[Monke]) -> None:
    for monke in monkes:
        monke = [x.strip() for x in monke.split("\n")]
        items = extract_numbers(monke[1])
        operation = get_operation(monke[2])
        divisor = extract_numbers(monke[3])[0]
        passed_idx = extract_numbers(monke[4])[0]
        failed_idx = extract_numbers(monke[5])[0]
        monkes_list.append(
            Monke(items, operation, divisor, passed_idx, failed_idx)
        )


def solve(monkes_list: list[Monke], part: int) -> int:
    # we have to find modulo which is divisible by every divisor
    # obvious answer is a product of divisors but optimal is lcm of divisors.
    # lcm and product if every divisor is prime are the same, so don't bother.
    divs_prod = prod([monke.divisor for monke in monkes_list])
    NROUNDS = 20 if part == 1 else 10_000
    for _ in range(NROUNDS):
        for monke in monkes_list:
            while len(monke.items) > 0:
                item = monke.items.pop(0)
                item = perform_operation(monke.operation, item, divs_prod)
                if part == 1:
                    item = item // 3
                monke.inspected += 1
                if item % monke.divisor == 0:
                    monkes_list[monke.passed_idx].items.append(item)
                else:
                    monkes_list[monke.failed_idx].items.append(item)
    monke_business = prod(
        sorted([monke.inspected for monke in monkes_list], reverse=True)[:2]
    )
    return monke_business


if __name__ == "__main__":
    with open("input.txt") as f:
        monkes = f.read().strip("\n").split("\n\n")

    monkes_list: list[Monke] = []
    populate_monkes(monkes, monkes_list)
    print(f"Part #1: {solve(deepcopy(monkes_list), 1)}")
    print(f"Part #2: {solve(deepcopy(monkes_list), 2)}")
