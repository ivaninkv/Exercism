module main

fn count_nucleotides(strand string) !map[string]int {
	for r in strand.runes() {
		if !'ACGT'.contains(r.str()) {
			return error('${r.str()} is not a valid nucleotide!')
		}
	}

	return {
		'A': strand.count('A')
		'C': strand.count('C')
		'G': strand.count('G')
		'T': strand.count('T')
	}
}

fn main() {
	count_nucleotides('AGXXACT') or { println(err) }
}
