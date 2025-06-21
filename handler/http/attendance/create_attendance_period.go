// Package attendance ...
package attendance

import (
	"attendance/constant"
	"attendance/model/payload"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// CreateAttendacePeriod ...
// @Summary Create Attendance Period
// @Description Create Attendance Period
// @Tags Attendance
// @Accept json
// @Produce json
// @Router /v1/attendances/period [post]
// @Success 200
// @Response 400
// @Response 500
// @Security BearerAuth
func (h *Handler) CreateAttendacePeriod(c *fiber.Ctx) error {
	req := payload.CreateAttendancePeriodRequest{}
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
	err := h.Service.Payroll.CreateAttendancePeriod(c.Context(), req)
	if err != nil {
		return c.Status(constant.GetHTTPStatus(err)).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(http.StatusOK)
}
