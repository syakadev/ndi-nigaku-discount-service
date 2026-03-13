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

type TransactionDiscountAppliedHandler struct {
	Controller *discount_ctrl.TransactionDiscountAppliedController
	Validate   *validator.Validate
}

func NewTransactionDiscountAppliedHandler(
	r fiber.Router,
	vld *validator.Validate,
	controller *discount_ctrl.TransactionDiscountAppliedController,
) {
	handler := &TransactionDiscountAppliedHandler{
		Controller: controller,
		Validate:   vld,
	}

	rStrict := r.Group("transaction-discount-applied")
	rStrict.Get("/", handler.ListTransactionDiscountApplied)
	rStrict.Get("/by-target/:id_unik", handler.ListTransactionDiscountAppliedByTargetID)
	rStrict.Post("/", handler.CreateTransactionDiscountApplied)
	rStrict.Delete("/:id_unik", handler.DeleteTransactionDiscountApplied)
}

// ListTransactionDiscountApplied
//
//	@Summary		List Transaction Discount Applied
//	@Description	Get all transaction discount applied
//	@Tags			Transaction Discount Applied
//	@Produce		json
//	@Param			search	query		string					false	"Search keyword"
//	@Param			page	query		int						false	"Page Number"
//	@Param			size	query		int						false	"Page Size"
//	@Success		200		{object}	resmodel.DatasResponse	"Product discount applied retrieved successfully"
//	@Failure		400		{object}	utils.RequestError		"Bad request"
//	@Failure		404		{object}	utils.RequestError		"Product discount applied not found"
//	@Failure		500		{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/transaction-discount-applied [get]
func (h *TransactionDiscountAppliedHandler) ListTransactionDiscountApplied(c *fiber.Ctx) error {
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
	datas, pagination, err := h.Controller.ListTransactionDiscountApplied(c.Context(), request)
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
		Message:    "Berhasil mengambil data diskon transaksi yang diterapkan",
		Datas:      datas,
		Pagination: pagination,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// ListTransactionDiscountAppliedByTargetID
//
//	@Summary		List Transaction Discount Applied By Discount Target ID
//	@Description	Get a list of transaction discount applied by its discount target ID
//	@Tags			Transaction Discount Applied
//	@Produce		json
//	@Param			id_unik	path		string					true	"Transaction Discount Target ID"
//	@Success		200		{object}	resmodel.DatasResponse	"Transaction discount applied retrieved successfully"
//	@Failure		400		{object}	utils.RequestError		"Bad request"
//	@Failure		404		{object}	utils.RequestError		"Transaction discount applied not found"
//	@Failure		500		{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/transaction-discount-applied/by-target/{id_unik} [get]
func (h *TransactionDiscountAppliedHandler) ListTransactionDiscountAppliedByTargetID(c *fiber.Ctx) error {
	// Parse
	transactionDiscountTargetID := c.Params("id_unik")
	var request reqmodel.ListRequest
	if err := c.QueryParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("format query tidak valid"),
		))
	}

	// Call Controller
	datas, pagination, err := h.Controller.ListTransactionDiscountAppliedByTargetID(c.Context(), transactionDiscountTargetID, request)
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
		Message:    "Berhasil mengambil data diskon transaksi yang diterapkan",
		Datas:      datas,
		Pagination: pagination,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// CreateTransactionDiscountApplied
//
//	@Summary		Create Transaction Discount Applied
//	@Description	Create a new transaction discount applied
//	@Tags			Transaction Discount Applied
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string										true	"User ID to check"
//	@Param			request			body		reqmodel.CreateTransactionDiscountApplied	true	"Create transaction discount applied"
//	@Success		201				{object}	resmodel.NoDataResponse						"Transaction discount applied created"
//	@Failure		400				{object}	utils.RequestError							"Bad request"
//	@Failure		500				{object}	utils.RequestError							"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/transaction-discount-applied [post]
func (h *TransactionDiscountAppliedHandler) CreateTransactionDiscountApplied(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	var request reqmodel.CreateTransactionDiscountApplied
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
	err := h.Controller.CreateTransactionDiscountApplied(c.Context(), request)
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
		Message: "Berhasil membuat diskon transaksi yang diterapkan",
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

// DeleteTransactionDiscountApplied
//
//	@Summary		Delete Transaction Discount Applied
//	@Description	Delete a transaction discount applied
//	@Tags			Transaction Discount Applied
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string					true	"User ID to check"
//	@Param			id_unik			path		string					true	"Transaction Discount Applied ID"
//	@Success		200				{object}	resmodel.NoDataResponse	"Transaction discount applied deleted"
//	@Failure		400				{object}	utils.RequestError		"Bad request"
//	@Failure		500				{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/transaction-discount-applied/{id_unik} [delete]
func (h *TransactionDiscountAppliedHandler) DeleteTransactionDiscountApplied(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	transactionDiscountAppliedID := c.Params("id_unik")

	// Call Controller
	err := h.Controller.DeleteTransactionDiscountApplied(c.Context(), transactionDiscountAppliedID, authUserID)
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
		Message: "Berhasil menghapus diskon transaksi yang diterapkan",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
