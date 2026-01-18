package response

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success   bool   `json:"success"`
	Status    int    `json:"status"`
	Code      int    `json:"code,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	Errors    any    `json:"errors,omitempty"`
}

func Success(
	c *fiber.Ctx,
	status int,
	code Code,
	data any,
) error {
	return c.Status(status).JSON(Response{
		Success: true,
		Status:  status,
		Code:    code.Code,
		Message: code.Message,
		Data:    data,
	})
}

func Error(
	c *fiber.Ctx,
	status int,
	code Code,

	errors any,
) error {
	return c.Status(status).JSON(Response{
		Success:   false,
		Status:    status,
		Code:      code.Code,
		ErrorCode: code.ErrorCode,
		Message:   code.Message,
		Errors:    errors,
	})
}
