package middleware

import (
	"errors"
	"log"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

type PanicResponse struct {
	Issue string `json:"issue"`
}

func Panic() string {
	jsonData, err := sonic.Marshal(PanicResponse{
		Issue: "kontol",
	})

	if err != nil {
		log.Fatal(err)
	}

	return string(jsonData)
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	var panicResponse PanicResponse
	response := fiber.Map{
		"error": err.Error(),
	}

	switch code {
	case fiber.ErrBadRequest.Code:
		sonic.Unmarshal([]byte(err.Error()), &panicResponse)
		response = fiber.Map{
			"issue": panicResponse.Issue,
		}
		break
	}

	return c.Status(code).JSON(response)
}
