package dna

import "errors"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, ch := range d {
		switch ch {
		case 'A', 'C', 'G', 'T':
			h[ch]++
		default:
			return nil, errors.New("fatal error: something went wrong")
		}
	}
	return h, nil
}
