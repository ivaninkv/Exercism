const std = @import("std");

/// Writes a reversed copy of `s` to `buffer`.
pub fn reverse(buffer: []u8, s: []const u8) []u8 {
    for (s, 0..) |str, i| {
        buffer[s.len - i - 1] = str;
    }

    return buffer[0..s.len];
}

pub fn main() void {
    var buffer: [80]u8 = undefined;
    const input = "Hello, World!";
    const reversed = reverse(&buffer, input);
    std.debug.print("Original: {s}\nReversed: {s}\n", .{ input, reversed });
}
