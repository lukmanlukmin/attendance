// Package payroll ...
package payroll

import (
	"attendance/constant"
	"attendance/model/payload"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetSummaryPayslip ...
// @Summary Get Payslip Summary
// @Description Get Payslip Summary
// @Tags Payroll
// @Accept json
// @Produce json
// @Router /v1/payrolls/{id}/payslips [get]
// @Param id path string true "payroll ID"
// @Param page query int false "page number"
// @Param per_page query int false "per page"
// @Success 200 {object} payload.PayslipSummaryResponse
// @Response 400
// @Response 500
// @Security BearerAuth
func (h *Handler) GetSummaryPayslip(c *fiber.Ctx) error {
	id := c.Params("id")
	payrollID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(payload.ErrorResponse{
			Error: err.Error(),
		})
	}
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))

	res, err := h.Service.Payroll.GetResumePayslip(c.Context(), payrollID, uint64(page), uint64(perPage))
	if err != nil {
		return c.Status(constant.GetHTTPStatus(err)).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(res)
}
