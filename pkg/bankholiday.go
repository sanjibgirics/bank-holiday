package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Below are the variables which will be holding the processed data and which will be returning
// on respective API call.

// Holds bank holidays that fall within 2023 from all regions.
var holidays2023 []Event

// Holds bank holidays for the region "“England and Wales” where “bunting” is false.
var holidaysEnglandAndWalesNoBunting []Event

// Holds bank holidays title and date that fall within 2024 from all regions.
var holidays2024TitleAndDate []struct {
	Title string `json:"title"`
	Date  string `json:"date"`
}

// Fetch all the bank holidays data from the intended url.
func fetchBankHolidayData() BankHolidays {
	log.Println("Fetching all the bank-holidays data.")
	url := "https://www.gov.uk/bank-holidays.json"
	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Errorf("error while fetching data from the url %v: %w", url, err))
	}
	defer resp.Body.Close()

	var bankHolidaysData BankHolidays
	if err := json.NewDecoder(resp.Body).Decode(&bankHolidaysData); err != nil {
		panic(fmt.Errorf("error while decoding the data: %w", err))
	}

	return bankHolidaysData
}

// Process all the fetched bank-holiday data according to the endpoint requirements and store
// them.
func processBankHolidayData(bankHolidayData BankHolidays) {
	log.Println("Processing all the bank-holidays data.")

	for regionName, region := range bankHolidayData {
		for _, event := range region.Events {
			// Handles bank holidays for the region "England and Wales" where “bunting” is false.
			if regionName == "england-and-wales" && !event.Bunting {
				holidaysEnglandAndWalesNoBunting = append(holidaysEnglandAndWalesNoBunting,
					event)
			}

			year, err := getYearFromDate(event.Date)
			if err != nil {
				panic(fmt.Errorf("error while getting year from the holiday date: %w", err))
			}

			// Handles bank holidays that fall within 2023 from all regions.
			if year == 2023 {
				holidays2023 = append(holidays2023, event)
			} else if year == 2024 {
				// Handles bank holidays title and date that fall within 2024 from all regions.
				holidayTitleAndDate := struct {
					Title string `json:"title"`
					Date  string `json:"date"`
				}{event.Title, event.Date}
				holidays2024TitleAndDate = append(holidays2024TitleAndDate, holidayTitleAndDate)
			}
		}
	}
}

// Fetch and process the bank holidays data.
func FetchAndProcessBankHolidayData() {
	bankHolidayData := fetchBankHolidayData()
	processBankHolidayData(bankHolidayData)
}

// Functions which returns the data stored to the endpoints as per their requirements.

// Returns all bank holidays that fall within 2023 from all regions.
func getHolidays2023Data() ([]Event, error) {
	return holidays2023, nil
}

// Returns all bank holidays for the region "England and Wales" where “bunting” is false.
func getHolidaysEnglandAndWalesNoBuntingData() ([]Event, error) {
	return holidaysEnglandAndWalesNoBunting, nil
}

// Returns all bank holidays title and date that fall within 2024 from all regions.
func getHolidays2024TitleAndDateData() ([]struct {
	Title string `json:"title"`
	Date  string `json:"date"`
}, error) {
	return holidays2024TitleAndDate, nil
}
