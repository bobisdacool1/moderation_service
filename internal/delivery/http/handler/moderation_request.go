package handler

import (
	"context"
	"fmt"
	"net/http"

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

func (h *ModerationRequestHandler) ApproveHandler(c *fiber.Ctx) error {
	return c.SendString("approved")
}

func (h *ModerationRequestHandler) CreateHandler(c *fiber.Ctx) error {
	r := new(entity.ModerationRequest)

	err := c.BodyParser(r)
	if err != nil {
		err = fmt.Errorf("body parse failed: %w", err)
		fmt.Print(err)
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	err = h.ModerationRequestUsecase.CreateModerationRequest(c.Context(), r)
	if err != nil {
		err = fmt.Errorf("create request failed: %w", err)
		fmt.Print(err)
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	err = c.SendStatus(http.StatusCreated)
	if err != nil {
		err = fmt.Errorf("create request failed: %w", err)
		fmt.Print(err)
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
