package expenses

import "errors"

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
	filteredRecords := []Record{}
	for _, expenseRecord := range in {
		isValid := predicate(expenseRecord)

		if isValid {
			filteredRecords = append(filteredRecords, expenseRecord)
		}
	}
	return filteredRecords
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(rec Record) bool {
		if rec.Day >= p.From && rec.Day <= p.To {
			return true
		}
		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(rec Record) bool {
		return rec.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var totalExpense float64
	for _, rec := range in {
		if rec.Day >= p.From && rec.Day <= p.To {
			totalExpense += rec.Amount
		}
	}
	return totalExpense
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var hasCategory bool
	var totalExpense float64
	for _, rec := range in {

		if rec.Category == c {
			hasCategory = true
		}

		if rec.Day >= p.From && rec.Day <= p.To && rec.Category == c {
			totalExpense += rec.Amount
		}
	}

	if !hasCategory {
		return 0, errors.New("unknown category entertainment")
	}

	return totalExpense, nil
}
