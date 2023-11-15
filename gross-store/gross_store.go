package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, unitExists := units[unit]
	// if the unit is non-existent, return false
	if !unitExists {
		return false
	} else {
		_, itemExists := bill[item]
		// if the item doesn't exist, add it to the customer bill
		if !itemExists {
			bill[item] = units[unit]
			return true
		} else {
			// if item is already in the customer bill, increment it by the value in the units map
			bill[item] += units[unit]
			return true
		}
	}

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, itemExists := bill[item]
	_, unitExists := units[unit]
	// item is not in the customer bill or unit doesn't exist
	if !itemExists || !unitExists {
		return false
	}

	currQuantity := bill[item]
	newQuantity := currQuantity - units[unit]

	if newQuantity < 0 {
		return false
	} else if newQuantity == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] -= units[unit]
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, exits := bill[item]
	if !exits {
		return 0, false
	}
	return bill[item], true
}
