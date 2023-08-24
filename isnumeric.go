package asciiArtColor

// to check if it was a number or not.
// #Q.5,pt.11.
func IsNumeric(s string) bool {
	for _, w := range s {
		if isLt(w) == false {
			return false
		}
	}
	return true
}

// one by one.
func isLt(woo rune) bool {
	if woo == ',' {
		return false
	}
	return (woo >= '0' && woo <= '9')
}
