package helper

import (
	"encoding/json"
	"net/http"
)

// ReadRequestBody decodes a JSON request body into the provided result struct.
// It automatically panics if the decoding fails.
//
// Parameters:
//   - r: the incoming HTTP request containing a JSON body
//   - result: a pointer to the destination structure to populate
//
// Example:
//
//	var book request.CreateRequest
//	helper.ReadRequestBody(r, &book)
func ReadRequestBody(r *http.Request, result interface{}) {
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(result)
	ErrorPanic(err)
}

// WriteRequestBody encodes the result object as JSON and writes it to the response writer.
// It also sets the `Content-Type` to `application/json`.
//
// Parameters:
//   - w: the HTTP response writer
//   - result: the data to be encoded and sent back to the client
//
// Example:
//
//	response := Response{Code: 200, Status: "OK", Data: myData}
//	helper.WriteRequestBody(w, response)
func WriteRequestBody(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(result)
	ErrorPanic(err)
}
