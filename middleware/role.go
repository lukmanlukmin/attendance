// Package middleware ...
package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Roles ...
func (m *Midleware) Roles(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rawRoleVal := c.Locals("role")
		if rawRoleVal == nil {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{
				"error": "missing role",
			})
		}

		roleStr, ok := rawRoleVal.(string)
		if !ok || roleStr == "" {
			return c.Status(http.StatusForbidden).JSON(fiber.Map{
				"error": "invalid role type",
			})
		}

		userRoles := strings.Split(roleStr, ",")
		roleSet := make(map[string]struct{})
		for _, role := range userRoles {
			roleSet[strings.ToLower(strings.TrimSpace(role))] = struct{}{}
		}

		for _, allowed := range allowedRoles {
			if _, ok := roleSet[strings.ToLower(allowed)]; ok {
				return c.Next()
			}
		}

		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": "access denied",
		})
	}
}
