package resistorcolorduo

const (
	black = iota
	brown
	red
	orange
	yellow
	green
	blue
	violet
	grey
	white
)

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	if len(colors) < 2 {
		return 0
	}
	return colorCode(colors[0])*10 + colorCode(colors[1])
}

func colorCode(color string) int {
	switch color {
	case "black":
		return black
	case "brown":
		return brown
	case "red":
		return red
	case "orange":
		return orange
	case "yellow":
		return yellow
	case "green":
		return green
	case "blue":
		return blue
	case "violet":
		return violet
	case "grey":
		return grey
	case "white":
		return white
	default:
		return -1
	}
}
