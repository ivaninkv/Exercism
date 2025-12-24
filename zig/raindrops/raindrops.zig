const std = @import("std");

pub fn convert(buffer: []u8, n: u32) []const u8 {
    var stream = std.io.fixedBufferStream(buffer);
    const writer = stream.writer();
    var wrote = false;

    if (n % 3 == 0) {
        writer.writeAll("Pling") catch unreachable;
        wrote = true;
    }
    if (n % 5 == 0) {
        writer.writeAll("Plang") catch unreachable;
        wrote = true;
    }
    if (n % 7 == 0) {
        writer.writeAll("Plong") catch unreachable;
        wrote = true;
    }

    if (!wrote) {
        writer.print("{d}", .{n}) catch unreachable;
    }

    return stream.getWritten();
}
