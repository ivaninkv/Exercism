const std = @import("std");

pub fn isLeapYear(year: u32) bool {
    return (year % 4 == 0 and year % 100 != 0) or year % 400 == 0;
}

pub fn main() !void {
    std.debug.print("Year 2015 is leap: {}\n", .{isLeapYear(2015)});
    std.debug.print("Year 2024 is leap: {}\n", .{isLeapYear(2024)});
}
