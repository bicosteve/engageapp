package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/engageapp/pkg/entities"
)

// ReadJSON -> helper function for reading json from client
// Takes json, decodes for error.
// Returns error if any
func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// decoding the data received from client into go types
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	if err != nil {
		return err
	}

	// Check that json value from client is valid
	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must contain a single JSON value")
	}

	return nil
}

// WriteJSON -> helper for writing json to  client
// Marshal's data
// Sets header content type
func WriteJSON(
	w http.ResponseWriter, status int, data interface{}, headers ...http.Header,
) error {
	toSee, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	// Set Header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(toSee)
	if err != nil {
		return err
	}

	return nil
}

// ErrorJSON -> helper to returns json error to client
func ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload entities.JSONResponse

	payload.Error = true
	payload.Message = err.Error()

	return WriteJSON(w, statusCode, payload)

}

func GenerateAccessToken() (string, error) {

	return "", nil
}
