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

type ProductDiscountAppliedHandler struct {
	Controller *discount_ctrl.ProductDiscountAppliedController
	Validate   *validator.Validate
}

func NewProductDiscountAppliedHandler(
	r fiber.Router,
	vld *validator.Validate,
	controller *discount_ctrl.ProductDiscountAppliedController,
) {
	handler := &ProductDiscountAppliedHandler{
		Controller: controller,
		Validate:   vld,
	}

	rStrict := r.Group("product-discount-applied")
	rStrict.Get("/", handler.ListProductDiscountApplied)
	rStrict.Get("/by-discount/:id", handler.ListProductDiscountAppliedByDiscountID)
	rStrict.Post("/", handler.CreateProductDiscountApplied)
	rStrict.Delete("/:id", handler.DeleteProductDiscountApplied)
}

// ListProductDiscountApplied
//
//	@Summary		List Product Discount Applied
//	@Description	Get all product discount applied
//	@Tags			Product Discount Applied
//	@Produce		json
//	@Param			search	query		string					false	"Search keyword"
//	@Param			page	query		int						false	"Page Number"
//	@Param			size	query		int						false	"Page Size"
//	@Success		200		{object}	resmodel.DatasResponse	"Product discount applied retrieved successfully"
//	@Failure		400		{object}	utils.RequestError		"Bad request"
//	@Failure		404		{object}	utils.RequestError		"Product discount applied not found"
//	@Failure		500		{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/product-discount-applied [get]
func (h *ProductDiscountAppliedHandler) ListProductDiscountApplied(c *fiber.Ctx) error {
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
	datas, pagination, err := h.Controller.ListProductDiscountApplied(c.Context(), request)
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
		Message:    "Berhasil mengambil data diskon produk yang diterapkan",
		Datas:      datas,
		Pagination: pagination,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// ListProductDiscountAppliedByDiscountID
//
//	@Summary		List Product Discount Applied By Discount ID
//	@Description	Get a list of product discount applied by its discount ID
//	@Tags			Product Discount Applied
//	@Produce		json
//	@Param			id	path		string					true	"Product Discount ID"
//	@Success		200	{object}	resmodel.DatasResponse	"Product discount applied retrieved successfully"
//	@Failure		400	{object}	utils.RequestError		"Bad request"
//	@Failure		404	{object}	utils.RequestError		"Product discount applied not found"
//	@Failure		500	{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/product-discount-applied/by-discount/{id} [get]
func (h *ProductDiscountAppliedHandler) ListProductDiscountAppliedByDiscountID(c *fiber.Ctx) error {
	// Parse
	productDiscountID := c.Params("id")
	var request reqmodel.ListRequest
	if err := c.QueryParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("format query tidak valid"),
		))
	}

	// Call Controller
	datas, pagination, err := h.Controller.ListProductDiscountAppliedByProductDiscountID(c.Context(), productDiscountID, request)
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
		Message:    "Berhasil mengambil data diskon produk yang diterapkan",
		Datas:      datas,
		Pagination: pagination,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// CreateProductDiscountApplied
//
//	@Summary		Create Product Discount Applied
//	@Description	Create a new product discount applied
//	@Tags			Product Discount Applied
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string									true	"User ID to check"
//	@Param			request			body		reqmodel.CreateProductDiscountApplied	true	"Create product discount applied"
//	@Success		201				{object}	resmodel.NoDataResponse					"Product discount applied created"
//	@Failure		400				{object}	utils.RequestError						"Bad request"
//	@Failure		500				{object}	utils.RequestError						"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/product-discount-applied [post]
func (h *ProductDiscountAppliedHandler) CreateProductDiscountApplied(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	var request reqmodel.CreateProductDiscountApplied
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
	err := h.Controller.CreateProductDiscountApplied(c.Context(), request)
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
		Message: "Berhasil membuat diskon produk yang diterapkan",
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

// DeleteProductDiscountApplied
//
//	@Summary		Delete Product Discount Applied
//	@Description	Delete a product discount applied
//	@Tags			Product Discount Applied
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string					true	"User ID to check"
//	@Param			id				path		string					true	"Product Discount Applied ID"
//	@Success		200				{object}	resmodel.NoDataResponse	"Product discount applied deleted"
//	@Failure		400				{object}	utils.RequestError		"Bad request"
//	@Failure		500				{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/product-discount-applied/{id} [delete]
func (h *ProductDiscountAppliedHandler) DeleteProductDiscountApplied(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	productDiscountAppliedID := c.Params("id")

	// Call Controller
	err := h.Controller.DeleteProductDiscountApplied(c.Context(), productDiscountAppliedID, authUserID)
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
		Message: "Berhasil menghapus diskon produk yang diterapkan",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
