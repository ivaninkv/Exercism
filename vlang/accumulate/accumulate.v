module main

fn accumulate_ints(values []int, operation fn (int) int) []int {
	mut res := []int{}

	for val in values {
		res << operation(val)
	}

	return res
}

// Because V functions cannot be overloaded[1], make another function
//  called `accumulate_strs` that does the same thing for strings
// instead of ints

fn accumulate_strs(values []string, operation fn (string) string) []string {
	mut res := []string{}

	for val in values {
		res << operation(val)
	}

	return res
}
