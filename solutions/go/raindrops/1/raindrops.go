package raindrops

func Convert(number int) string {
    if number % 5 == 0 && number % 3 == 0 && number % 7 == 0{
        return "PlingPlangPlong"
	}else if number % 5 == 0 && number % 3 == 0 {
        return "PlingPlang"
    }else if number % 5 == 0 && number % 7 == 0 {
        return "PlangPlong"
    }else if number % 3 == 0 && number % 7 == 0 {
        return "PlingPlong"
    }else if number % 5 == 0 {
        return "Plang"
    }else if number % 3 == 0 {
        return "Pling"
    }else if number % 7 == 0 {
        return "Plong"
    }else {
        return itoa(number)
    }
}

func itoa(n int) string {
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