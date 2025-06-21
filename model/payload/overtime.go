// Package payload ...
package payload

// SubmitOvertimeRequest ...
type SubmitOvertimeRequest struct {
	Hours int `json:"hours" validate:"required,min=1"`
}
