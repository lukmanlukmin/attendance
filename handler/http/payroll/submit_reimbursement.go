// Package payroll ...
package payroll

import (
	"attendance/constant"
	"attendance/model/payload"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// SubmitReimbursement ...
// @Summary Submit Reimburse
// @Description Submit Reimburse
// @Tags Payroll
// @Accept json
// @Produce json
// @Router /v1/payrolls/overtime [post]
// @Success 200
// @Response 400
// @Response 500
// @Security BearerAuth
func (h *Handler) SubmitReimbursement(c *fiber.Ctx) error {
	req := payload.SubmitReimbursementRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(payload.ErrorResponse{
			Error: err.Error(),
		})
	}
	if err := h.Validate.Struct(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(payload.ErrorResponse{
			Error: err.Error(),
		})
	}
	err := h.Service.Payroll.SubmitReimbursement(c.Context(), req)
	if err != nil {
		return c.Status(constant.GetHTTPStatus(err)).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(http.StatusOK)
}
