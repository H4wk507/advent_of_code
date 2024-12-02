const std = @import("std");

fn p1(ll: []u64, rr: []u64) u64 {
    var s: u64 = 0;
    for (ll, rr) |l, r| {
        s += if (r > l) r - l else l - r;
    }
    return s;
}

fn p2(ll: []u64, rr: []u64, allocator: std.mem.Allocator) !u64 {
    var map = std.AutoHashMap(u64, u64).init(
        allocator,
    );

    for (rr) |r| {
        if (map.get(r)) |v| {
            try map.put(r, v + 1);
        } else {
            try map.put(r, 1);
        }
    }

    var s: u64 = 0;
    for (ll) |l| {
        const v = map.get(l) orelse 0;
        s += l * v;
    }
    return s;
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    const file_content = try std.fs.cwd().readFileAlloc(allocator, "dat1", 1024 * 1024);

    var ll = std.ArrayList(u64).init(allocator);
    var rr = std.ArrayList(u64).init(allocator);
    var rows = std.mem.splitSequence(u8, file_content, "\n");
    while (rows.next()) |row| {
        if (row.len == 0) {
            continue;
        }
        var split = std.mem.splitSequence(u8, row, "   ");
        const l = try std.fmt.parseInt(u64, split.next() orelse "impossible", 10);
        const r = try std.fmt.parseInt(u64, split.next() orelse "impossible", 10);
        try ll.append(l);
        try rr.append(r);
    }

    std.mem.sort(u64, ll.items, {}, std.sort.asc(u64));
    std.mem.sort(u64, rr.items, {}, std.sort.asc(u64));

    std.debug.print("{d}\n", .{p1(ll.items, rr.items)});
    std.debug.print("{d}\n", .{try p2(ll.items, rr.items, allocator)});
}
