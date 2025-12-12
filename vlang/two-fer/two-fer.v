module main

fn two_fer(name string) string {
	n := match name {
		'' { 'you' }
		else { name }
	}
	return 'One for ${n}, one for me.'
}
