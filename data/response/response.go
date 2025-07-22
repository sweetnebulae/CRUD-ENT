package response

// Response is a standard wrapper used for HTTP responses.
//
// It includes a status code, a message status, and optionally returned data.
type Response struct {
	Code   int         `json:"code"`           // HTTP status code (e.g., 200, 404, 500)
	Status string      `json:"status"`         // Status message (e.g., "OK", "Error")
	Data   interface{} `json:"data,omitempty"` // Actual data payload, optional
}
