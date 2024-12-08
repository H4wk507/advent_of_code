const std = @import("std");

fn p1(grid: [][]u8, i: usize, j: usize, dx: i32, dy: i32) usize {
    const n: i32 = @intCast(grid.len);
    const m: i32 = @intCast(grid[0].len);

    const word = "MAS";
    var y: i32 = @intCast(i);
    var x: i32 = @intCast(j);

    for (word) |target| {
        x += dx;
        y += dy;
        if (x < 0 or y < 0 or y >= n or x >= m) {
            break;
        }
        if (grid[@intCast(y)][@intCast(x)] != target) {
            break;
        }
        if (target == word[word.len - 1]) return 1;
    }
    return 0;
}

fn p2(grid: [][]u8, i: usize, j: usize) usize {
    const y: i32 = @intCast(i);
    const x: i32 = @intCast(j);
    const n: i32 = @intCast(grid.len);
    const m: i32 = @intCast(grid[0].len);

    if (!(y - 1 >= 0 and x - 1 >= 0 and y + 1 < n and x + 1 < m)) {
        return 0;
    }

    const tl = grid[@intCast(y - 1)][@intCast(x - 1)];
    const br = grid[@intCast(y + 1)][@intCast(x + 1)];
    const tr = grid[@intCast(y - 1)][@intCast(x + 1)];
    const bl = grid[@intCast(y + 1)][@intCast(x - 1)];

    const diag1 = (tl == 'M' and br == 'S') or (tl == 'S' and br == 'M');
    const diag2 = (tr == 'M' and bl == 'S') or (tr == 'S' and bl == 'M');

    return if (diag1 and diag2) 1 else 0;
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();
    const file_content = try std.fs.cwd().readFileAlloc(allocator, "dat", 1024 * 1024);
    var lines_iterator = std.mem.splitSequence(u8, file_content, "\n");

    const first_line = lines_iterator.next() orelse return error.EmptyFile;
    const M = first_line.len;

    var N: usize = 1;
    while (lines_iterator.next()) |_| {
        N += 1;
    }

    const grid = try allocator.alloc([]u8, N);
    for (grid) |*row| {
        row.* = try allocator.alloc(u8, M);
    }

    lines_iterator = std.mem.splitSequence(u8, file_content, "\n");
    var k: usize = 0;
    while (lines_iterator.next()) |line| {
        for (line, 0..) |char, j| {
            grid[k][j] = char;
        }
        k += 1;
    }

    var cnt: usize = 0;
    var cnt2: usize = 0;
    for (grid, 0..N) |row, i| {
        for (row, 0..M) |cell, j| {
            if (cell == 'X') {
                cnt += p1(grid, i, j, -1, 0);
                cnt += p1(grid, i, j, 1, 0);
                cnt += p1(grid, i, j, 0, -1);
                cnt += p1(grid, i, j, 0, 1);
                cnt += p1(grid, i, j, -1, -1);
                cnt += p1(grid, i, j, -1, 1);
                cnt += p1(grid, i, j, 1, -1);
                cnt += p1(grid, i, j, 1, 1);
            }
            if (cell == 'A') {
                cnt2 += p2(grid, i, j);
            }
        }
    }
    std.debug.print("{}\n", .{cnt});
    std.debug.print("{}\n", .{cnt2});
}
