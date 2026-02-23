package handlers

import (
	"errors"
	auth_ctrl "template/service/auth/controller/auth"
	reqmodel "template/service/auth/model/request"
	resmodel "template/service/auth/model/response"
	"template/service/auth/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	Controller *auth_ctrl.PostController
	Validate   *validator.Validate
}

func NewPostHandler(
	r fiber.Router,
	vld *validator.Validate,
	controller *auth_ctrl.PostController,
) {
	handler := &PostHandler{
		Controller: controller,
		Validate:   vld,
	}

	rStrict := r.Group("post")
	rStrict.Get("/", handler.ListPost)
	rStrict.Get("/:id", handler.GetPostByID)
	rStrict.Post("/", handler.CreatePost)
	rStrict.Put("/:id", handler.UpdatePost)
	rStrict.Delete("/:id", handler.DeletePost)
}

// ListPost
//
//	@Summary		List Post
//	@Description	Mendapatkan seluruh post
//	@Tags			Post
//	@Produce		json
//	@Param			search	query		string					false	"Search keyword"
//	@Param			page	query		int						false	"Page Number"
//	@Param			size	query		int						false	"Page Size"
//	@Success		200		{object}	resmodel.DatasResponse	"Post berhasil diambil"
//	@Failure		400		{object}	utils.RequestError		"Bad request"
//	@Failure		404		{object}	utils.RequestError		"Post not found"
//	@Failure		500		{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/post [get]
func (h *PostHandler) ListPost(c *fiber.Ctx) error {
	// Parse
	var request reqmodel.ListRequest
	if err := c.QueryParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("Format query tidak valid"),
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
	datas, pagination, err := h.Controller.ListPost(c.Context(), request)
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
		Message:    "Post berhasil diambil",
		Datas:      datas,
		Pagination: pagination,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetPostByID
//
//	@Summary		Get Post By ID
//	@Description	Mendapatkan post berdasarkan ID
//	@Tags			Post
//	@Produce		json
//	@Param			id	path		string					true	"Post ID"
//	@Success		200	{object}	resmodel.DataResponse	"Post berhasil diambil"
//	@Failure		400	{object}	utils.RequestError		"Bad request"
//	@Failure		404	{object}	utils.RequestError		"Post not found"
//	@Failure		500	{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/post/{id} [get]
func (h *PostHandler) GetPostByID(c *fiber.Ctx) error {
	// Parse
	postID := c.Params("id")

	// Call Controller
	data, err := h.Controller.GetPostByID(c.Context(), postID)
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
		Message: "Post berhasil diambil",
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// CreatePost
//
//	@Summary		Create Post
//	@Description	Membuat post baru
//	@Tags			Post
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string					true	"User ID to check"
//	@Param			request			body		reqmodel.CreatePost		true	"Create post"
//	@Success		201				{object}	resmodel.NoDataResponse	"Post Created"
//	@Failure		400				{object}	utils.RequestError		"Bad request"
//	@Failure		500				{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/post [post]
func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	var request reqmodel.CreatePost
	request.AuthUserID = authUserID
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("Format JSON tidak valid"),
		))
	}

	// Validate Input
	if err := h.Validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("Validasi gagal: "+err.Error()),
		))
	}

	// Call Controller
	err := h.Controller.CreatePost(c.Context(), request)
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
		Message: "Post berhasil dibuat",
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

// UpdatePost
//
//	@Summary		Update Post
//	@Description	Membuat post baru
//	@Tags			Post
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string					true	"User ID to check"
//	@Param			id				path		string					true	"Post ID"
//	@Param			request			body		reqmodel.UpdatePost		true	"Update post"
//	@Success		201				{object}	resmodel.NoDataResponse	"Post Updated"
//	@Failure		400				{object}	utils.RequestError		"Bad request"
//	@Failure		500				{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/post/{id} [put]
func (h *PostHandler) UpdatePost(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	postID := c.Params("id")
	var request reqmodel.UpdatePost
	request.ID = postID
	request.AuthUserID = authUserID
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("Format JSON tidak valid"),
		))
	}

	// Validate Input
	if err := h.Validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("Validasi gagal: "+err.Error()),
		))
	}

	// Call Controller
	err := h.Controller.UpdatePost(c.Context(), request)
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
		Message: "Post berhasil diubah",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// DeletePost
//
//	@Summary		Delete Post
//	@Description	Menghapus post
//	@Tags			Post
//	@Produce		json
//	@Accept			json
//	@Param			X-Auth-User-Id	header		string					true	"User ID to check"
//	@Param			id				path		string					true	"Post ID"
//	@Success		201				{object}	resmodel.NoDataResponse	"Post Deleted"
//	@Failure		400				{object}	utils.RequestError		"Bad request"
//	@Failure		500				{object}	utils.RequestError		"Server error"
//	@Security		ApiKeyAuth
//	@Router			/ndi/post/{id} [delete]
func (h *PostHandler) DeletePost(c *fiber.Ctx) error {
	// Get Header
	authUserID := c.Get("X-Auth-User-Id")
	if authUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(utils.GeneralErrorResponse(
			fiber.StatusBadRequest,
			errors.New("header X-Auth-User-Id tidak ditemukan"),
		))
	}

	// Parse
	postID := c.Params("id")

	// Call Controller
	err := h.Controller.DeletePost(c.Context(), postID, authUserID)
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
		Message: "Post berhasil dihapus",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
