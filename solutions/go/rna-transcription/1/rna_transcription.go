package strand

func ToRNA(dna string) string {
    res := ""
	for _, ch := range dna{
        switch ch {
            case 'G':
            res += "C"
            case 'C':
            res += "G"
            case 'T':
            res += "A"
            case 'A':
            res += "U"            
        }
    }
	return res
}