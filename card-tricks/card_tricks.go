package cards

// checks if the index is out of bounds
func isOutOfBounds(slice []int, index int) bool {
	if index >= len(slice) || index < 0 {
		return true
	} else {
		return false
	}
}

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if isOutOfBounds(slice, index) {
		return -1
	} else {
		return slice[index]
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if isOutOfBounds(slice, index) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	// go treats variadic parameters as slices. Cool!
	if values == nil {
		return slice
	} else {
		values = append(values, slice...)
		return values
	}
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if isOutOfBounds(slice, index) {
		return slice
	} else {
		return append(slice[:index], slice[index+1:]...)
	}
}
