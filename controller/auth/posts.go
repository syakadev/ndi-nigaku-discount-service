package auth_ctrl

import (
	"context"
	"net/http"
	"time"

	dbprocess "service/discount/api/database/process"
	reqmodel "service/discount/api/model/request"
	resmodel "service/discount/api/model/response"
	"service/discount/api/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostController struct {
	contextTimeout time.Duration
	pgxConn        *pgxpool.Pool
	log            *utils.AppLogger
}

func NewPostController(
	conn *pgxpool.Pool,
	timeout time.Duration,
	log *utils.AppLogger,
) *PostController {
	return &PostController{
		pgxConn:        conn,
		contextTimeout: timeout,
		log:            log,
	}
}

func (c *PostController) ListPost(ctx context.Context, request reqmodel.ListRequest) ([]interface{}, interface{}, error) {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Get
	posts, pagination, err := dbprocess.ListPost(reqCtx, c.pgxConn, request)
	if err != nil {
		c.log.Error("gagal memproses daftar post:", err)
		return nil, nil, utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengambil daftar post",
		}
	}
	if len(posts) == 0 {
		return nil, nil, utils.StandardError{
			Code:    http.StatusNotFound,
			Message: "Tidak ada data post yang tersedia",
		}
	}

	// Return Success
	return posts, pagination, nil
}

func (c *PostController) GetPostByID(ctx context.Context, postID string) (*resmodel.Post, error) {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Get
	post, err := dbprocess.GetPostByID(reqCtx, c.pgxConn, postID)
	if err != nil {
		c.log.Error("gagal memproses post:", err)
		return nil, utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengambil post",
		}
	}

	// Return Success
	return post, nil
}

func (c *PostController) CreatePost(ctx context.Context, request reqmodel.CreatePost) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Create
	err := dbprocess.CreatePost(reqCtx, c.pgxConn, request)
	if err != nil {
		c.log.Error("gagal memproses post:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal membuat post",
		}
	}

	// Return Success
	return nil
}

func (c *PostController) UpdatePost(ctx context.Context, request reqmodel.UpdatePost) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Update
	err := dbprocess.UpdatePost(reqCtx, c.pgxConn, request)
	if err != nil {
		c.log.Error("gagal memproses post:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengubah post",
		}
	}

	// Return Success
	return nil
}

func (c *PostController) DeletePost(ctx context.Context, postID, authUserID string) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Delete
	err := dbprocess.DeletePost(reqCtx, c.pgxConn, postID, authUserID)
	if err != nil {
		c.log.Error("gagal memproses post:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal menghapus post",
		}
	}

	// Return Success
	return nil
}
