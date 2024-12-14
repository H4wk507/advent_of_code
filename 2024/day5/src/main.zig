const std = @import("std");

inline fn array_contains(comptime T: type, haystack: []T, needle: T) bool {
    return for (haystack) |element| {
        if (element == needle) break true;
    } else false;
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    const file_content = try std.fs.cwd().readFileAlloc(allocator, "dat", 1024 * 1024);
    var lines_it = std.mem.splitSequence(u8, file_content, "\n");

    var map = std.AutoHashMap(u64, std.ArrayList(u64)).init(allocator);
    defer {
        var it = map.valueIterator();
        while (it.next()) |list| {
            list.deinit();
        }
        map.deinit();
    }
    while (lines_it.next()) |line| {
        if (line.len == 0) break;

        var pages = std.mem.splitSequence(u8, line, "|");
        const before = try std.fmt.parseInt(u64, pages.next() orelse return error.InvalidFormat, 10);
        const y = try std.fmt.parseInt(u64, pages.next() orelse return error.InvalidFormat, 10);

        var entry = try map.getOrPut(y);
        if (!entry.found_existing) {
            entry.value_ptr.* = std.ArrayList(u64).init(allocator);
        }
        try entry.value_ptr.append(before);
    }

    var sum1: u64 = 0;
    var sum2: u64 = 0;
    while (lines_it.next()) |line| {
        if (line.len == 0) break;

        var nums = std.ArrayList(u64).init(allocator);
        defer nums.deinit();
        var nums_it = std.mem.splitSequence(u8, line, ",");

        while (nums_it.next()) |n_string| {
            if (n_string.len == 0) break;
            try nums.append(try std.fmt.parseInt(u64, n_string, 10));
        }

        if (nums.items.len == 0) continue;

        var ok = true;
        var consumed = std.ArrayList(u64).init(allocator);
        defer consumed.deinit();

        for (nums.items) |num| {
            if (map.get(num)) |befores| {
                for (befores.items) |before| {
                    if (!array_contains(u64, nums.items, before)) continue;
                    if (!array_contains(u64, consumed.items, before)) {
                        ok = false;
                        break;
                    }
                }
                if (!ok) break;
            }
            try consumed.append(num);
        }

        const middle_idx = nums.items.len / 2;
        if (ok) {
            sum1 += nums.items[middle_idx];
        } else {
            var valid_arr = try allocator.alloc(u64, nums.items.len);
            defer allocator.free(valid_arr);
            @memset(valid_arr, 0);

            for (nums.items) |num| {
                var idx: usize = 0;
                if (map.get(num)) |befores| {
                    for (befores.items) |before| {
                        if (array_contains(u64, nums.items, before)) {
                            idx += 1;
                        }
                    }
                }
                valid_arr[idx] = num;
            }
            sum2 += valid_arr[middle_idx];
        }
    }
    std.debug.print("sum1: {}\n", .{sum1});
    std.debug.print("sum2: {}\n", .{sum2});
}
