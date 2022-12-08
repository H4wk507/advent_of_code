from typing import Optional


class Directory:
    def __init__(self, name: str, parent: Optional["Directory"]):
        self.name = name
        self.parent = parent
        self.files_size: int = 0
        self.subdirs: dict[str, Directory] = {}

    def size(self) -> int:
        """Recursively get directory size."""
        return sum(d.size() for d in self.subdirs.values()) + self.files_size

    def get_directories_sizes(
        self, path: str = "", sizes: dict[str, int] = {}
    ) -> dict[str, int]:
        """Get sizes of all directiories from self."""
        sizes[path + self.name] = self.size()
        for subdir in self.subdirs.values():
            subdir.get_directories_sizes(path + self.name, sizes)
        return sizes


def solve(lines: list[str], part: int) -> int:
    pwd = root = Directory("/", None)
    for line in lines:
        match line.split():
            case ["$", "cd", "/"]:
                pwd = root
            case ["$", "cd", ".."]:
                pwd = pwd.parent
            case ["$", "cd", folder]:
                if folder not in pwd.subdirs:
                    pwd.subdirs[folder] = Directory(folder, pwd)
                pwd = pwd.subdirs[folder]
            case ["dir", folder]:
                pwd.subdirs[folder] = Directory(folder, pwd)
            case [filesize, filename]:
                if filesize.isdecimal():
                    pwd.files_size += int(filesize)

    sizes = root.get_directories_sizes().values()
    if part == 1:
        return sum(size for size in sizes if size <= 100000)
    elif part == 2:
        disk_space = 70_000_000
        needed = 30_000_000
        used = root.size()
        to_free = used + needed - disk_space
        return min(size for size in sizes if to_free <= size < used)


if __name__ == "__main__":
    file = "input.txt"
    with open(file) as f:
        lines = f.read().strip("\n").split("\n")
    print(f"Part #1: {solve(lines, 1)}")
    print(f"Part #2: {solve(lines, 2)}")
