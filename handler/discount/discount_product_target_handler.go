package discount

import (
	"errors"
	discount_ctrl "service/discount/api/controller/discount"
	reqmodel "service/discount/api/model/request"
	resmodel "service/discount/api/model/response"
	"service/discount/api/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type DiscountProductTargetHandler struct {
	Controller *discount_ctrl.PostController
	Validate   *validator.Validate
}

func NewDiscountProductTargetHandler(
	r fiber.Router,
	vld *validator.Validate,
	controller *discount_ctrl.PostController,
) {
	handler := &DiscountProductTargetHandler{
		Controller: controller,
		Validate:   vld,
	}

	rStrict := r.Group("discount-product")
	rStrict.Get("/", handler.ListDiscountProductTarget)
	rStrict.Get("/:id", handler.GetDiscountProductTargetByID)
	rStrict.Post("/", handler.CreateDiscountProductTarget)
	rStrict.Put("/:id", handler.UpdateDiscountProductTarget)
	rStrict.Delete("/:id", handler.DeleteDiscountProductTarget)
}

// ListDiscountProductTarget
//
//	@Summary		List Discount Product Target
//	@Description	Get all discount product targets
//	@Tags			Discount Product Target
//	@Produce		json
//	@Param			search	query		string					false	"Search keyword"
//	@Param			page	query		int						false	"Page Number"
//	@Param			size	query		int						false	"Page Size"
//	@Success		200		{object}	resmodel.DatasResponse	"Discount product targets retrieved successfully"
//	@Failure		400		{object}	utils.RequestError		"Bad request"
//	@Failure		404		{object}	utils.RequestError		"Discount product targets not found"
//	@Failure		500		{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount-product [get]
func (h *DiscountProductTargetHandler) ListDiscountProductTarget(c *fiber.Ctx) error {
	// Parse
	var request reqmodel.ListRequest
	if err := c.QueryParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("format query tidak valid"),
		))
	}

	// Validate Input
	if err := h.Validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("validasi gagal: "+err.Error()),
		))
	}

	// Call Controller
	datas, pagination, err := h.Controller.ListDiscountProductTarget(c.Context(), request)
	if err != nil {
		if reqErr, ok := err.(utils.RequestError); ok {
			return c.Status(reqErr.StatusCode).JSON(reqErr)
		}
		if stdErr, ok := err.(utils.StandardError); ok {
			return c.Status(stdErr.Code).JSON(stdErr)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(utils.GeneralErrorResponse(
			fiber.StatusInternalServerError,
			err,
		))
	}

	// Serve JSON
	response := resmodel.DatasResponse{
		Success:    true,
		Message:    "Berhasil mengambil data target diskon produk",
		Datas:      datas,
		Pagination: pagination,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetDiscountProductTargetByID
//
//	@Summary		Get Discount Product Target By ID
//	@Description	Get a discount product target by its ID
//	@Tags			Discount Product Target
//	@Produce		json
//	@Param			id	path		string					true	"Discount Product Target ID"
//	@Success		200	{object}	resmodel.DataResponse	"Discount product target retrieved successfully"
//	@Failure		400	{object}	utils.RequestError		"Bad request"
//	@Failure		404	{object}	utils.RequestError		"Discount product target not found"
//	@Failure		500	{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount-product/{id} [get]
func (h *DiscountProductTargetHandler) GetDiscountProductTargetByID(c *fiber.Ctx) error {
	// Parse
	productID := c.Params("id")

	// Call Controller
	data, err := h.Controller.GetDiscountProductTargetByID(c.Context(), productID)
	if err != nil {
		if reqErr, ok := err.(utils.RequestError); ok {
			return c.Status(reqErr.StatusCode).JSON(reqErr)
		}
		if stdErr, ok := err.(utils.StandardError); ok {
			return c.Status(stdErr.Code).JSON(stdErr)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(utils.GeneralErrorResponse(
			fiber.StatusInternalServerError,
			err,
		))
	}

	// Serve JSON
	response := resmodel.DataResponse{
		Success: true,
		Message: "Berhasil mengambil data target diskon produk",
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// CreateDiscountProductTarget
//
//	@Summary		Create Discount Product Target
//	@Description	Create a new discount product target
//	@Tags			Discount Product Target
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string									true	"User ID to check"
//	@Param			request			body		reqmodel.CreateDiscountProductTarget	true	"Create discount product target"
//	@Success		201				{object}	resmodel.NoDataResponse					"Discount product target created"
//	@Failure		400				{object}	utils.RequestError						"Bad request"
//	@Failure		500				{object}	utils.RequestError						"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount-product [post]
func (h *DiscountProductTargetHandler) CreateDiscountProductTarget(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	var request reqmodel.CreateDiscountProductTarget
	request.AuthUserID = authUserID
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("format JSON tidak valid"),
		))
	}

	// Validate Input
	if err := h.Validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("validasi gagal: "+err.Error()),
		))
	}

	// Call Controller
	err := h.Controller.CreateDiscountProductTarget(c.Context(), request)
	if err != nil {
		if reqErr, ok := err.(utils.RequestError); ok {
			return c.Status(reqErr.StatusCode).JSON(reqErr)
		}
		if stdErr, ok := err.(utils.StandardError); ok {
			return c.Status(stdErr.Code).JSON(stdErr)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(utils.GeneralErrorResponse(
			fiber.StatusInternalServerError,
			errors.New(err.Error()),
		))
	}

	// Serve JSON
	response := resmodel.NoDataResponse{
		Success: true,
		Message: "Berhasil membuat target diskon produk",
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

// UpdateDiscountProductTarget
//
//	@Summary		Update Discount Product Target
//	@Description	Update an existing discount product target
//	@Tags			Discount Product Target
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string									true	"User ID to check"
//	@Param			id				path		string									true	"Discount Product Target ID"
//	@Param			request			body		reqmodel.UpdateDiscountProductTarget	true	"Update discount product target"
//	@Success		200				{object}	resmodel.NoDataResponse					"Discount product target updated"
//	@Failure		400				{object}	utils.RequestError						"Bad request"
//	@Failure		500				{object}	utils.RequestError						"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount-product/{id} [put]
func (h *DiscountProductTargetHandler) UpdateDiscountProductTarget(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	productID := c.Params("id")
	var request reqmodel.UpdateDiscountProductTarget
	request.ID = productID
	request.AuthUserID = authUserID
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("format JSON tidak valid"),
		))
	}

	// Validate Input
	if err := h.Validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("validasi gagal: "+err.Error()),
		))
	}

	// Call Controller
	err := h.Controller.UpdateDiscountProductTarget(c.Context(), request)
	if err != nil {
		if reqErr, ok := err.(utils.RequestError); ok {
			return c.Status(reqErr.StatusCode).JSON(reqErr)
		}
		if stdErr, ok := err.(utils.StandardError); ok {
			return c.Status(stdErr.Code).JSON(stdErr)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(utils.GeneralErrorResponse(
			fiber.StatusInternalServerError,
			errors.New(err.Error()),
		))
	}

	// Serve JSON
	response := resmodel.NoDataResponse{
		Success: true,
		Message: "Berhasil mengubah target diskon produk",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteDiscountProductTarget
//
//	@Summary		Delete Discount Product Target
//	@Description	Delete a discount product target
//	@Tags			Discount Product Target
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string					true	"User ID to check"
//	@Param			id				path		string					true	"Discount Product Target ID"
//	@Success		200				{object}	resmodel.NoDataResponse	"Discount product target deleted"
//	@Failure		400				{object}	utils.RequestError		"Bad request"
//	@Failure		500				{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount-product/{id} [delete]
func (h *DiscountProductTargetHandler) DeleteDiscountProductTarget(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	productID := c.Params("id")

	// Call Controller
	err := h.Controller.DeleteDiscountProductTarget(c.Context(), productID, authUserID)
	if err != nil {
		if reqErr, ok := err.(utils.RequestError); ok {
			return c.Status(reqErr.StatusCode).JSON(reqErr)
		}
		if stdErr, ok := err.(utils.StandardError); ok {
			return c.Status(stdErr.Code).JSON(stdErr)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(utils.GeneralErrorResponse(
			fiber.StatusInternalServerError,
			errors.New(err.Error()),
		))
	}

	// Serve JSON
	response := resmodel.NoDataResponse{
		Success: true,
		Message: "Berhasil menghapus target diskon produk",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
