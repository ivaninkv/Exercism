module main

fn grains_on_square(square int) !u64 {
	if square < 1 || square > 64 {
		return error('invalid square number should return error')
	}
	mut result := u64(1)
	for i := 1; i < square; i++ {
		result *= 2
	}

	return result
}

fn total_grains_on_board() u64 {
	mut res := u64(0)

	for i := 1; i <= 64; i++ {
		res += grains_on_square(i) or { return 0 }
	}

	return res
}

fn main() {
	println(total_grains_on_board())
}
