// Package payroll ...
package payroll

import (
	"attendance/constant"
	"attendance/model/payload"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreatePayroll ...
// @Summary Create Payroll
// @Description Create Payroll
// @Tags Payroll
// @Accept json
// @Produce json
// @Router /v1/payrolls/period/{id} [post]
// @Param id path string true "period ID"
// @Success 200
// @Response 400
// @Response 500
// @Security BearerAuth
func (h *Handler) CreatePayroll(c *fiber.Ctx) error {
	id := c.Params("id")
	periodID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(payload.ErrorResponse{
			Error: err.Error(),
		})
	}
	err = h.Service.Payroll.CreatePayroll(c.Context(), periodID)
	if err != nil {
		return c.Status(constant.GetHTTPStatus(err)).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(http.StatusOK)
}
