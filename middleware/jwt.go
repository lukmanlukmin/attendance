// Package middleware ...
package middleware

import (
	"attendance/constant"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lukmanlukmin/go-lib/util"
)

// JWT ...
func (m *Midleware) JWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		authHeaderLower := strings.ToLower(authHeader)
		if !strings.HasPrefix(authHeaderLower, "bearer ") {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing or invalid Authorization header",
			})
		}
		tokenStr := strings.TrimSpace(authHeader[len("Bearer "):])

		claimData, err := util.ValidateJWT(m.Conf.Security.JWTSecret, tokenStr)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		requiredClaims := []string{constant.ContextUserIDKey, constant.ContextRoleKey}
		for _, key := range requiredClaims {
			if _, ok := claimData[key]; !ok {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"error": fmt.Sprintf("jwt. missing required claim: %s", key),
				})
			}
		}

		userIDStr, ok := claimData[constant.ContextUserIDKey].(string)
		if !ok || userIDStr == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "jwt. missing user_id",
			})
		}
		if _, err := uuid.Parse(userIDStr); err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "jwt. invalid user_id",
			})
		}
		employeeIDStr, ok := claimData[constant.ContextEmployeeIDKey].(string)
		if ok && employeeIDStr != "" {
			if _, err := uuid.Parse(employeeIDStr); err == nil {
				c.Locals(constant.ContextEmployeeIDKey, employeeIDStr)
			}
		}

		role, ok := claimData[constant.ContextRoleKey].(string)
		if !ok || role == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "jwt. invalid username",
			})
		}

		c.Locals(constant.ContextUserIDKey, userIDStr)
		c.Locals(constant.ContextRoleKey, role)
		c.Locals(constant.ContextIPKey, c.IP())
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.NewString()
		}
		c.Locals(constant.ContextRequestIDKey, requestID)
		return c.Next()
	}
}
