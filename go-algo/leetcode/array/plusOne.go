package array

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			break
		}
	}
	if digits[0] == 0 {
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}
