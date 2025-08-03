package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    step := 0
	if n <= 0 {
		return 0, errors.New("input must be a positive integer")
	}
    for n !=1 {
		if n % 2 == 0{
        	n = n / 2
    	}else {
        	n = n*3 +1   
    	}
        step++
    }
    return step, nil
}
