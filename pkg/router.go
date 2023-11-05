package pkg

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// RegisterRoutes register all the endpoints needed to the router
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/getHolidays2023", getHolidays2023).Methods("GET")
	r.HandleFunc("/getHolidaysEnglandAndWalesNoBunting", getHolidaysEnglandAndWalesNoBunting).
		Methods("GET")
	r.HandleFunc("/getHolidays2024TitleAndDate", getHolidays2024TitleAndDate).
		Methods("GET")
}

// getHolidays2023 is the handle function for the endpoint /getHolidays2023 which returns all
// bank holidays that fall within 2023 from all regions.
func getHolidays2023(w http.ResponseWriter, r *http.Request) {
	// Retrieve required data
	holidays2023Data, err := getHolidays2023Data()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(holidays2023Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// getHolidaysEnglandAndWalesNoBunting is the handle function for the endpoint /getHolidaysEnglandAndWalesNoBunting
// which returns all bank holidays for the region "England and Wales" where “bunting” is false.
func getHolidaysEnglandAndWalesNoBunting(w http.ResponseWriter, r *http.Request) {
	// Retrieve required data
	holidaysEnglandAndWalesNoBuntingData, err := getHolidaysEnglandAndWalesNoBuntingData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(holidaysEnglandAndWalesNoBuntingData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// getHolidays2024TitleAndDate is the handle function for the endpoint /getHolidays2024TitleAndDate
// which returns all bank holidays title and date that fall within 2024 from all regions.
func getHolidays2024TitleAndDate(w http.ResponseWriter, r *http.Request) {
	// Retrieve required data
	holidays2024TitleAndDateData, err := getHolidays2024TitleAndDateData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(holidays2024TitleAndDateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handle homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to UK bank holiday application"))
}
