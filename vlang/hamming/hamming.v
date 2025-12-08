module main

fn distance(a string, b string) !int {
	if a.len_utf8() != b.len_utf8() {
		return error('lengths must match!')
	}
	mut distance := 0
	for i in 0 .. a.len_utf8() {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance
}

fn main() {
	println(distance('ABBA', 'BABA') or { return })
}
