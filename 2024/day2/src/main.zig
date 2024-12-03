const std = @import("std");

fn is_safe(ll: *std.ArrayList(u8), maybe_skip_idx: ?usize) bool {
    var order: i32 = 0;
    var ok = true;

    var temp = std.ArrayList(u8).init(ll.allocator);
    defer temp.deinit();

    for (ll.items, 0..) |item, i| {
        if (maybe_skip_idx) |skip_idx| {
            if (i != skip_idx) {
                temp.append(item) catch unreachable;
            }
        } else {
            temp.append(item) catch unreachable;
        }
    }

    for (0..temp.items.len - 1) |i| {
        const diff = @as(i32, temp.items[i + 1]) - @as(i32, temp.items[i]);

        if (diff > 0) {
            if (order == -1) {
                ok = false;
                break;
            }
            order = 1;
        }
        if (diff < 0) {
            if (order == 1) {
                ok = false;
                break;
            }
            order = -1;
        }
        if (@abs(diff) < 1 or @abs(diff) > 3) {
            ok = false;
            break;
        }
    }
    return ok;
}

fn p(rr: *std.mem.SplitIterator(u8, .sequence), al: std.mem.Allocator) !void {
    var cnt1: u32 = 0;
    var cnt2: u32 = 0;
    while (rr.next()) |r| {
        if (r.len == 0) {
            continue;
        }
        var split = std.mem.splitSequence(u8, r, " ");
        var ll = std.ArrayList(u8).init(al);
        while (split.next()) |s| {
            const n = std.fmt.parseInt(u8, s, 10) catch unreachable;
            try ll.append(n);
        }
        var ok = false;
        if (is_safe(&ll, null)) cnt1 += 1;
        if (is_safe(&ll, 0)) ok = true;

        for (0..ll.items.len - 1) |i| {
            const diff = @as(i32, ll.items[i + 1]) - @as(i32, ll.items[i]);
            if (@abs(diff) < 1 or @abs(diff) > 3) {
                if (is_safe(&ll, i) or is_safe(&ll, i + 1)) {
                    ok = true;
                }
                break;
            }
            if (i + 2 < ll.items.len) {
                const diff2 = @as(i32, ll.items[i + 2]) - @as(i32, ll.items[i + 1]);
                if ((diff > 0) != (diff2 > 0)) { // sign compare
                    if (is_safe(&ll, i) or is_safe(&ll, i + 1) or is_safe(&ll, i + 2)) {
                        ok = true;
                    }
                    break;
                }
            }
        }
        if (ok) {
            cnt2 += 1;
        }
    }
    std.debug.print("{d} {d}\n", .{ cnt1, cnt2 });
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    const file_content = try std.fs.cwd().readFileAlloc(allocator, "dat", 1024 * 1024);

    var reports = std.mem.splitSequence(u8, file_content, "\n");
    try p(&reports, allocator);
}
