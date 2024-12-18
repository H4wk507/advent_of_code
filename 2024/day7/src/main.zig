const std = @import("std");

fn concat(x: u64, y: u64) u64 {
    const y_digits = @as(u64, @intFromFloat(@floor(@log10(@as(f64, @floatFromInt(y)))) + 1));
    return x * std.math.pow(u64, 10, y_digits) + y;
}

fn solve(test_val: u64, acc: u64, nums: std.ArrayList(u64), idx: usize, p2: bool) bool {
    if (acc > test_val) return false;

    if (idx >= nums.items.len) {
        return acc == test_val;
    }

    const add_fold = acc + nums.items[idx];
    const mul_fold = acc * nums.items[idx];

    const add = solve(test_val, add_fold, nums, idx + 1, p2);
    const mul = solve(test_val, mul_fold, nums, idx + 1, p2);
    const con = if (p2) solve(test_val, concat(acc, nums.items[idx]), nums, idx + 1, p2) else false;

    return add or mul or con;
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    const file_content = try std.fs.cwd().readFileAlloc(allocator, "dat", 1024 * 1024);

    var cnt1: u64 = 0;
    var cnt2: u64 = 0;
    var lines_it = std.mem.splitSequence(u8, file_content, "\n");
    while (lines_it.next()) |line| {
        if (line.len == 0) continue;
        var color_split = std.mem.splitSequence(u8, line, ":");
        const t = color_split.next() orelse unreachable;
        const r = color_split.next() orelse unreachable;

        const test_val = try std.fmt.parseInt(u64, t, 10);

        var nums = std.ArrayList(u64).init(allocator);
        var rems_it = std.mem.splitSequence(u8, r, " ");
        while (rems_it.next()) |rem| {
            if (rem.len == 0) continue;
            const rem_val = try std.fmt.parseInt(u64, rem, 10);
            try nums.append(rem_val);
        }
        cnt1 += if (solve(test_val, nums.items[0], nums, 1, false)) test_val else 0;
        cnt2 += if (solve(test_val, nums.items[0], nums, 1, true)) test_val else 0;
    }
    std.debug.print("{}\n", .{cnt1});
    std.debug.print("{}\n", .{cnt2});
}
