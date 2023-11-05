package pkg

// Schema for Event
type Event struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Notes   string `json:"notes"`
	Bunting bool   `json:"bunting"`
}

// Schema for Region
type Region struct {
	Division string  `json:"division"`
	Events   []Event `json:"events"`
}

// Schema for BankHolidays which contain all the data
type BankHolidays map[string]Region
