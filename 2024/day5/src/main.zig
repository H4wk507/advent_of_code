const std = @import("std");

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    const file_content = try std.fs.cwd().readFileAlloc(allocator, "dat", 1024 * 1024);

    var sections = std.mem.splitSequence(u8, file_content, "\n\n");
    const rules_section = sections.next() orelse return error.InvalidFormat;
    const updates_section = sections.next() orelse return error.InvalidFormat;

    var map = std.AutoHashMap(u64, u128).init(allocator);
    defer map.deinit();

    var rules_lines = std.mem.splitSequence(u8, rules_section, "\n");
    while (rules_lines.next()) |line| {
        if (line.len == 0) break;

        var pages = std.mem.splitSequence(u8, line, "|");
        const before = try std.fmt.parseInt(u7, pages.next() orelse return error.InvalidFormat, 10);
        const y = try std.fmt.parseInt(u64, pages.next() orelse return error.InvalidFormat, 10);

        const entry = try map.getOrPut(y);
        if (!entry.found_existing) {
            entry.value_ptr.* = 0;
        }
        entry.value_ptr.* |= (@as(u128, 1) << before);
    }

    var sum1: u64 = 0;
    var sum2: u64 = 0;
    var updates = std.mem.splitSequence(u8, updates_section, "\n");
    while (updates.next()) |line| {
        if (line.len == 0) break;

        var nums_buf: [128]u64 = undefined;
        var nums_count: usize = 0;

        var nums_it = std.mem.splitSequence(u8, line, ",");
        while (nums_it.next()) |n_string| {
            if (n_string.len == 0) break;
            nums_buf[nums_count] = try std.fmt.parseInt(u64, n_string, 10);
            nums_count += 1;
        }

        if (nums_count == 0) continue;

        var current_set: u128 = 0;
        for (nums_buf[0..nums_count]) |num| {
            current_set |= (@as(u128, 1) << @as(u7, @intCast(num)));
        }

        var ok = true;
        var consumed: u128 = 0;

        for (nums_buf[0..nums_count]) |num| {
            if (map.get(num)) |before_set| {
                if ((before_set & current_set) != 0 and (before_set & current_set) != (before_set & consumed)) {
                    ok = false;
                    break;
                }
            }
            consumed |= @as(u128, 1) << @as(u7, @intCast(num));
        }

        const middle_idx = nums_count / 2;
        if (ok) {
            sum1 += nums_buf[middle_idx];
        } else {
            var valid: [128]u64 = undefined;

            for (nums_buf[0..nums_count]) |num| {
                var pos: usize = 0;
                if (map.get(num)) |before_set| {
                    pos = @popCount(before_set & current_set);
                }
                valid[pos] = num;
            }
            sum2 += valid[middle_idx];
        }
    }
    std.debug.print("sum1: {}\n", .{sum1});
    std.debug.print("sum2: {}\n", .{sum2});
}
