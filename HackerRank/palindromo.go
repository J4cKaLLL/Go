func solution(inputString string) bool {    
	full := len(inputString)
	middle := len(inputString) / 2
	middleMod := len(inputString) % 2
    first := inputString[:middle]	
	if middleMod != 0 {
		middle = middle + 1
	}	
	var slice string
	for i := full; i > middle; i-- {
		slice = slice + string(inputString[i-1])
	}
	if slice == first {
		return true
	} else {
		return false
	}
}
