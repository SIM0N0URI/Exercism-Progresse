package raindrops

func Convert(n int) string {
	s := ""
 	if n%3 == 0 {	
		s += "Pling"
	}
	if n%5 == 0 {
		s += "Plang"
	}
	if n%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		return Itoa(n)
	}
 return s
}
func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""
	for n > 0 {
		digit := (n % 10)+'0'
		result = string(rune(digit)) + result
		n = n / 10
	}
	return result
}