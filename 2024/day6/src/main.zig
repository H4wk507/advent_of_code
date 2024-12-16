const std = @import("std");

fn run(grid: *[][]u8, startx: usize, starty: usize, N: usize, M: usize) usize {
    var pos_x: i64 = @as(i64, @intCast(startx));
    var pos_y: i64 = @as(i64, @intCast(starty));
    var x: i64 = @as(i64, @intCast(startx));
    var y: i64 = @as(i64, @intCast(starty));
    var dx: i64 = 0;
    var dy: i64 = -1;
    var steps: usize = 0;
    grid.*[@as(usize, @intCast(pos_y))][@as(usize, @intCast(pos_x))] = 'X';
    while (x >= 0 and x < M and y >= 0 and y < N and steps < N * M) {
        if (grid.*[@as(usize, @intCast(y))][@as(usize, @intCast(x))] == '#') {
            const tmp = dx;
            dx = -dy;
            dy = tmp;
            x = pos_x;
            y = pos_y;
        } else {
            pos_x = x;
            pos_y = y;
            grid.*[@as(usize, @intCast(y))][@as(usize, @intCast(x))] = 'X';
            steps += 1;
        }
        x += dx;
        y += dy;
    }
    return if (steps >= N * M) 1 else 0;
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const a = arena.allocator();
    const file_content = try std.fs.cwd().readFileAlloc(a, "dat", 1024 * 1024);
    var lines_iterator = std.mem.splitSequence(u8, file_content, "\n");
    const first_line = lines_iterator.next() orelse return error.EmptyFile;
    const M = first_line.len;
    var N: usize = 1;
    while (lines_iterator.next()) |line| {
        if (line.len > 0) N += 1;
    }
    var grid = try a.alloc([]u8, N);
    lines_iterator = std.mem.splitSequence(u8, file_content, "\n");
    var pos_x: usize = 0;
    var pos_y: usize = 0;
    var row_idx: usize = 0;
    while (lines_iterator.next()) |line| {
        if (line.len == 0) continue;
        grid[row_idx] = try a.alloc(u8, M);
        @memcpy(grid[row_idx][0..], line[0..]);
        for (line, 0..) |cell, j| {
            if (cell == '^') {
                pos_x = j;
                pos_y = row_idx;
            }
        }
        row_idx += 1;
    }

    var path_grid = try a.alloc([]u8, N);
    for (path_grid) |*row| row.* = try a.alloc(u8, M);
    @memcpy(path_grid[0..], grid[0..]);
    _ = run(&path_grid, pos_x, pos_y, N, M);
    // p1
    var cnt: usize = 0;
    for (path_grid) |row| {
        for (row) |cell| {
            if (cell == 'X') {
                cnt += 1;
            }
        }
    }
    // p2
    var cnt2: usize = 0;
    var test_grid = try a.alloc([]u8, N);
    for (test_grid) |*row| row.* = try a.alloc(u8, M);
    for (0.., path_grid) |i, row| {
        for (0.., row) |j, cell| {
            if (cell != 'X' or (i == pos_y and j == pos_x)) continue;
            for (0..N) |y| @memcpy(test_grid[y][0..], grid[y][0..]);
            test_grid[i][j] = '#';
            cnt2 += run(&test_grid, pos_x, pos_y, N, M);
        }
    }

    std.debug.print("p1: {}\n", .{cnt});
    std.debug.print("p2: {}\n", .{cnt2});
}
