// Package attendance ...
package attendance

import (
	"attendance/constant"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Submit ...
// @Summary Submit Attendance
// @Description Submit Attendance
// @Tags Attendance
// @Accept json
// @Produce json
// @Router /v1/attendances [post]
// @Success 200
// @Response 400
// @Response 500
// @Security BearerAuth
func (h *Handler) Submit(c *fiber.Ctx) error {
	err := h.Service.Attendance.SubmitAttendance(c.Context())
	if err != nil {
		return c.Status(constant.GetHTTPStatus(err)).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(http.StatusOK)
}
