package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"ModerationService/internal/entity"
)

type (
	ModerationRequestHandler struct {
		ModerationRequestUsecase ModerationRequestUsecase
	}

	ModerationRequestUsecase interface {
		CreateModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
		ApproveModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
		DeclineModerationRequest(ctx context.Context, request *entity.ModerationRequest) error
		NextModerationRequest(ctx context.Context) (*entity.ModerationRequest, error)
	}
)

func NewModerationRequestHandler(moderationRequestUsecase ModerationRequestUsecase) *ModerationRequestHandler {
	return &ModerationRequestHandler{
		ModerationRequestUsecase: moderationRequestUsecase,
	}
}

func (h *ModerationRequestHandler) DeclineHandler(c *fiber.Ctx) error {
	return c.SendString("declined")
}

func (h *ModerationRequestHandler) NextHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"id": "abc123", "text": "пример"})
}

func (h *ModerationRequestHandler) ApproveHandler(c *fiber.Ctx) error {
	return c.SendString("approved")
}

func (h *ModerationRequestHandler) CreateHandler(c *fiber.Ctx) error {
	return c.SendString("created")
}
