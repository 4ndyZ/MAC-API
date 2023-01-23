package main

import (
	"encoding/csv"
	"github.com/gofiber/fiber/v2"
	"io"
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Log.Logger.Warn().Str("error", err.Error()).Msg("Error while closing the HTTP response body.")
		}
	}(resp.Body)
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
	reg := regexp.MustCompile(" [^(a-zA-Z0-9)]*$") // Remove spaces after last alphanumeric char
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

func (a *App) GetOUI(c *fiber.Ctx) error {
	input := strings.ToUpper(c.Params("oui")) // Convert OUI to upper case
	request := RemoveAllNonHex(input)         // Remove all non-HEX chars

	if len(request) != 6 { // Length of OUI has to be 6
		Log.Logger.Debug().Str("request-typ", "v1/oui").Str("input", input).Str("error", "Invalid OUI").Msg("Invalid API request")
		return a.respondWithError(c, http.StatusBadRequest, "Invalid OUI")
	}

	request = SplitMAC(request) // Split OUI into format XX-XX-XX

	// Check if request matches OUI entry in data
	for _, oui := range *a.Data {
		if oui.OUI == request {
			Log.Logger.Debug().Str("request-typ", "v1/oui").Str("vendor", oui.Vendor).Str("oui", oui.OUI).Str("typ", oui.Typ).Str("Address", oui.Address).Msg("API request")
			return a.respondWithJSON(c, http.StatusOK, oui)
		}
	}
	Log.Logger.Debug().Str("request-typ", "v1/oui").Str("request", request).Str("error", "Invalid OUI").Msg("Invalid API request")
	return a.respondWithError(c, http.StatusNotFound, "Invalid OUI")
}

func (a *App) GetMAC(c *fiber.Ctx) error {
	input := strings.ToUpper(c.Params("mac")) // Convert OUI to upper case
	request := RemoveAllNonHex(input)         // Remove all non-HEX chars

	if len(request) != 12 { // Length of OUI has to be 6
		Log.Logger.Debug().Str("request-typ", "v1/mac").Str("input", input).Str("error", "Invalid OUI").Msg("Invalid API request")
		return a.respondWithError(c, http.StatusBadRequest, "Invalid OUI")
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
			Log.Logger.Debug().Str("request-typ", "v1/mac").Str("mac", mac.MAC).Str("vendor", mac.Vendor).Str("oui", mac.OUI).Str("typ", mac.Typ).Str("Address", mac.Address).Msg("API request")
			return a.respondWithJSON(c, http.StatusOK, mac)
		}
	}
	Log.Logger.Debug().Str("request-typ", "v1/mac").Str("request", request).Str("error", "OUI not found").Msg("Invalid API request")
	return a.respondWithError(c, http.StatusNotFound, "OUI not found")
}

func (a *App) NotFound(c *fiber.Ctx) error {
	return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": "Invalid request"})
}

// Send a payload of JSON content
func (a *App) respondWithJSON(c *fiber.Ctx, code int, payload interface{}) error {
	return c.Status(code).JSON(payload)
}

// Send a JSON error message
func (a *App) respondWithError(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(map[string]string{"error": message})
}
