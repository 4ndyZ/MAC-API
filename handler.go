package main

import (
	"encoding/csv"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"strings"
)

func (a *App) GetMACData(httpCSV string) ([][]string, error) {
	// Perform HTTP call
	resp, err := http.Get(httpCSV)
	if err != nil {
		return [][]string{}, err
	}
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	body, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (a *App) ParseMACData(body *[][]string) ([]OUI, error) {
	// Slice with parsed data
	var ouis []OUI
	// Regex
	reg, _ := regexp.Compile(" [^(a-zA-Z0-9)]*$") // Remove spaces after last alphanumeric char
	// Search for data using Regex
	for _, row := range (*body)[1:] { // Skip first element because it contains headers
		// Struct
		var oui OUI
		// Build struct
		oui.Vendor = row[2]
		oui.OUI = SplitMAC(row[1])
		oui.Typ = row[0]
		oui.Address = reg.ReplaceAllString(row[3], "")
		// Append to parsed data the slice
		ouis = append(ouis, oui)
	}
	return ouis, nil
}

func (a *App) GetOUI(w http.ResponseWriter, r *http.Request) {
	input := strings.ToUpper((mux.Vars(r))["oui"]) // Convert OUI to upper case

	// Remove all non HEX chars
	reg, _ := regexp.Compile("[^0-9A-F]+")
	request := reg.ReplaceAllString(input, "")

	if len(request) != 6 { // Length of OUI has to be 6
		Log.Logger.Debug().Str("request-typ", "v1/oui").Str("input", input).Str("error", "Invalid OUI").Msg("Invalid API request")
		a.respondWithError(w, http.StatusBadRequest, "Invalid OUI")
		return
	}

	request = SplitMAC(request) // Split OUI into format XX-XX-XX

	// Check if request matches OUI entry in data
	for _, oui := range *a.Data {
		if oui.OUI == request {
			a.respondWithJSON(w, http.StatusOK, oui)
			Log.Logger.Debug().Str("request-typ", "v1/oui").Str("vendor", oui.Vendor).Str("oui", oui.OUI).Str("typ", oui.Typ).Str("Address", oui.Address).Msg("API request")
			return
		}
	}

	a.respondWithError(w, http.StatusNotFound, "Invalid OUI")
	Log.Logger.Debug().Str("request-typ", "v1/oui").Str("request", request).Str("error", "Invalid OUI").Msg("Invalid API request")
}

func (a *App) GetMAC(w http.ResponseWriter, r *http.Request) {
	input := strings.ToUpper((mux.Vars(r))["mac"]) // Convert OUI to upper case

	// Remove all non HEX chars
	reg, _ := regexp.Compile("[^0-9A-F]+")
	request := reg.ReplaceAllString(input, "")

	if len(request) != 12 { // Length of OUI has to be 6
		Log.Logger.Debug().Str("request-typ", "v1/mac").Str("input", input).Str("error", "Invalid OUI").Msg("Invalid API request")
		a.respondWithError(w, http.StatusBadRequest, "Invalid OUI")
		return
	}

	requestOUI := SplitMAC(request[0:6]) // Get OUI from MAC

	request = SplitMAC(request) // Split OUI into format XX-XX-XX-XX-XX-XX

	// Check if request OUI matches OUI entry in data
	for _, oui := range *a.Data {
		if oui.OUI == requestOUI {
			mac := MAC{
				MAC:     request,
				Vendor:  oui.Vendor,
				OUI:     oui.OUI,
				Typ:     oui.Typ,
				Address: oui.Address,
			}
			a.respondWithJSON(w, http.StatusOK, mac)
			Log.Logger.Debug().Str("request-typ", "v1/mac").Str("mac", mac.MAC).Str("vendor", mac.Vendor).Str("oui", mac.OUI).Str("typ", mac.Typ).Str("Address", mac.Address).Msg("API request")
			return
		}
	}

	a.respondWithError(w, http.StatusNotFound, "OUI not found")
	Log.Logger.Debug().Str("request-typ", "v1/mac").Str("request", request).Str("error", "OUI not found").Msg("Invalid API request")
}

func (a *App) NotFound(w http.ResponseWriter, r *http.Request) {
	a.respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
}

// Send a payload of JSON content
func (a *App) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Send a JSON error message
func (a *App) respondWithError(w http.ResponseWriter, code int, message string) {
	a.respondWithJSON(w, code, map[string]string{"error": message})
}
