// Package attendance ...
package attendance

import (
	"attendance/constant"
	"attendance/model/payload"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Overtime ...
// @Summary Request Overtime
// @Description Request Overtime
// @Tags Attendance
// @Accept json
// @Produce json
// @Router /v1/attendances/overtime [post]
// @Success 200
// @Response 400
// @Response 500
// @Security BearerAuth
func (h *Handler) Overtime(c *fiber.Ctx) error {
	req := payload.SubmitOvertimeRequest{}
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
	err := h.Service.Attendance.SubmitOvertime(c.Context(), req)
	if err != nil {
		return c.Status(constant.GetHTTPStatus(err)).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(http.StatusOK)
}
