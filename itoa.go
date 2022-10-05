package exam

func Itoa(n int) string {
	s := ""
	var arrayOfRunes []rune
	if n < 0 {
		for n != 0 {
			arrayOfRunes = append(arrayOfRunes, rune(-(n%10)+48))
			n /= 10
		}
		arrayOfRunes = append(arrayOfRunes, '-')
	} else if n > 0 {
		for n != 0 {
			arrayOfRunes = append(arrayOfRunes, rune(n%10+48))
			n /= 10
		}
	} else {
		arrayOfRunes = append(arrayOfRunes, '0')
	}
	for i := len(arrayOfRunes) - 1; i >= 0; i-- {
		s += string(arrayOfRunes[i])
	}
	return s
}
