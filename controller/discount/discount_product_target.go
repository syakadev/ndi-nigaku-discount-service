package discount_ctrl

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

func NewDiscountProductTargetController(
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

func (c *PostController) ListDiscountProductTarget(ctx context.Context, request reqmodel.ListRequest) ([]interface{}, interface{}, error) {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Get
	productsTarget, pagination, err := dbprocess.ListDiscountProductTarget(reqCtx, c.pgxConn, request)
	if err != nil {
		c.log.Error("gagal memproses daftar diskon product:", err)
		return nil, nil, utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengambil daftar diskon product",
		}
	}
	if len(productsTarget) == 0 {
		return nil, nil, utils.StandardError{
			Code:    http.StatusNotFound,
			Message: "Tidak ada data diskon product yang tersedia",
		}
	}

	// Return Success
	return productsTarget, pagination, nil
}

func (c *PostController) GetDiscountProductTargetByID(ctx context.Context, productID string) (*resmodel.DiscountProductTarget, error) {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Get
	product, err := dbprocess.GetDiscountProductTargetByID(reqCtx, c.pgxConn, productID)
	if err != nil {
		c.log.Error("gagal memproses product:", err)
		return nil, utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengambil product",
		}
	}

	// Return Success
	return product, nil
}

func (c *PostController) CreateDiscountProductTarget(ctx context.Context, request reqmodel.CreateDiscountProductTarget) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Create
	err := dbprocess.CreateDiscountProductTarget(reqCtx, c.pgxConn, request)
	if err != nil {
		c.log.Error("gagal memproses diskon product:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal membuat diskon product",
		}
	}

	// Return Success
	return nil
}

func (c *PostController) UpdateDiscountProductTarget(ctx context.Context, request reqmodel.UpdateDiscountProductTarget) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Update
	err := dbprocess.UpdateDiscountProductTarget(reqCtx, c.pgxConn, request)
	if err != nil {
		c.log.Error("gagal memproses diskon produc:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal mengubah diskon product",
		}
	}

	// Return Success
	return nil
}

func (c *PostController) DeleteDiscountProductTarget(ctx context.Context, productID, authUserID string) error {
	// Context Timeout
	reqCtx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	// Delete
	err := dbprocess.DeleteDiscountProductTarget(reqCtx, c.pgxConn, productID, authUserID)
	if err != nil {
		c.log.Error("gagal memproses diskon product:", err)
		return utils.StandardError{
			Code:    http.StatusInternalServerError,
			Message: "Gagal menghapus diskon product",
		}
	}

	// Return Success
	return nil
}
