const std = @import("std");

pub fn isPangram(str: []const u8) bool {
    var counts: [26]usize = .{0} ** 26;

    for (str) |c| {
        if (!std.ascii.isAlphabetic(c)) {
            continue;
        }
        const lower = std.ascii.toLower(c);
        const index = lower - 'a';
        counts[index] += 1;
    }

    for (counts) |c| {
        if (c == 0) {
            return false;
        }
    }
    return true;
}
