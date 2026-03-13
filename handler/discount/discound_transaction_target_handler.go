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

type DiscountTransactionTargetHandler struct {
	Controller *discount_ctrl.DiscountTransactionTargetController
	Validate   *validator.Validate
}

func NewDiscountTransactionTargetHandler(
	r fiber.Router,
	vld *validator.Validate,
	controller *discount_ctrl.DiscountTransactionTargetController,
) {
	handler := &DiscountTransactionTargetHandler{
		Controller: controller,
		Validate:   vld,
	}

	rStrict := r.Group("discount-transaction")
	rStrict.Get("/", handler.ListDiscountTransactionTarget)
	rStrict.Get("/:id_unik", handler.GetDiscountTransactionTargetByID)
	rStrict.Post("/", handler.CreateDiscountTransactionTarget)
	rStrict.Put("/:id_unik", handler.UpdateDiscountTransactionTarget)
	rStrict.Delete("/:id_unik", handler.DeleteDiscountTransactionTarget)
}

// ListDiscountTransactionTarget
//
//	@Summary		List Discount Transaction Target
//	@Description	Get all discount transaction targets
//	@Tags			Discount Transaction Target
//	@Produce		json
//	@Param			search	query		string					false	"Search keyword"
//	@Param			page	query		int						false	"Page Number"
//	@Param			size	query		int						false	"Page Size"
//	@Success		200		{object}	resmodel.DatasResponse	"Discount transaction targets retrieved successfully"
//	@Failure		400		{object}	utils.RequestError		"Bad request"
//	@Failure		404		{object}	utils.RequestError		"Discount transaction targets not found"
//	@Failure		500		{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount-transaction [get]
func (h *DiscountTransactionTargetHandler) ListDiscountTransactionTarget(c *fiber.Ctx) error {
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
	datas, pagination, err := h.Controller.ListDiscountTransactionTarget(c.Context(), request)
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
		Message:    "Berhasil mengambil data target diskon transaksi",
		Datas:      datas,
		Pagination: pagination,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetDiscountTransactionTargetByID
//
//	@Summary		Get Discount Transaction Target By ID
//	@Description	Get a discount transaction target by its ID
//	@Tags			Discount Transaction Target
//	@Produce		json
//	@Param			id_unik	path		string					true	"Discount Transaction Target ID"
//	@Success		200		{object}	resmodel.DataResponse	"Discount transaction target retrieved successfully"
//	@Failure		400		{object}	utils.RequestError		"Bad request"
//	@Failure		404		{object}	utils.RequestError		"Discount transaction target not found"
//	@Failure		500		{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount-transaction/{id_unik} [get]
func (h *DiscountTransactionTargetHandler) GetDiscountTransactionTargetByID(c *fiber.Ctx) error {
	// Parse
	discountTransactionTargetID := c.Params("id_unik")

	// Call Controller
	data, err := h.Controller.GetDiscountTransactionTargetByID(c.Context(), discountTransactionTargetID)
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
		Message: "Berhasil mengambil data target diskon transaksi",
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// CreateDiscountTransactionTarget
//
//	@Summary		Create Discount Transaction Target
//	@Description	Create a new discount transaction target
//	@Tags			Discount Transaction Target
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string										true	"User ID to check"
//	@Param			request			body		reqmodel.CreateDiscountTransactionTarget	true	"Create discount transaction target"
//	@Success		201				{object}	resmodel.NoDataResponse						"Discount transaction target created"
//	@Failure		400				{object}	utils.RequestError							"Bad request"
//	@Failure		500				{object}	utils.RequestError							"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount-transaction [post]
func (h *DiscountTransactionTargetHandler) CreateDiscountTransactionTarget(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	var request reqmodel.CreateDiscountTransactionTarget
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
	err := h.Controller.CreateDiscountTransactionTarget(c.Context(), request)
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
		Message: "Berhasil membuat target diskon transaksi",
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

// UpdateDiscountTransactionTarget
//
//	@Summary		Update Discount Transaction Target
//	@Description	Update an existing discount transaction target
//	@Tags			Discount Transaction Target
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string										true	"User ID to check"
//	@Param			id_unik			path		string										true	"Discount Transaction Target ID"
//	@Param			request			body		reqmodel.UpdateDiscountTransactionTarget	true	"Update discount transaction target"
//	@Success		200				{object}	resmodel.NoDataResponse						"Discount transaction target updated"
//	@Failure		400				{object}	utils.RequestError							"Bad request"
//	@Failure		500				{object}	utils.RequestError							"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount-transaction/{id_unik} [put]
func (h *DiscountTransactionTargetHandler) UpdateDiscountTransactionTarget(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	discountTransactionTargetID := c.Params("id_unik")
	var request reqmodel.UpdateDiscountTransactionTarget
	request.ID = discountTransactionTargetID
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
	err := h.Controller.UpdateDiscountTransactionTarget(c.Context(), request)
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
		Message: "Berhasil mengubah target diskon transaksi",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteDiscountTransactionTarget
//
//	@Summary		Delete Discount Transaction Target
//	@Description	Delete a discount transaction target
//	@Tags			Discount Transaction Target
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string					true	"User ID to check"
//	@Param			id_unik			path		string					true	"Discount Transaction Target ID"
//	@Success		200				{object}	resmodel.NoDataResponse	"Discount transaction target deleted"
//	@Failure		400				{object}	utils.RequestError		"Bad request"
//	@Failure		500				{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/discount-transaction/{id_unik} [delete]
func (h *DiscountTransactionTargetHandler) DeleteDiscountTransactionTarget(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	discountTransactionTargetID := c.Params("id_unik")

	// Call Controller
	err := h.Controller.DeleteDiscountTransactionTarget(c.Context(), discountTransactionTargetID, authUserID)
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
		Message: "Berhasil menghapus target diskon transaksi",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
