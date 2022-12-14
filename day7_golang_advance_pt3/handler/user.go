package handler

import (
	"user-management/entity/response"
	"user-management/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(usercase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: usercase}
}

func (h UserHandler) CreateUser(ctx *fiber.Ctx) error {
	userRequest := response.CreateUserRequest{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	if err := h.userUsecase.CreateUser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to create user",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.BaseResponse{
		Code:    fiber.StatusCreated,
		Message: "user created successfully",
		Data:    nil,
	})
}

func (h UserHandler) GetList(ctx *fiber.Ctx) error {
	users, err := h.userUsecase.GetList()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "get list users failed",
			Data:    nil,
		})
	}

	if len(users) < 1 {
		return ctx.Status(fiber.StatusNoContent).JSON(response.BaseResponse{
			Code:    fiber.StatusNoContent,
			Message: "no user found",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "successfully get all users",
		Data:    users,
	})
}

func (h UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	userRequest := response.UpdateUserRequest{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	if err := h.userUsecase.UpdateUser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to update user",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.BaseResponse{
		Code:    fiber.StatusCreated,
		Message: "user updated successfully",
		Data:    nil,
	})
}

func (h UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	userRequest := response.DeleteUserRequest{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "invalid req body",
			Data:    nil,
		})
	}
	if err := h.userUsecase.DeleteUser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to delete user",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.BaseResponse{
		Code:    fiber.StatusCreated,
		Message: "user deleted successfully",
		Data:    nil,
	})
}
