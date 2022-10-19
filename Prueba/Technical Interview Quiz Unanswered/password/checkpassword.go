package password

import "fmt"

// CheckPassword returns an error if the input string is an invalid password per the
// rules below. Otherwise, it returns nil.
//
// MUST have a capital letter
// MUST contain at least one number
// MUST contain a punctuation mark or mathematical symbol
// MUST NOT contain the word "password"
// MUST be longer than 7 characters and shorter than 31 characters
//
func CheckPassword(pass string) error {

	var capitalLetter = false
	var oneNumber = false
	var PMark_MSymbol = false
	var lenRestriction = false
	var res = false
	var inicial = 0
	var counter = 0

	const word string = "J4cKaL=="

	// Validates if there is a capital letter in the string
	for i := 0; i < len(pass); i++ {
		if pass[i] >= 65 && pass[i] <= 90 {
			capitalLetter = true
		}
	}

	// Validates if the string contains at least one number
	for i := 0; i < len(pass); i++ {
		if pass[i] >= 48 && pass[i] <= 57 {
			oneNumber = true
		}
	}

	// Validates if there if it contains punctuation mark or mathematical symbol
	for i := 0; i < len(pass); i++ {
		if (pass[i] >= 33 && pass[i] <= 47) || pass[i] >= 58 && pass[i] <= 63 {
			//fmt.Println(pass[i])
			PMark_MSymbol = true
		}
	}

	// Validates the string has not the string "password"
	for i := 0; i < len(word); i++ {
		for j := 0; j < len(pass); j++ {
			if word[i] == pass[j] {
				inicial = j
				for k := inicial; k < (len(pass[k:])); k++ {
					//fmt.Println("longitud", len(pass[k:])+inicial-1)
					if len(word) <= len(pass[k:])+inicial+1 {
						fmt.Printf("Valor inicial %v\n, valor final %v\n", len(word), len(pass[inicial:]))
						if word[counter] == pass[len(word)+inicial] {
							counter++
							fmt.Println(counter)
						}
					}
				}
			}
		}
	}
	res = false
	if counter == len(word) {
		res = true
	}

	// Validates if the string is longer than 7 characters and shorter than 31 characters
	if len(pass) > 7 && len(pass) < 31 {
		lenRestriction = true
	}

	fmt.Printf("\nResult:\n \tCapital Letter: %v\n\tOne Number: %v\n\tPMark_MSymbol: %v\n\tlenRestriction: %v\n", capitalLetter, oneNumber, PMark_MSymbol, lenRestriction)

	//panic("implement me")

	panic("implement me")
}
