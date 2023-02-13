package controller

import "github.com/gofiber/fiber/v2"

// GetIrancell godoc
// @Summary Shows Irancells
// @Description get the status of server.
// @Tags irancell
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/irancell [get]
func GetIrancell(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]any{
		"message": "message",
	})
}
