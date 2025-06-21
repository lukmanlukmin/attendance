package payload

// ErrorResponse ...
type ErrorResponse struct {
	Error interface{} `json:"error"`
}
