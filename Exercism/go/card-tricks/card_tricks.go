package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
	panic("Please implement the FavoriteCards function")
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index > len(slice)-1 || index < 0 {
		return -1
	} else {
		return slice[index]
	}

	panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {

	if index > len(slice)-1 || index < 0 {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice

	panic("Please implement the SetItem function")
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
	panic("Please implement the PrependItems function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	/*fmt.Printf("Este es el slice Completo: %d\n", slice)
	fmt.Printf("Este es el indice: %d\n", index)
	fmt.Printf("Este es el slice Modificado: %d\n", append(slice[:index], slice[index+1:]...))*/
	if index > len(slice)-1 || index < 0 {
		return slice
	} else {
		return append(slice[:index], slice[index+1:]...)
	}
	panic("Please implement the RemoveItem function")
}
