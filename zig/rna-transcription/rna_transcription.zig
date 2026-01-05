const std = @import("std");
const mem = std.mem;

pub fn toRna(allocator: mem.Allocator, dna: []const u8) mem.Allocator.Error![]const u8 {
    var rna = try allocator.alloc(u8, dna.len);

    for (dna, 0..) |nucleotide, i| {
        switch (nucleotide) {
            'G' => rna[i] = 'C',
            'C' => rna[i] = 'G',
            'T' => rna[i] = 'A',
            'A' => rna[i] = 'U',
            else => {
                allocator.free(rna);
                return error.OutOfMemory;
            },
        }
    }

    return rna;
}
