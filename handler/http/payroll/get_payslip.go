// Package payroll ...
package payroll

import (
	"attendance/constant"
	"attendance/model/payload"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetPayslip ...
// @Summary Get Payslip Detail
// @Description Get Payslip Detail
// @Tags Payroll
// @Accept json
// @Produce json
// @Router /v1/payrolls/{id}/payslip [get]
// @Param id path string true "period ID"
// @Success 200 {object} payload.PayslipDetailResponse
// @Response 400
// @Response 500
// @Security BearerAuth
func (h *Handler) GetPayslip(c *fiber.Ctx) error {
	id := c.Params("id")
	payrollID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(payload.ErrorResponse{
			Error: err.Error(),
		})
	}
	res, err := h.Service.Payroll.GetPayslip(c.Context(), payrollID)
	if err != nil {
		return c.Status(constant.GetHTTPStatus(err)).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(res)
}
