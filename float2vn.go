package num2vn

import "strconv"

func Float2Vn(number float64) string {
	if number == 0 {
		return "Không"
	}

	intNum := int64(number)
	if number == float64(intNum) {
		// Not a fraction number
		return Int2Vn(intNum)
	}

	var fractNum int64
	var vn, fractStr string

	// Convert float -0.00456 to string "-0.00456"
	fractStr = strconv.FormatFloat(number, 'f', -1, 64)

	r := []rune(fractStr)
	dotHit := false

	// Search the decimal dot .
	// Convert all '0' after dot to "không"
	// Convert the fraction digits to words
	for i, rLen := 0, len(r); i < rLen; i++ {

		switch r[i] {
		case '.':
			dotHit = true
		case '0':
			if dotHit {
				vn = "không " + vn
			}
		default:
			if dotHit {
				fractNum, _ = strconv.ParseInt(fractStr[i:], 10, 64)
				vn += int2VnStr(fractNum)
				i = rLen
			}
		}
	}

	if intNum == 0 {
		if number >= 0 {
			return "Không phẩy " + vn
		} else {
			return "Âm không phẩy " + vn
		}
	}

	return Int2Vn(intNum) + " phẩy " + vn
}
