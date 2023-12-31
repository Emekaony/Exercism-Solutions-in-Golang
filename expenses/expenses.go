package expenses

import "fmt"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	records := []Record{}
	for _, record := range in {
		if predicate(record) {
			records = append(records, record)
		}
	}
	return records
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		ans := false
		if r.Day >= p.From && r.Day <= p.To {
			ans = true
		}
		return ans
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		ans := false
		if r.Category == c {
			ans = true
		}
		return ans
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	recordsInPeriod := Filter(in, ByDaysPeriod(p))
	sum := 0.0
	for _, r := range recordsInPeriod {
		sum += r.Amount
	}
	return sum
}

type CategoryError struct {
	category string
}

// when something implements the Error() method.
// if u have to return an error, then the Error() method will be called.
func (c *CategoryError) Error() string {
	return fmt.Sprintf("error(unknown category %s)", c.category)
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	recordsInCategory := Filter(in, ByCategory(c))
	// this is when there are no records in the given category
	if len(recordsInCategory) == 0 {
		// Had to define my own custom categoryError type and return a reference to it!
		return 0, &CategoryError{
			category: c,
		}
	} else {
		recordsInPeriod := Filter(recordsInCategory, ByDaysPeriod(p))
		if len(recordsInPeriod) == 0 {
			return 0, nil
		} else {
			return TotalByPeriod(recordsInPeriod, p), nil
		}
	}
}
