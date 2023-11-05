package pkg

import "time"

// getYearFromDate returns year from a given date
func getYearFromDate(dateString string) (int, error) {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return 0, err
	}

	return date.Year(), nil
}
