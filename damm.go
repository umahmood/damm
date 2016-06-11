package damm

// opTable the operation table it may be obtained from the totally
// anti-symmetric quasigroup in Damm's doctoral dissertation page 111.
// http://archiv.ub.uni-marburg.de/diss/z2004/0516/pdf/dhmd.pdf
func opTable() [10][10]byte {
	var table [10][10]byte
	table[0] = [10]byte{0, 3, 1, 7, 5, 9, 8, 6, 4, 2}
	table[1] = [10]byte{7, 0, 9, 2, 1, 5, 4, 8, 6, 3}
	table[2] = [10]byte{4, 2, 0, 6, 8, 7, 1, 3, 5, 9}
	table[3] = [10]byte{1, 7, 5, 0, 9, 8, 3, 4, 2, 6}
	table[4] = [10]byte{6, 1, 2, 3, 0, 4, 5, 9, 7, 8}
	table[5] = [10]byte{3, 6, 7, 4, 2, 0, 9, 5, 8, 1}
	table[6] = [10]byte{5, 8, 6, 9, 7, 2, 0, 1, 3, 4}
	table[7] = [10]byte{8, 9, 4, 5, 3, 6, 2, 0, 1, 7}
	table[8] = [10]byte{9, 4, 3, 8, 6, 1, 7, 2, 0, 5}
	table[9] = [10]byte{2, 5, 8, 1, 4, 3, 6, 7, 9, 0}
	return table
}

// checkDigits calculates the check digit of n.
func checkDigit(n int) int {
	var digits []byte
	for n > 0 {
		digits = append(digits, byte(n%10))
		n = n / 10
	}
	table := opTable()
	l := len(digits) - 1
	var interim byte
	for l >= 0 {
		interim = table[interim][digits[l]]
		l = l - 1
	}
	return int(interim)
}

// Calc calculates the check digit of n and appends it as a trailing digit to n.
func Calc(n int) int {
	c := checkDigit(n)
	n = (n << 3) + (n << 1) + c
	return n
}

// Validate n against the included trailing check digit.
func Validate(n int) bool {
	if checkDigit(n) != 0 {
		return false
	}
	return true
}
