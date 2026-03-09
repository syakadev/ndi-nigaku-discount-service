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

type DiscountHandler struct {
	Controller *discount_ctrl.DiscountController
	Validate   *validator.Validate
}

func NewDiscountHandler(
	r fiber.Router,
	vld *validator.Validate,
	controller *discount_ctrl.DiscountController,
) {
	handler := &DiscountHandler{
		Controller: controller,
		Validate:   vld,
	}

	rStrict := r.Group("discount")
	rStrict.Get("/", handler.ListDiscount)
	rStrict.Get("/:id_unik", handler.GetDiscountByID)
	rStrict.Post("/", handler.CreateDiscount)
	rStrict.Put("/:id_unik", handler.UpdateDiscount)
	rStrict.Delete("/:id_unik", handler.DeleteDiscount)
}

// ListDiscount
//
//	@Summary		List Discount
//	@Description	Get all discounts
//	@Tags			Discount
//	@Produce		json
//	@Param			search	query		string					false	"Search keyword"
//	@Param			page	query		int						false	"Page Number"
//	@Param			size	query		int						false	"Page Size"
//	@Success		200		{object}	resmodel.DatasResponse	"Discounts retrieved successfully"
//	@Failure		400		{object}	utils.RequestError		"Bad request"
//	@Failure		404		{object}	utils.RequestError		"Discounts not found"
//	@Failure		500		{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount [get]
func (h *DiscountHandler) ListDiscount(c *fiber.Ctx) error {
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
	datas, pagination, err := h.Controller.ListDiscount(c.Context(), request)
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
		Message:    "Berhasil mengambil data diskon",
		Datas:      datas,
		Pagination: pagination,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetDiscountByID
//
//	@Summary		Get Discount By ID
//	@Description	Get a discount by its ID
//	@Tags			Discount
//	@Produce		json
//	@Param			id_unik	path		string					true	"Discount ID"
//	@Success		200	{object}	resmodel.DataResponse	"Discount retrieved successfully"
//	@Failure		400	{object}	utils.RequestError		"Bad request"
//	@Failure		404	{object}	utils.RequestError		"Discount not found"
//	@Failure		500	{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount/{id_unik} [get]
func (h *DiscountHandler) GetDiscountByID(c *fiber.Ctx) error {
	// Parse
	discountID := c.Params("id_unik")

	// Call Controller
	data, err := h.Controller.GetDiscountByID(c.Context(), discountID)
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
		Message: "Berhasil mengambil data diskon",
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// CreateDiscount
//
//	@Summary		Create Discount
//	@Description	Create a new discount
//	@Tags			Discount
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string					true	"User ID to check"
//	@Param			request			body		reqmodel.CreateDiscount	true	"Create discount"
//	@Success		201				{object}	resmodel.NoDataResponse	"Discount created"
//	@Failure		400				{object}	utils.RequestError		"Bad request"
//	@Failure		500				{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount [post]
func (h *DiscountHandler) CreateDiscount(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	var request reqmodel.CreateDiscount
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
	err := h.Controller.CreateDiscount(c.Context(), request)
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
		Message: "Berhasil membuat diskon",
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

// UpdateDiscount
//
//	@Summary		Update Discount
//	@Description	Update an existing discount
//	@Tags			Discount
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string					true	"User ID to check"
//	@Param			id_unik			path		string					true	"Discount ID"
//	@Param			request			body		reqmodel.UpdateDiscount	true	"Update discount"
//	@Success		200				{object}	resmodel.NoDataResponse	"Discount updated"
//	@Failure		400				{object}	utils.RequestError		"Bad request"
//	@Failure		500				{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount/{id_unik} [put]
func (h *DiscountHandler) UpdateDiscount(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	discountID := c.Params("id_unik")
	var request reqmodel.UpdateDiscount
	request.ID = discountID
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
	err := h.Controller.UpdateDiscount(c.Context(), request)
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
		Message: "Berhasil mengubah diskon",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteDiscount
//
//	@Summary		Delete Discount
//	@Description	Delete a discount
//	@Tags			Discount
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string					true	"User ID to check"
//	@Param			id_unik			path		string					true	"Discount ID"
//	@Success		200				{object}	resmodel.NoDataResponse	"Discount deleted"
//	@Failure		400				{object}	utils.RequestError		"Bad request"
//	@Failure		500				{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount/{id_unik} [delete]
func (h *DiscountHandler) DeleteDiscount(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	discountID := c.Params("id_unik")

	// Call Controller
	err := h.Controller.DeleteDiscount(c.Context(), discountID, authUserID)
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
		Message: "Berhasil menghapus diskon",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
