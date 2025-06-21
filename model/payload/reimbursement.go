// Package payload ...
package payload

// SubmitReimbursementRequest ...
type SubmitReimbursementRequest struct {
	Amount      int    `json:"amount" validate:"required,gt=0"`
	Description string `json:"description" validate:"required"`
}
