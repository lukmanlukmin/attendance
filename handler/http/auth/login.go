// Package auth ...
package auth

import (
	"attendance/constant"
	"attendance/model/payload"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Login ...
// @Summary Login
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Router /v1/authorize/login [post]
// @Param payload body payload.LoginCredential true "User Credential Payload"
// @Success 200 {object} payload.Token
func (h *Handler) Login(c *fiber.Ctx) error {
	req := payload.LoginCredential{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	if err := h.Validate.Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	res, err := h.Service.Auth.ValidateUserByCredential(c.Context(), req)
	if err != nil {
		return c.Status(constant.GetHTTPStatus(err)).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(res)
}
