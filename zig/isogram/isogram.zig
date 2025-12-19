const std = @import("std");

pub fn isIsogram(str: []const u8) bool {
    var counts: [26]usize = .{0} ** 26;

    for (str) |c| {
        if (!std.ascii.isAlphabetic(c)) {
            continue;
        }

        const lower = std.ascii.toLower(c);
        const index = lower - 'a';
        if (counts[index] > 0) {
            return false;
        }
        counts[index] += 1;
    }

    return true;
}

pub fn main() void {
    std.debug.print("Is isogram {}\n", .{isIsogram("Alphabet")});
}
