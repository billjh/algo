package misc

// GCD calculates the greatest common divisor of two input numbers
// when any input number is zero, returns zero
func GCD(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
