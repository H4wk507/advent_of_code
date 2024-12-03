const std = @import("std");

const Lexer = struct {
    content: []u8,
    content_len: usize,
    cursor: usize,
};

fn parse_dont(lexer: *Lexer, do: *bool) void {
    const cur = lexer.cursor;
    const dont = "don't()";
    var j: usize = 0;
    for (dont) |c| {
        if (lexer.content[lexer.cursor] == c) {
            lexer.cursor += 1;
            j += 1;
        } else {
            break;
        }
    }
    if (j == dont.len) {
        lexer.cursor -= 1;
        do.* = false;
    } else {
        lexer.cursor = cur;
    }
}

fn parse_do(lexer: *Lexer, do: *bool) void {
    const cur = lexer.cursor;
    const do_keyword = "do()";
    var j: usize = 0;
    for (do_keyword) |c| {
        if (lexer.content[lexer.cursor] == c) {
            lexer.cursor += 1;
            j += 1;
        } else {
            break;
        }
    }
    if (j == do_keyword.len) {
        do.* = true;
        lexer.cursor -= 1;
    } else {
        lexer.cursor = cur;
    }
}

fn parse_number(lexer: *Lexer) u32 {
    var n: u32 = 0;
    while (lexer.cursor < lexer.content_len and std.ascii.isDigit(lexer.content[lexer.cursor])) {
        n = (n * 10) + (@as(u32, lexer.content[lexer.cursor] - '0'));
        lexer.cursor += 1;
    }
    return n;
}

fn parse_mul(lexer: *Lexer, do: bool) u32 {
    const cur = lexer.cursor;
    const mul = "mul";
    var j: usize = 0;
    for (mul) |m| {
        if (lexer.content[lexer.cursor] == m) {
            lexer.cursor += 1;
            j += 1;
        } else {
            break;
        }
    }
    if (j != mul.len) {
        lexer.cursor = cur;
        return 0;
    }

    if (lexer.content[lexer.cursor] != '(') {
        lexer.cursor = cur;
        return 0;
    }
    lexer.cursor += 1;

    if (!std.ascii.isDigit(lexer.content[lexer.cursor])) {
        lexer.cursor = cur;
        return 0;
    }

    const number = parse_number(lexer);

    if (lexer.content[lexer.cursor] != ',') {
        lexer.cursor = cur;
        return 0;
    }
    lexer.cursor += 1;

    if (!std.ascii.isDigit(lexer.content[lexer.cursor])) {
        lexer.cursor = cur;
        return 0;
    }
    const number2 = parse_number(lexer);

    if (lexer.content[lexer.cursor] != ')') {
        lexer.cursor = cur;
        return 0;
    }
    if (do) {
        return number * number2;
    }
    return 0;
}

fn p1(lexer: *Lexer) void {
    var res: u32 = 0;
    while (lexer.cursor < lexer.content_len) {
        res += parse_mul(lexer, true);
        lexer.cursor += 1;
    }
    std.debug.print("p1: {d}\n", .{res});
}

fn p2(lexer: *Lexer) void {
    var res: u32 = 0;
    var do = true;
    while (lexer.cursor < lexer.content_len) {
        res += parse_mul(lexer, do);
        parse_do(lexer, &do);
        parse_dont(lexer, &do);
        lexer.cursor += 1;
    }
    std.debug.print("p2: {d}\n", .{res});
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    const file_content = try std.fs.cwd().readFileAlloc(allocator, "dat", 1024 * 1024);

    var lexer1 = Lexer{ .content = file_content, .content_len = file_content.len, .cursor = 0 };
    var lexer2 = Lexer{ .content = file_content, .content_len = file_content.len, .cursor = 0 };
    p1(&lexer1);
    p2(&lexer2);
}
